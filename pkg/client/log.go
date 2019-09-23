package client

import (
	"context"
	"fmt"
	"io"
	"log"

	"./logs"
	"./patch"
	eventstream "github.com/Gameye/messaging-client-go/pkg/eventstream"
)

func SubscribeLogEvents(gameyeClient GameyeClient, matchKey string, onStateChanged func(logs.State)) (err error) {
	ctx := context.Background()
	url := fmt.Sprintf("%s/fetch/log", gameyeClient.config.Endpoint)

	logs.SubscribeState(onStateChanged)
	decoder, err := eventstream.Create(
		ctx,
		url,
		map[string]string{"matchKey": matchKey},
		getStreamHeaders(gameyeClient.config),
	)

	if err != nil {
		return err
	}

	go func() {
		for {
			patches := []patch.Patch{}
			decoder.Decode(&patches)

			if err == io.EOF {
				break
			} else if err != nil {
				log.Println(err)
			} else if len(patches) > 0 {
				logs.Dispatch(&patches)
			}
		}

		logs.UnsubscribeState(onStateChanged)
	}()

	return nil
}
