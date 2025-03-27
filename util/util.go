package util

import (
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/hsmtkk/kabu-station-api-go/api"
	"github.com/hsmtkk/kabu-station-api-go/api/market_code"
)

type Utility interface {
	AtTheMoney() (int, error)
	NthMonth(nthMonth int) (time.Time, error)
}

type utilityImpl struct {
	logger    *slog.Logger
	apiClient api.Client
}

func New(logger *slog.Logger, apiClient api.Client) Utility {
	return &utilityImpl{logger, apiClient}
}

// nth is beginning from 0
func (u *utilityImpl) NthMonth(nthMonth int) (time.Time, error) {
	u.logger.Debug("NthMonth", "nthMonth", nthMonth)
	futureResp, err := u.apiClient.SymbolnameFutureGet(api.NK225mini, 0)
	if err != nil {
		return time.Time{}, err
	}
	firstMonth, err := parseSymbolName(futureResp.SymbolName)
	if err != nil {
		return time.Time{}, err
	}
	result := firstMonth.AddDate(0, nthMonth, 0)
	u.logger.Debug("NthMonth", "response", result)
	return result, nil
}

func parseSymbolName(symbolname string) (time.Time, error) {
	elems := strings.Split(symbolname, " ")
	if len(elems) != 2 {
		return time.Time{}, fmt.Errorf("failed to parse symbol name: %s", symbolname)
	}
	parsed, err := time.Parse("06/01", elems[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse year and month: %s", elems[1])
	}
	return parsed, nil
}

// ATMオプションの権利行使価格 先物ミニの価格を125の倍数に丸める
func (u *utilityImpl) AtTheMoney() (int, error) {
	u.logger.Debug("AtTheMoney")
	// 先物ミニのシンボル
	futureResp, err := u.apiClient.SymbolnameFutureGet(api.NK225mini, 0)
	if err != nil {
		return 0, err
	}
	// 先物ミニの価格
	boardResp, err := u.apiClient.BoardGet(futureResp.Symbol, market_code.WholeDay)
	if err != nil {
		return 0, err
	}
	price := boardResp.CurrentPrice
	roundPrice := int(price/125) * 125
	u.logger.Debug("AtTheMoney", "response", roundPrice)
	return roundPrice, nil
}
