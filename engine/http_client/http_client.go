package http_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mega/engine/logger"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HttpClient struct {
	client *http.Client
}

// NewHttpClient 创建 HTTP 客户端
func NewHttpClient(timeout time.Duration) *HttpClient {
	if timeout <= 0 {
		timeout = 10 * time.Second
	}
	return &HttpClient{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (h *HttpClient) Get(
	rawURL string,
	params map[string]string,
	headers map[string]string,
) ([]byte, error) {

	// 默认 Content-Type（如果外面没传）
	if headers == nil {
		headers = make(map[string]string)
	}
	if _, ok := headers["Content-Type"]; !ok {
		headers["Content-Type"] = "application/json"
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	// query 参数
	q := u.Query()
	if headers["Content-Type"] == "application/json" {

	} else if headers["Content-Type"] == "application/x-www-form-urlencoded" {
		for k, v := range params {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}

	// req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, u.String(), nil)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, u.String(), strings.NewReader(q.Encode()))
	if err != nil {
		return nil, err
	}

	// headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Warn("http_client Get resp error=", err)
		return nil, nil
	} else {
		return bodyData, nil
	}
}

func (h *HttpClient) Post(
	rawURL string,
	params map[string]string,
	headers map[string]string,
) ([]byte, error) {

	// 默认 Content-Type（如果外面没传）
	if headers == nil {
		headers = make(map[string]string)
	}
	if _, ok := headers["Content-Type"]; !ok {
		headers["Content-Type"] = "application/json"
	}

	u, err2 := url.Parse(rawURL)
	if err2 != nil {
		return nil, err2
	}

	var req *http.Request
	var err error
	var bodyBytes []byte
	// query 参数
	q := u.Query()
	if headers["Content-Type"] == "application/json" {
		// map -> JSON
		bodyBytes, err = json.Marshal(params)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			u.String(),
			// nil,
			bytes.NewReader(bodyBytes),
		)
	} else if headers["Content-Type"] == "application/x-www-form-urlencoded" {
		for k, v := range params {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
		req, err = http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			u.String(),
			strings.NewReader(q.Encode()),
		)
	}
	if err != nil {
		return nil, err
	}

	// headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 可选：检查 HTTP 状态码
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("http status %d", resp.StatusCode)
	}

	bodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Warn("http_client Post resp error=", err)
		return nil, err
	}

	return bodyData, nil
}
