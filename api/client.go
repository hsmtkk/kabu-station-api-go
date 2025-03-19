package api

import (
	"fmt"
	"net/http"

	"github.com/hsmtkk/kabu-station-api-go/api/token"
)

const PAPER_PORT = 18081
const LIVE_PORT = 18080
const BASE_PATH = "/kabusapi"

type Client interface {
}

func NewPaper(apiPassword string) (Client, error) {
	return newClient(apiPassword, PAPER_PORT)
}

func NewLive(apiPassword string) (Client, error) {
	return newClient(apiPassword, LIVE_PORT)
}

func newClient(apiPassword string, port int) (Client, error) {
	baseURL := fmt.Sprintf("http://localhost:%d%s", port, BASE_PATH)
	impl := clientImpl{baseURL: baseURL, token: ""}
	if err := impl.getToken(apiPassword); err != nil {
		return nil, err
	}
	return &impl, nil
}

type clientImpl struct {
	httpClient *http.Client
	baseURL    string
	token      string
}

func (c *clientImpl) getToken(apiPassword string) error {
	clt := token.New(c.httpClient, c.baseURL)
	_, token, err := clt.Token(apiPassword)
	if err != nil {
		return err
	}
	c.token = token
	return nil
}
