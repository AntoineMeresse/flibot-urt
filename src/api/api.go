package api

import (
	"net/http"
)

type Api struct {
	BridgeUrl string
	BridgeLocalUrl string
	UjmUrl string
	Apikey string
	Client http.Client
}

