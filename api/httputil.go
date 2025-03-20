package api

import (
	"fmt"
	"io"
	"net/http"
)

func (c *clientImpl) invokeHTTP(req *http.Request) ([]byte, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.Do failed: %w", err)
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll failed: %w", err)
	}
	return respBytes, nil
}

func (c *clientImpl) invokeHTTPWithTokenHeader(req *http.Request) ([]byte, error) {
	req.Header.Set("X-API-KEY", c.token)
	fmt.Printf("X-API-KEY: %s\n", c.token) // debug
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.Do failed: %w", err)
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll failed: %w", err)
	}
	return respBytes, nil
}
