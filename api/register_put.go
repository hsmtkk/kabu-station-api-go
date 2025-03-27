package api

import (
	"encoding/json"
	"fmt"

	"github.com/hsmtkk/kabu-station-api-go/api/exchange_code"
)

type registerPutRequest struct {
	Symbols []SymbolMarketCode `json:"Symbols"`
}

type registerPutResponse struct {
	Code       int                `json:"Code"`
	Message    string             `json:"Message"`
	RegistList []SymbolMarketCode `json:"RegistList"`
}

type SymbolMarketCode struct {
	Symbol   string                     `json:"Symbol"`
	Exchange exchange_code.ExchangeCode `json:"Exchange"`
}

func (c *clientImpl) RegisterPut(symbols []SymbolMarketCode) ([]SymbolMarketCode, error) {
	c.logger.Debug("RegisterPut begin")
	endpoint := fmt.Sprintf("%s/register", c.baseURL)
	req := registerPutRequest{
		Symbols: symbols,
	}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal failed: %w", err)
	}
	fmt.Printf("%s\n", string(reqBytes)) // debug
	respBytes, err := c.putWithToken(endpoint, reqBytes)
	if err != nil {
		return nil, err
	}
	result := registerPutResponse{}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal failed: %w", err)
	}
	if result.Code != 0 {
		return nil, fmt.Errorf("got non 0 code %d: %s", result.Code, result.Message)
	}
	c.logger.Debug("RegisterPut end", "response", result)
	return result.RegistList, nil
}
