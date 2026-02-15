package api

import (
	"io"
	"net/http"
)

type Api struct {
	BridgeUrl      string
	BridgeLocalUrl string
	UjmUrl         string
	Apikey         string
	Client         http.Client
	ServerUrl      string
	DiscordWebhook string
}

// UjmGetWithBody Sending a body inside a get request is a bad practice, we should change that on ujm side.
func (api *Api) UjmGetWithBody(url string, body io.Reader) (resp *http.Response, err error) {
	request, err := http.NewRequest(http.MethodGet, url, body)

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	return api.Client.Do(request)
}
