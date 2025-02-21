package NetworkManager

import (
	"io"
	"net/http"
)

type FetchCallReader interface {
	FetchResponse(url string) *http.Response
}

type BaseNetworkManager struct {
	url string
}

func NewBaseNetworkManager(url string) *BaseNetworkManager {
	return &BaseNetworkManager{url: url}
}

func (nm *BaseNetworkManager) FetchResponseBody(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return nil, err
	} else {
		return resp.Body, nil
	}
}
