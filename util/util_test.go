package util_test

import (
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/hsmtkk/kabu-station-api-go/api"
	"github.com/hsmtkk/kabu-station-api-go/util"
	"github.com/stretchr/testify/assert"
)

func TestNthMonth(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	apiPassword := os.Getenv("API_PASSWORD")
	client, err := api.NewLive(logger, apiPassword)
	assert.Nil(t, err)
	utility := util.New(logger, client)
	atm, err := utility.AtTheMoney()
	assert.Nil(t, err)
	fmt.Printf("%d\n", atm)
	month, err := utility.NthMonth(1)
	assert.Nil(t, err)
	fmt.Printf("%v\n", month)
}
