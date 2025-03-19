package token

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client interface {
	Token(apiPassword string) (int, string, error)
}

func New(client *http.Client, baseURL string) Client {
	return &clientImpl{client, baseURL}
}

type clientImpl struct {
	client  *http.Client
	baseURL string
}

type requestSchema struct {
	APIPassword string `json:"APIPassword"`
}

type responseSchema struct {
	ResultCode int    `json:"ResultCode"`
	Token      string `json:"Token"`
}

func (c *clientImpl) Token(apiPassword string) (int, string, error) {
	url := fmt.Sprintf("%s/token", c.baseURL)
	reqBytes, err := json.Marshal(requestSchema{APIPassword: apiPassword})
	if err != nil {
		return 0, "", fmt.Errorf("json.Marshal failed: %w", err)
	}
	resp, err := c.client.Post(url, "application/json", bytes.NewReader(reqBytes))
	if err != nil {
		return 0, "", fmt.Errorf("http.Post failed: %w", err)
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, "", fmt.Errorf("io.ReadAll failed: %w", err)
	}
	result := responseSchema{}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return 0, "", fmt.Errorf("json.Unmarshal failed: %w", err)
	}
	return result.ResultCode, result.Token, nil
}
