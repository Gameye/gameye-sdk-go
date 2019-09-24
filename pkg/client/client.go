package client

import (
	"fmt"
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
func NewGameyeClient(config GameyeClientConfig) (client GameyeClient, err error) {
	err = config.validate()
	if err != nil {
		return GameyeClient{}, err
	}

	httpClient := &http.Client{}

	client = GameyeClient{
		config,
		httpClient,
	}

	return client, nil
}

const heartbeat = 10 * 1000

func getCommandHeaders(config GameyeClientConfig) map[string]string {
	return map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", config.Token),
	}
}

func getStreamHeaders(config GameyeClientConfig) map[string]string {
	return map[string]string{
		"Authorization":        fmt.Sprintf("Bearer %s", config.Token),
		"Accept":               "application/x-ndjson",
		"x-heartbeat-interval": fmt.Sprintf("%d", heartbeat),
	}
}
