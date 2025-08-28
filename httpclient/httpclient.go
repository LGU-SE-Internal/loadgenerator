package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/Lincyaw/loadgenerator/stats"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

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

type HttpClient struct {
	client   *http.Client
	headers  map[string]string
	reqCount int
	mu       sync.Mutex
	tracer   trace.Tracer
	timeout  time.Duration
	MaxLat   time.Duration
}

func NewCustomClient() *HttpClient {
	transport := &http.Transport{
		MaxIdleConns:        100,              // 最大空闲连接数
		MaxIdleConnsPerHost: 10,               // 每个主机的最大空闲连接数
		IdleConnTimeout:     90 * time.Second, // 空闲连接超时
		DisableKeepAlives:   false,            // 启用keep-alive
	}

	httpClient := &http.Client{
		Transport: transport,
		Timeout:   20 * time.Second, // 默认20秒超时
	}

	tracer := otel.Tracer("loadgenerator/httpclient")

	return &HttpClient{
		client:  httpClient,
		headers: make(map[string]string),
		tracer:  tracer,
		timeout: 20 * time.Second,
	}
}

func (c *HttpClient) AddHeader(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.headers[key] = value
}

func (c *HttpClient) GetTimeout() time.Duration {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.timeout
}

func (c *HttpClient) SendRequest(method, url string, body interface{}) (*http.Response, error) {
	return c.SendRequestWithContext(context.Background(), method, url, body)
}

func (c *HttpClient) SendRequestWithContext(ctx context.Context, method, url string, body interface{}) (*http.Response, error) {
	startTime := time.Now()

	ctx, span := c.tracer.Start(ctx, fmt.Sprintf("HTTP %s %s", method, url),
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			attribute.String("http.method", method),
			attribute.String("http.url", url),
		))
	defer span.End()

	c.mu.Lock()
	c.reqCount++
	reqCountSnapshot := c.reqCount
	c.mu.Unlock()

	span.SetAttributes(attribute.Int("http.request_count", reqCountSnapshot))

	jsonData, err := json.Marshal(body)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Failed to marshal request body")
		span.SetAttributes(
			attribute.String("error.type", "marshal_error"),
			attribute.String("error.message", err.Error()),
		)
		return nil, err
	}

	span.SetAttributes(attribute.Int("http.request_content_length", len(jsonData)))

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Failed to create HTTP request")
		span.SetAttributes(
			attribute.String("error.type", "request_creation_error"),
			attribute.String("error.message", err.Error()),
		)
		return nil, err
	}

	c.mu.Lock()
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}
	c.mu.Unlock()

	propagator := otel.GetTextMapPropagator()
	propagator.Inject(ctx, propagation.HeaderCarrier(req.Header))

	resp, err := c.client.Do(req)

	latency := time.Since(startTime)
	c.MaxLat = max(c.MaxLat, latency)
	logrus.Warnf("HTTP %s %s took %v, max latency: %v", method, url, latency, c.MaxLat)
	statsObj := stats.GlobalLatencyManager.GetOrCreateStats(url, method)
	statsObj.AddLatency(latency)

	span.SetAttributes(attribute.Int64("http.request_duration_ms", latency.Nanoseconds()/1000000))

	if err != nil {
		span.RecordError(err)
		if ctx.Err() == context.DeadlineExceeded {
			span.SetStatus(codes.Error, "HTTP request timeout")
			span.SetAttributes(
				attribute.String("error.type", "timeout"),
				attribute.Int64("timeout_duration_ms", c.GetTimeout().Milliseconds()),
			)
			return nil, fmt.Errorf("request timeout after %v: %w", c.GetTimeout(), err)
		}
		if err.Error() == "context deadline exceeded" {
			span.SetStatus(codes.Error, "HTTP request timeout")
			span.SetAttributes(
				attribute.String("error.type", "context_timeout"),
				attribute.Int64("timeout_duration_ms", c.GetTimeout().Milliseconds()),
			)
			return nil, fmt.Errorf("request timeout after %v: %w", c.GetTimeout(), err)
		}
		span.SetStatus(codes.Error, "HTTP request failed")
		span.SetAttributes(
			attribute.String("error.type", "network_error"),
			attribute.String("error.message", err.Error()),
		)
		return nil, err
	}

	span.SetAttributes(
		attribute.Int("http.status_code", resp.StatusCode),
		attribute.String("http.status_text", resp.Status),
	)

	if resp.StatusCode >= 400 {
		span.SetStatus(codes.Error, fmt.Sprintf("HTTP error: %s", resp.Status))
	} else {
		span.SetStatus(codes.Ok, "Request completed successfully")
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Failed to read response body")
		return nil, err
	}
	resp.Body.Close()

	span.SetAttributes(attribute.Int("http.response_content_length", len(respBody)))

	resp.Body = io.NopCloser(bytes.NewBuffer(respBody))

	return resp, nil
}
