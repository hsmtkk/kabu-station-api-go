package api_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hsmtkk/kabu-station-api-go/api"
	"github.com/stretchr/testify/assert"
)

func TestSymbolnameFutureGet(t *testing.T) {
	apiPassword := os.Getenv("API_PASSWORD")
	clt, err := api.NewLive(apiPassword)
	assert.Nil(t, err)
	assert.NotNil(t, clt)
	result, err := clt.SymbolnameFutureGet(api.NK225, 0)
	assert.Nil(t, err)
	fmt.Printf("%v\n", result)
}
