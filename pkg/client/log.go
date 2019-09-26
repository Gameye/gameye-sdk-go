package client

import (
	"context"
	"fmt"
	"io"
	"log"

	"./logs"
	"./patch"
	"github.com/Gameye/messaging-client-go/pkg/eventstream"
)

// SubscribeLogEvents adds a subscriber to log events
func SubscribeLogEvents(gameyeClient GameyeClient, matchKey string, onStateChanged func(logs.State)) (err error) {
	ctx := context.Background()
	url := fmt.Sprintf("%s/query/log", gameyeClient.config.Endpoint)

	logs.SubscribeState("client.log.internal", onStateChanged)
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
			var patches []patch.Patch
			err = decoder.Decode(&patches)

			if err == io.EOF {
				break
			} else if err != nil {
				log.Println(err)
				break
			} else if patches != nil && len(patches) > 0 {
				logs.Dispatch(&patches)
			}
		}

		logs.UnsubscribeState("client.log.internal")
	}()

	return nil
}
