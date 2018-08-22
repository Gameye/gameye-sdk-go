package clients

import (
	"net/http"
)

/*
GameyeClient is a simple wrapper for the gameye api, please use
NewGameyeClient to create an instance
*/
type GameyeClient struct {
	config     GameyeClientConfig
	httpClient *http.Client
}

/*
NewGameyeClient will create a new GameyeClient
*/
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
