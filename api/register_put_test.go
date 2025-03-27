package api_test

import (
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/hsmtkk/kabu-station-api-go/api"
	"github.com/hsmtkk/kabu-station-api-go/api/exchange_code"
	"github.com/stretchr/testify/assert"
)

func TestRegisterPut(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	apiPassword := os.Getenv("API_PASSWORD")
	clt, err := api.NewLive(logger, apiPassword)
	assert.Nil(t, err)
	assert.NotNil(t, clt)
	result, err := clt.RegisterPut([]api.SymbolMarketCode{{"130047718", exchange_code.WholeDay}})
	assert.Nil(t, err)
	assert.NotNil(t, result)
	fmt.Printf("%v\n", result)
}
