package client

import (
	"fmt"

	messages "../messages/command"
	"github.com/Gameye/messaging-client-go/pkg/command"
)

// StartMatch starts a new match with the given paramters
func StartMatch(client GameyeClient,
	matchKey string,
	gameKey string,
	locationKeys []string,
	templateKey string,
	config map[string]interface{},
	endCallbackURL string,
) (err error) {

	action := messages.StartMatchCommand{
		Type: "start-match",
		Payload: messages.StartMatchCommandPayload{
			MatchKey:       matchKey,
			GameKey:        gameKey,
			LocationKeys:   locationKeys,
			TemplateKey:    templateKey,
			Config:         config,
			EndCallbackURL: endCallbackURL,
		},
	}

	url := fmt.Sprintf("%s/action/%s", client.config.Endpoint, action.Type)
	err = command.Invoke(
		url,
		action.Payload,
		getCommandHeaders(client.config),
	)
	return err
}

// StopMatch ends a match with the given matchKey
func StopMatch(client GameyeClient, matchKey string) (err error) {

	action := messages.StopMatchCommand{
		Type: "stop-match",
		Payload: messages.StopMatchCommandPayload{
			MatchKey: matchKey,
		},
	}

	url := fmt.Sprintf("%s/action/%s", client.config.Endpoint, action.Type)
	err = command.Invoke(
		url,
		action.Payload,
		getCommandHeaders(client.config),
	)
	return err
}
