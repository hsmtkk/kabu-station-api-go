package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *clientImpl) post(endpoint string, body []byte) ([]byte, error) {
	c.logger.Debug("post", "endpoint", endpoint)
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.Post failed: %w", err)
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll failed: %w", err)
	}
	c.logger.Debug("post", "response", string(respBytes))
	return respBytes, nil
}

func (c *clientImpl) getWithToken(endpoint string, query map[string]string) ([]byte, error) {
	c.logger.Debug("getWithToken", "endpoint", endpoint, "query", query)
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("url.Parse failed: %w", err)
	}
	q := u.Query()
	for k, v := range query {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	queryURL := u.String()
	req, err := http.NewRequest(http.MethodGet, queryURL, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest failed: %w", err)
	}
	req.Header.Set("X-API-KEY", c.token)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.Do failed: %w", err)
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll failed: %w", err)
	}
	c.logger.Debug("getWithToken", "response", string(respBytes))
	return respBytes, nil
}

func (c *clientImpl) putWithToken(endpoint string) ([]byte, error) {
	c.logger.Debug("putWithToken", "endpoint", endpoint)
	req, err := http.NewRequest(http.MethodPut, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest failed: %w", err)
	}
	req.Header.Set("X-API-KEY", c.token)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.Do failed: %w", err)
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll failed: %w", err)
	}
	c.logger.Debug("putWithToken", "response", string(respBytes))
	return respBytes, nil
}
