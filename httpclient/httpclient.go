package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"sync"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// RequestStats 保存每个请求的统计信息。
type RequestStats struct {
	Success      int
	Failed       int
	RequestBody  []string
	ResponseBody []string
}

type RequestStatsKey struct {
	URL    string
	Method string
}

// HttpClient 是自定义的 HTTP 客户端，包含请求计数器和头信息。
type HttpClient struct {
	client       *http.Client
	headers      map[string]string
	reqCount     int
	mu           sync.Mutex
	requestStats map[RequestStatsKey]RequestStats
	tracer       trace.Tracer
}

// NewCustomClient 创建并返回一个新的 HttpClient 实例。
func NewCustomClient() *HttpClient {
	// 使用 otelhttp 包装的 HTTP 客户端进行自动 instrumentation
	httpClient := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	// 获取全局 tracer
	tracer := otel.Tracer("loadgenerator/httpclient")

	return &HttpClient{
		client:       httpClient,
		headers:      make(map[string]string),
		requestStats: make(map[RequestStatsKey]RequestStats),
		tracer:       tracer,
	}
}

// AddHeader 向 HttpClient 添加一个头信息。
func (c *HttpClient) AddHeader(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.headers[key] = value
}

// SendRequest 发送 HTTP 请求并统计请求数量和详细信息。
func (c *HttpClient) SendRequest(method, url string, body interface{}) (*http.Response, error) {
	return c.SendRequestWithContext(context.Background(), method, url, body)
}

// SendRequestWithContext 发送带上下文的 HTTP 请求并统计请求数量和详细信息。
func (c *HttpClient) SendRequestWithContext(ctx context.Context, method, url string, body interface{}) (*http.Response, error) {
	// 创建 span
	ctx, span := c.tracer.Start(ctx, fmt.Sprintf("HTTP %s", method),
		trace.WithAttributes(
			attribute.String("http.method", method),
			attribute.String("http.url", url),
		))
	defer span.End()

	c.mu.Lock()
	c.reqCount++
	c.mu.Unlock()

	// 将 body 转换为 JSON
	jsonData, err := json.Marshal(body)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Failed to marshal request body")
		return nil, err
	}

	// 添加请求体大小到 span 属性
	span.SetAttributes(attribute.Int("http.request_content_length", len(jsonData)))

	// 创建新的 HTTP 请求，使用带有 trace 信息的 context
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Failed to create HTTP request")
		return nil, err
	}

	// 添加头信息，需要持有锁避免并发读写
	c.mu.Lock()
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}
	c.mu.Unlock()

	// 发送请求并记录请求和响应信息
	resp, err := c.client.Do(req)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "HTTP request failed")
		return nil, err
	}

	// 添加响应状态码到 span 属性
	span.SetAttributes(
		attribute.Int("http.status_code", resp.StatusCode),
		attribute.String("http.status_text", resp.Status),
	)

	// 根据状态码设置 span 状态
	if resp.StatusCode >= 400 {
		span.SetStatus(codes.Error, fmt.Sprintf("HTTP error: %s", resp.Status))
	} else {
		span.SetStatus(codes.Ok, "Request completed successfully")
	}

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Failed to read response body")
		return nil, err
	}
	// 关闭响应体
	resp.Body.Close()

	// 添加响应体大小到 span 属性
	span.SetAttributes(attribute.Int("http.response_content_length", len(respBody)))

	// 重新创建响应体以便后续处理
	resp.Body = io.NopCloser(bytes.NewBuffer(respBody))

	// 记录请求和响应信息
	c.logRequestResponse(req, resp, jsonData, respBody)

	return resp, nil
}

// logRequestResponse 记录请求和响应信息。
func (c *HttpClient) logRequestResponse(req *http.Request, resp *http.Response, reqBody, respBody []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	key := RequestStatsKey{
		URL:    req.URL.String(),
		Method: req.Method,
	}
	if _, ok := c.requestStats[key]; !ok {
		c.requestStats[key] = RequestStats{
			RequestBody:  make([]string, 0),
			ResponseBody: make([]string, 0),
		}
	}

	// 从当前上下文获取 span（如果存在）
	if span := trace.SpanFromContext(req.Context()); span.IsRecording() {
		// 记录请求/响应信息到 span
		span.AddEvent("request_logged", trace.WithAttributes(
			attribute.String("request.url", req.URL.String()),
			attribute.String("request.method", req.Method),
			attribute.Int("response.status_code", resp.StatusCode),
		))
	}

	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		value := c.requestStats[key]
		value.Success += 1

		if rand.Int()%10 == 0 {
			value.RequestBody = append(value.RequestBody, string(reqBody))
			value.ResponseBody = append(value.ResponseBody, string(respBody))
		}
		c.requestStats[key] = value
		return
	}
	value := c.requestStats[key]
	value.Failed += 1
	if rand.Int()%5 == 0 {
		value.RequestBody = append(value.RequestBody, string(reqBody))
		value.ResponseBody = append(value.ResponseBody, string(respBody))
	}
	c.requestStats[key] = value
}

// GetRequestCount 返回已经发送的请求数量。
func (c *HttpClient) GetRequestCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.reqCount
}

// GetRequestStats 返回所有请求的统计信息。
func (c *HttpClient) GetRequestStats() map[RequestStatsKey]RequestStats {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 创建 requestStats 的副本
	statsCopy := make(map[RequestStatsKey]RequestStats, len(c.requestStats))
	for key, value := range c.requestStats {
		newv := RequestStats{}
		newv.Success = value.Success
		newv.Failed = value.Failed
		newv.RequestBody = make([]string, len(value.RequestBody))
		newv.ResponseBody = make([]string, len(value.ResponseBody))
		copy(newv.RequestBody, value.RequestBody)
		copy(newv.ResponseBody, value.ResponseBody)
		statsCopy[key] = newv
	}
	return statsCopy
}
func GenerateMarkdownTable(data map[RequestStatsKey]RequestStats) string {
	var sb strings.Builder

	// 表头
	sb.WriteString("| URL | Method | Success | Failed | Request Body | Response Body |\n")
	sb.WriteString("| --- | ------ | ------- | ------ | ------------ | ------------- |\n")

	// 遍历 map 并生成表格行
	for key, stats := range data {
		requestBody := strings.Join(stats.RequestBody, "<br>")
		responseBody := strings.Join(stats.ResponseBody, "<br>")
		sb.WriteString(fmt.Sprintf("| %s | %s | %d | %d | %s | %s |\n",
			key.URL, key.Method, stats.Success, stats.Failed, requestBody, responseBody))
	}

	return sb.String()
}
