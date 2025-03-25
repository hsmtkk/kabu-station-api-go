package api

import (
	"encoding/json"
	"fmt"
)

type requestSchema struct {
	APIPassword string `json:"APIPassword"`
}

type responseSchema struct {
	ResultCode int    `json:"ResultCode"`
	Token      string `json:"Token"`
}

func (c *clientImpl) Token(apiPassword string) (int, string, error) {
	c.logger.Debug("Token", "apiPassword(SHA256)", hash(apiPassword))
	url := fmt.Sprintf("%s/token", c.baseURL)
	reqBytes, err := json.Marshal(requestSchema{APIPassword: apiPassword})
	if err != nil {
		return 0, "", fmt.Errorf("json.Marshal failed: %w", err)
	}
	respBytes, err := c.post(url, reqBytes)
	if err != nil {
		return 0, "", err
	}
	result := responseSchema{}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return 0, "", fmt.Errorf("json.Unmarshal failed: %w", err)
	}
	c.logger.Debug("Token", "response", result)
	return result.ResultCode, result.Token, nil
}
