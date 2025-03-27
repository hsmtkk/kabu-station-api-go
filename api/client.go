package api

import (
	"crypto/sha256"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/hsmtkk/kabu-station-api-go/api/market_code"
	"github.com/hsmtkk/kabu-station-api-go/api/option_code"
	"github.com/hsmtkk/kabu-station-api-go/api/put_or_call"
)

const PAPER_PORT = 18081
const LIVE_PORT = 18080
const BASE_PATH = "/kabusapi"

type Client interface {
	SymbolnameFutureGet(futureCode FutureCode, derivMonth int) (SymbolnameFutureGetResponse, error)
	SymbolnameOptionGet(optionCode option_code.OptionCode, derivMonth int, putOrCall put_or_call.PutOrCall, strikePrice int) (SymbolnameOptionGetResponse, error)
	SymbolnameOptionMiniGet(derivMonth int, derivWeekly int, putOrCall put_or_call.PutOrCall, strikePrice int) (SymbolnameOptionMiniGetResponse, error)
	BoardGet(symbolCode string, marketCode market_code.MarketCode) (BoardGetResponse, error)
	RegisterPut(symbols []SymbolMarketCode) ([]SymbolMarketCode, error)
	UnregisterAllPut() error
}

func NewPaper(logger *slog.Logger, apiPassword string) (Client, error) {
	return newClient(logger, apiPassword, PAPER_PORT)
}

func NewLive(logger *slog.Logger, apiPassword string) (Client, error) {
	return newClient(logger, apiPassword, LIVE_PORT)
}

func newClient(logger *slog.Logger, apiPassword string, port int) (Client, error) {
	logger.Debug("newClient", "apiPassword(SHA256)", hash(apiPassword), "port", port)
	baseURL := fmt.Sprintf("http://localhost:%d%s", port, BASE_PATH)
	impl := clientImpl{logger: logger, httpClient: http.DefaultClient, baseURL: baseURL, token: ""}
	_, token, err := impl.Token(apiPassword)
	if err != nil {
		return nil, err
	}
	impl.token = token
	return &impl, nil
}

type clientImpl struct {
	logger     *slog.Logger
	httpClient *http.Client
	baseURL    string
	token      string
}

func hash(s string) string {
	sum := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", sum)
}
