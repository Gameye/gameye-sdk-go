package clients

import (
	"net"
	"net/http"
	"testing"

	"github.com/Gameye/gameye-sdk-go/models"

	"github.com/Gameye/gameye-sdk-go/test"
	"github.com/stretchr/testify/assert"
)

func TestGameyeClient_SubscribeGame(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	mux := test.CreateAPITestServerMux(
		nil, nil,
	)

	var listener net.Listener
	listener, err = net.Listen("tcp", ":8084")
	if err != nil {
		return
	}
	defer listener.Close()
	go http.Serve(listener, mux)

	client := NewGameyeClient(GameyeClientConfig{
		Endpoint: "http://localhost:8084",
		Token:    "",
	})

	var sub *GameQuerySubscription
	sub, err = client.SubscribeGame()
	if err != nil {
		return
	}
	defer sub.Cancel()
}

func TestGameyeClient_QueryGame(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	mux := test.CreateAPITestServerMux(
		nil, nil,
	)

	var listener net.Listener
	listener, err = net.Listen("tcp", ":8085")
	if err != nil {
		return
	}
	defer listener.Close()
	go http.Serve(listener, mux)

	client := NewGameyeClient(GameyeClientConfig{
		Endpoint: "http://localhost:8085",
		Token:    "",
	})

	var state *models.GameQueryState
	state, err = client.QueryGame()
	if err != nil {
		return
	}
	assert.NotNil(t, state)
}
