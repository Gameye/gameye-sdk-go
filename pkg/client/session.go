package client

import (
	"context"
	"fmt"
	"io"
	"log"

	"./session"
	"github.com/Gameye/messaging-client-go/pkg/eventstream"
	"github.com/Gameye/sdk-messages-go/pkg/event"
)

func SubscribeSessionEvents(gameyeClient GameyeClient, onStateChanged func(session.State)) (err error) {
	ctx := context.Background()
	url := fmt.Sprintf("%s/fetch/session", gameyeClient.config.Endpoint)

	session.SubscribeState(onStateChanged)
	decoder, err := eventstream.Create(ctx, url, nil, getStreamHeaders(gameyeClient.config))

	if err != nil {
		return err
	}

	go func() {
		for {
			var action event.UnionEvent
			decoder.Decode(&action)

			if err == io.EOF {
				break
			} else if err != nil {
				log.Println(err)
			} else {
				session.Dispatch(&action)
			}
		}

		session.UnsubscribeState(onStateChanged)
	}()

	return nil
}
