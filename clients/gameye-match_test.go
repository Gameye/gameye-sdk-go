package clients

import (
	"net"
	"net/http"
	"testing"

	"github.com/Gameye/gameye-sdk-go/models"

	"github.com/Gameye/gameye-sdk-go/test"
	"github.com/stretchr/testify/assert"
)

func TestGameyeClient_SubscribeMatch(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	mux := test.CreateAPITestServerMux(
		nil, nil,
	)

	var listener net.Listener
	listener, err = net.Listen("tcp", ":8086")
	if err != nil {
		return
	}
	defer listener.Close()
	go http.Serve(listener, mux)

	client := NewGameyeClient(GameyeClientConfig{
		Endpoint: "http://localhost:8086",
		Token:    "",
	})

	var sub *MatchQuerySubscription
	sub, err = client.SubscribeMatch()
	if err != nil {
		return
	}
	defer sub.Cancel()
}

func TestGameyeClient_QueryMatch(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	mux := test.CreateAPITestServerMux(
		nil, nil,
	)

	var listener net.Listener
	listener, err = net.Listen("tcp", ":8087")
	if err != nil {
		return
	}
	defer listener.Close()
	go http.Serve(listener, mux)

	client := NewGameyeClient(GameyeClientConfig{
		Endpoint: "http://localhost:8087",
		Token:    "",
	})

	var state *models.MatchQueryState
	state, err = client.QueryMatch()
	if err != nil {
		return
	}
	assert.NotNil(t, state)
}