package api

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type OptionCode string

const (
	NK225op     OptionCode = "NK225op"
	NK225miniop            = "NK225miniop"
)

type PutOrCall string

const (
	Put  PutOrCall = "P"
	Call           = "C"
)

type SymbolnameOptionGetResponse struct {
	Code       int    `json:"Code"`
	Message    string `json:"Message"`
	Symbol     string `json:"Symbol"`
	SymbolName string `json:"SymbolName"`
}

func (c *clientImpl) SymbolnameOptionGet(optionCode OptionCode, derivMonth int, putOrCall PutOrCall, strikePrice int) (SymbolnameOptionGetResponse, error) {
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
	fmt.Printf("%s\n", string(respBytes)) // debug
	if result.Code != 0 {
		return result, fmt.Errorf("got non 0 code %d: %s", result.Code, result.Message)
	}
	return result, nil
}
