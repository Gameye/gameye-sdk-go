package client

import (
	"fmt"

	command "github.com/Gameye/messaging-client-go/pkg/command"
	messages "github.com/Gameye/sdk-messages-go/pkg/command"
)

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
			EndCallbackUrl: endCallbackURL,
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
