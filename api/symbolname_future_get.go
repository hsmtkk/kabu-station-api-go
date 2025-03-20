package api

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type FutureCode string

const (
	NK225      FutureCode = "NK225"
	NK225mini             = "NK225mini"
	NK225micro            = "NK225micro"
	VI                    = "VI"
)

type symbolnameFutureGetResponse struct {
	Code       int    `json:"Code"`
	Message    string `json:"Message"`
	Symbol     string `json:"Symbol"`
	SymbolName string `json:"SymbolName"`
}

func (c *clientImpl) SymbolnameFutureGet(futureCode FutureCode, derivMonth int) (string, string, error) {
	endpoint := fmt.Sprintf("%s/symbolname/future", c.baseURL)
	respBytes, err := c.getWithToken(endpoint, map[string]string{
		"FutureCode": string(futureCode),
		"DerivMonth": strconv.Itoa(derivMonth),
	})
	if err != nil {
		return "", "", fmt.Errorf("io.ReadAll failed: %w", err)
	}
	result := symbolnameFutureGetResponse{}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return "", "", fmt.Errorf("json.Unmarshal failed: %w", err)
	}
	if result.Code != 0 {
		return "", "", fmt.Errorf("got non 0 code %d: %s", result.Code, result.Message)
	}
	return result.Symbol, result.SymbolName, nil
}
