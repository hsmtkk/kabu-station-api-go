package api_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hsmtkk/kabu-station-api-go/api"
	"github.com/stretchr/testify/assert"
)

func TestBoardGet(t *testing.T) {
	apiPassword := os.Getenv("API_PASSWORD")
	clt, err := api.NewLive(apiPassword)
	assert.Nil(t, err)
	assert.NotNil(t, clt)
	result, err := clt.BoardGet("130047718", api.WholeDay)
	assert.Nil(t, err)
	fmt.Printf("%v\n", result)
}
