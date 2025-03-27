package api_test

import (
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/hsmtkk/kabu-station-api-go/api"
	"github.com/hsmtkk/kabu-station-api-go/api/put_or_call"
	"github.com/stretchr/testify/assert"
)

func TestSymbolnameOptionMiniGet(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	apiPassword := os.Getenv("API_PASSWORD")
	clt, err := api.NewLive(logger, apiPassword)
	assert.Nil(t, err)
	assert.NotNil(t, clt)
	result, err := clt.SymbolnameOptionMiniGet(0, 0, put_or_call.Put, 0)
	assert.Nil(t, err)
	fmt.Printf("%v\n", result)
}
