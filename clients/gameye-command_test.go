package clients

import (
	"testing"

	"github.com/Gameye/gameye-sdk-go/test"
	"github.com/stretchr/testify/assert"
)

func TestGameyeClient_command(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	cancelChannel := make(chan struct{})
	go test.ListenAndServeApiTestServer(
		nil, nil, cancelChannel,
	)
	defer close(cancelChannel)

	client := NewGameyeClient(GameyeClientConfig{
		Endpoint: "http://localhost:8081",
		Token:    "",
	})

	err = client.command("noop", struct{}{})
	if err != nil {
		return
	}
}
