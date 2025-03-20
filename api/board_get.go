package api

import (
	"encoding/json"
	"fmt"
)

type MarketCode string

const (
	Tokyo    MarketCode = "1"  //東証
	WholeDay            = "2"  // 日通し
	Day                 = "23" // 日中
	Night               = "24" // 夜間
)

type BoardGetResponse struct {
	Code         int     `json:"Code"`
	Message      string  `json:"Message"`
	IV           float64 `json:"IV"`
	Gamma        float64 `json:"Gamma"`
	Theta        float64 `json:"Theta"`
	Vega         float64 `json:"Vega"`
	Delta        float64 `json:"Delta"`
	Symbol       string  `json:"Symbol"`
	SymbolName   string  `json:"SymbolName"`
	CurrentPrice float64 `json:"CurrentPrice"`
}

func (c *clientImpl) BoardGet(symbolCode string, marketCode MarketCode) (BoardGetResponse, error) {
	result := BoardGetResponse{}
	symbol := symbolCode + "@" + string(marketCode)
	endpoint := fmt.Sprintf("%s/board/%s", c.baseURL, symbol)
	respBytes, err := c.getWithToken(endpoint, map[string]string{
		"symbol": symbol,
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
	return result, nil
}
