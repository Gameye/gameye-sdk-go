package clients

import (
	"net/http"
)

type GameyeClient struct {
	config     GameyeClientConfig
	httpClient *http.Client
}

func NewGameyeClient(config GameyeClientConfig) (
	client GameyeClient,
) {
	config.validate()
	httpClient := &http.Client{}

	client = GameyeClient{
		config,
		httpClient,
	}

	return
}
