package api

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hsmtkk/kabu-station-api-go/api/put_or_call"
)

type SymbolnameOptionMiniGetResponse struct {
	Code       int    `json:"Code"`
	Message    string `json:"Message"`
	Symbol     string `json:"Symbol"`
	SymbolName string `json:"SymbolName"`
}

func (c *clientImpl) SymbolnameOptionMiniGet(derivMonth int, derivWeekly int, putOrCall put_or_call.PutOrCall, strikePrice int) (SymbolnameOptionMiniGetResponse, error) {
	c.logger.Debug("SymbolnameOptionMiniGet", "derivMonth", derivMonth, "derivWeekly", derivWeekly, "putOrCall", putOrCall, "strikePrice", strikePrice)
	result := SymbolnameOptionMiniGetResponse{}
	endpoint := fmt.Sprintf("%s/symbolname/minioptionweekly", c.baseURL)
	respBytes, err := c.getWithToken(endpoint, map[string]string{
		"DerivMonth":  strconv.Itoa(derivMonth),
		"DerivWeekly": strconv.Itoa(derivWeekly),
		"PutOrCall":   string(putOrCall),
		"StrikePrice": strconv.Itoa(strikePrice),
	})
	if err != nil {
		return result, fmt.Errorf("io.ReadAll failed: %w", err)
	}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return result, fmt.Errorf("json.Unmarshal failed: %w", err)
	}
	if result.Code != 0 {
		return result, fmt.Errorf("got non 0 code %d: %s", result.Code, result.Message)
	}
	c.logger.Debug("SymbolnameOptionMiniGet", "response", result)
	return result, nil
}
