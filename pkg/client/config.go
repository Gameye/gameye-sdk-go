package client

import (
	"errors"
	"os"
	"strings"
)

/*
GameyeClientConfig configures the GameyeClient
*/
type GameyeClientConfig struct {
	Endpoint string
	Token    string
}

/*
ErrMissingConfigField occurs when a config field is missing
*/
var ErrMissingConfigField = errors.New("missing field in config")

func (config *GameyeClientConfig) validate() (
	err error,
) {
	if strings.TrimSpace(config.Endpoint) == "" {
		config.Endpoint = os.Getenv("GAMEYE_API_ENDPOINT")
	}

	if strings.TrimSpace(config.Endpoint) == "" {
		err = ErrMissingConfigField
		return
	}

	if strings.TrimSpace(config.Token) == "" {
		config.Token = os.Getenv("GAMEYE_API_TOKEN")
	}
	if strings.TrimSpace(config.Token) == "" {
		err = ErrMissingConfigField
		return
	}

	return
}
