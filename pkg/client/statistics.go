package client

import (
	"context"
	"fmt"
	"io"
	"log"

	"./patch"
	"./statistics"
	"github.com/Gameye/messaging-client-go/pkg/eventstream"
)

// SubscribeStatisticsEvents adds a subscriber to statistic events
func SubscribeStatisticsEvents(gameyeClient GameyeClient, matchKey string, onStateChanged func(statistics.State)) (err error) {
	ctx := context.Background()
	url := fmt.Sprintf("%s/fetch/statistic", gameyeClient.config.Endpoint)

	statistics.SubscribeState("client.statistics.internal", onStateChanged)
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
				statistics.Dispatch(&patches)
			}
		}

		statistics.UnsubscribeState("client.statistics.internal")
	}()

	return nil
}
