package client

import (
	"context"
	"fmt"
	"io"
	"log"

	messages "../messages/event"
	"./session"
	"github.com/Gameye/messaging-client-go/pkg/eventstream"
)

// SubscribeSessionEvents adds a subscriber to session events
func SubscribeSessionEvents(gameyeClient GameyeClient, onStateChanged func(session.State)) (err error) {
	ctx := context.Background()
	url := fmt.Sprintf("%s/fetch/session", gameyeClient.config.Endpoint)

	session.SubscribeState("client.session.internal", onStateChanged)
	decoder, err := eventstream.Create(ctx, url, nil, getStreamHeaders(gameyeClient.config))

	if err != nil {
		return err
	}

	go func() {
		for {
			var action messages.UnionEvent
			err = decoder.Decode(&action)

			if err == io.EOF {
				break
			} else if err != nil {
				log.Println(err)
				break
			} else {
				session.Dispatch(&action)
			}
		}

		session.UnsubscribeState("client.session.internal")
	}()

	return nil
}
