package api_test

import (
	"log/slog"
	"os"
	"testing"

	"github.com/hsmtkk/kabu-station-api-go/api"
	"github.com/stretchr/testify/assert"
)

func TestUnregisterAllPut(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	apiPassword := os.Getenv("API_PASSWORD")
	clt, err := api.NewLive(logger, apiPassword)
	assert.Nil(t, err)
	assert.NotNil(t, clt)
	err = clt.UnregisterAllPut()
	assert.Nil(t, err)
}
