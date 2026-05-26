package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	headers    map[string]string
}

type Option func(*Client)

func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = d
	}
}

func WithHeader(key, value string) Option {
	return func(c *Client) {
		c.headers[key] = value
	}
}

func WithBearerToken(token string) Option {
	return WithHeader("Authorization", "Bearer "+token)
}

func WithAPIKey(key string) Option {
	return WithHeader("X-API-Key", key)
}

func New(baseURL string, opts ...Option) *Client {
	c := &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: 30 * time.Second},
		headers: map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Get gọi GET request, unmarshal response vào dest
func (c *Client) Get(ctx context.Context, path string, dest interface{}) error {
	return c.do(ctx, http.MethodGet, path, nil, dest)
}

// Post gọi POST request với body, unmarshal response vào dest
func (c *Client) Post(ctx context.Context, path string, body, dest interface{}) error {
	return c.do(ctx, http.MethodPost, path, body, dest)
}

// Put gọi PUT request
func (c *Client) Put(ctx context.Context, path string, body, dest interface{}) error {
	return c.do(ctx, http.MethodPut, path, body, dest)
}

// Delete gọi DELETE request
func (c *Client) Delete(ctx context.Context, path string, dest interface{}) error {
	return c.do(ctx, http.MethodDelete, path, nil, dest)
}

func (c *Client) do(ctx context.Context, method, path string, body, dest interface{}) error {
	var reqBody io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, reqBody)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	for k, v := range c.headers {
		req.Header.Set(k, v)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	// Coi 2xx là thành công
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(respBody))
	}

	if dest != nil {
		if err := json.Unmarshal(respBody, dest); err != nil {
			return fmt.Errorf("unmarshal response: %w", err)
		}
	}

	return nil
}
