package api

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type symbolnameOptionMiniGetResponse struct {
	Code       int    `json:"Code"`
	Message    string `json:"Message"`
	Symbol     string `json:"Symbol"`
	SymbolName string `json:"SymbolName"`
}

func (c *clientImpl) SymbolnameOptionMiniGet(derivMonth int, derivWeekly int, putOrCall PutOrCall, strikePrice int) (string, string, error) {
	endpoint := fmt.Sprintf("%s/symbolname/minioptionweekly", c.baseURL)
	respBytes, err := c.getWithToken(endpoint, map[string]string{
		"DerivMonth":  strconv.Itoa(derivMonth),
		"DerivWeekly": strconv.Itoa(derivWeekly),
		"PutOrCall":   string(putOrCall),
		"StrikePrice": strconv.Itoa(strikePrice),
	})
	if err != nil {
		return "", "", fmt.Errorf("io.ReadAll failed: %w", err)
	}
	result := symbolnameOptionMiniGetResponse{}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return "", "", fmt.Errorf("json.Unmarshal failed: %w", err)
	}
	if result.Code != 0 {
		return "", "", fmt.Errorf("got non 0 code %d: %s", result.Code, result.Message)
	}
	return result.Symbol, result.SymbolName, nil
}
