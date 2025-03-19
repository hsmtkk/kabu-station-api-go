package token_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/hsmtkk/kabu-station-api-go/api/token"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	apiPassword := os.Getenv("API_PASSWORD")
	clt := token.New(http.DefaultClient, "http://localhost:18080/kabusapi")
	resultCode, token, err := clt.Token(apiPassword)
	assert.Nil(t, err)
	assert.Equal(t, resultCode, 0)
	assert.NotEmpty(t, token)
}
