package util_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hsmtkk/kabu-station-api-go/api"
	"github.com/hsmtkk/kabu-station-api-go/util"
	"github.com/stretchr/testify/assert"
)

func TestNthMonth(t *testing.T) {
	apiPassword := os.Getenv("API_PASSWORD")
	client, err := api.NewLive(apiPassword)
	assert.Nil(t, err)
	result, err := util.NthMonth(client, 1)
	assert.Nil(t, err)
	fmt.Printf("%v\n", result)
}
