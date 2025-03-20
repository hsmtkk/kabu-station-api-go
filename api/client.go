package api

import (
	"fmt"
	"net/http"
)

const PAPER_PORT = 18081
const LIVE_PORT = 18080
const BASE_PATH = "/kabusapi"

type Client interface {
	SymbolnameFutureGet(futureCode FutureCode, derivMonth int) (string, string, error)
}

func NewPaper(apiPassword string) (Client, error) {
	return newClient(apiPassword, PAPER_PORT)
}

func NewLive(apiPassword string) (Client, error) {
	return newClient(apiPassword, LIVE_PORT)
}

func newClient(apiPassword string, port int) (Client, error) {
	baseURL := fmt.Sprintf("http://localhost:%d%s", port, BASE_PATH)
	impl := clientImpl{httpClient: http.DefaultClient, baseURL: baseURL, token: ""}
	_, token, err := impl.Token(apiPassword)
	if err != nil {
		return nil, err
	}
	impl.token = token
	return &impl, nil
}

type clientImpl struct {
	httpClient *http.Client
	baseURL    string
	token      string
}
