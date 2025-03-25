package api

import (
	"encoding/json"
	"fmt"
)

type unregisterAllResponse struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
}

func (c *clientImpl) UnregisterAllPut() error {
	c.logger.Debug("UnregisterAllPut")
	endpoint := fmt.Sprintf("%s/unregister/all", c.baseURL)
	respBytes, err := c.putWithToken(endpoint)
	if err != nil {
		return err
	}
	result := unregisterAllResponse{}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %w", err)
	}
	if result.Code != 0 {
		return fmt.Errorf("got non 0 code %d: %s", result.Code, result.Message)
	}
	c.logger.Debug("UnregisterAllPut", "response", result)
	return nil
}
