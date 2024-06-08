package httpclient

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
)

// HttpClient 是自定义的 HTTP 客户端，包含请求计数器和头信息。
type HttpClient struct {
	client   *http.Client
	headers  map[string]string
	reqCount int
	mu       sync.Mutex
}

// NewCustomClient 创建并返回一个新的 HttpClient 实例。
func NewCustomClient() *HttpClient {
	return &HttpClient{
		client:  &http.Client{},
		headers: make(map[string]string),
	}
}

// AddHeader 向 HttpClient 添加一个头信息。
func (c *HttpClient) AddHeader(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.headers[key] = value
}

// SendRequest 发送 HTTP 请求并统计请求数量。
func (c *HttpClient) SendRequest(method, url string, body interface{}) (*http.Response, error) {
	c.mu.Lock()
	c.reqCount++
	c.mu.Unlock()

	// 将 body 转换为 JSON
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// 创建新的 HTTP 请求
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// 添加头信息
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	// 发送请求
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetRequestCount 返回已经发送的请求数量。
func (c *HttpClient) GetRequestCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.reqCount
}
