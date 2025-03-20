package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", "", fmt.Errorf("url.Parse failed: %w", err)
	}
	q := u.Query()
	q.Set("FutureCode", string(futureCode))
	q.Set("DerivMonth", strconv.Itoa(derivMonth))
	u.RawQuery = q.Encode()
	queryURL := u.String()
	req, err := http.NewRequest(http.MethodGet, queryURL, nil)
	if err != nil {
		return "", "", fmt.Errorf("http.NewRequest failed: %w", err)
	}
	respBytes, err := c.invokeHTTPWithTokenHeader(req)
	if err != nil {
		return "", "", fmt.Errorf("io.ReadAll failed: %w", err)
	}
	fmt.Println(string(respBytes)) // debug
	result := symbolnameFutureGetResponse{}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return "", "", fmt.Errorf("json.Unmarshal failed: %w", err)
	}
	if result.Code != 0 {
		return "", "", fmt.Errorf("got non 0 code %d: %s", result.Code, result.Message)
	}
	return result.Symbol, result.SymbolName, nil
}
