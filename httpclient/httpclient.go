package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
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
}

func NewCustomClient() *HttpClient {
	httpClient := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	tracer := otel.Tracer("loadgenerator/httpclient")

	return &HttpClient{
		client:  httpClient,
		headers: make(map[string]string),
		tracer:  tracer,
	}
}

func (c *HttpClient) AddHeader(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.headers[key] = value
}

func (c *HttpClient) SendRequest(method, url string, body interface{}) (*http.Response, error) {
	return c.SendRequestWithContext(context.Background(), method, url, body)
}

func (c *HttpClient) SendRequestWithContext(ctx context.Context, method, url string, body interface{}) (*http.Response, error) {
	ctx, span := c.tracer.Start(ctx, fmt.Sprintf("HTTP %s", method),
		trace.WithAttributes(
			attribute.String("http.method", method),
			attribute.String("http.url", url),
		))
	defer span.End()

	c.mu.Lock()
	c.reqCount++
	c.mu.Unlock()

	jsonData, err := json.Marshal(body)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Failed to marshal request body")
		return nil, err
	}

	span.SetAttributes(attribute.Int("http.request_content_length", len(jsonData)))

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Failed to create HTTP request")
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
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "HTTP request failed")
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
