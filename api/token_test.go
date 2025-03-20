package api_test

import (
	"os"
	"testing"

	"github.com/hsmtkk/kabu-station-api-go/api"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	apiPassword := os.Getenv("API_PASSWORD")
	clt, err := api.NewLive(apiPassword)
	assert.Nil(t, err)
	assert.NotNil(t, clt)
}
