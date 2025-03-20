package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/hsmtkk/kabu-station-api-go/api"
)

// nth is beginning from 0
func NthMonth(client api.Client, nth int) (time.Time, error) {
	futureResp, err := client.SymbolnameFutureGet(api.NK225mini, 0)
	if err != nil {
		return time.Time{}, err
	}
	firstMonth, err := parseSymbolName(futureResp.SymbolName)
	if err != nil {
		return time.Time{}, err
	}
	result := firstMonth.AddDate(0, nth, 0)
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
