package api

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hsmtkk/kabu-station-api-go/api/option_code"
	"github.com/hsmtkk/kabu-station-api-go/api/put_or_call"
)

type SymbolnameOptionGetResponse struct {
	Code       int    `json:"Code"`
	Message    string `json:"Message"`
	Symbol     string `json:"Symbol"`
	SymbolName string `json:"SymbolName"`
}

func (c *clientImpl) SymbolnameOptionGet(optionCode option_code.OptionCode, derivMonth int, putOrCall put_or_call.PutOrCall, strikePrice int) (SymbolnameOptionGetResponse, error) {
	c.logger.Debug("SymbolnameOptionGet", "optionCode", optionCode, "derivMonth", derivMonth, "putOrCall", putOrCall, "strikePrice", strikePrice)
	result := SymbolnameOptionGetResponse{}
	endpoint := fmt.Sprintf("%s/symbolname/option", c.baseURL)
	respBytes, err := c.getWithToken(endpoint, map[string]string{
		"OptionCode":  string(optionCode),
		"DerivMonth":  strconv.Itoa(derivMonth),
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
	c.logger.Debug("SymbolnameOptionGet", "response", result)
	return result, nil
}
