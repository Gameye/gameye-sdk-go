package clients

import (
	"net"
	"net/http"
	"testing"

	"github.com/Gameye/gameye-sdk-go/models"

	"github.com/Gameye/gameye-sdk-go/test"
	"github.com/stretchr/testify/assert"
)

func TestGameyeClient_SubscribeStatistic(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	mux := test.CreateAPITestServerMux(
		nil, nil,
	)

	var listener net.Listener
	listener, err = net.Listen("tcp", ":8088")
	if err != nil {
		return
	}
	defer listener.Close()
	go http.Serve(listener, mux)

	client := NewGameyeClient(GameyeClientConfig{
		Endpoint: "http://localhost:8088",
		Token:    "",
	})

	var sub *StatisticQuerySubscription
	sub, err = client.SubscribeStatistic("match-y")
	if err != nil {
		return
	}
	defer sub.Cancel()
}

func TestGameyeClient_QueryStatistic(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	mux := test.CreateAPITestServerMux(
		nil, nil,
	)

	var listener net.Listener
	listener, err = net.Listen("tcp", ":8089")
	if err != nil {
		return
	}
	defer listener.Close()
	go http.Serve(listener, mux)

	client := NewGameyeClient(GameyeClientConfig{
		Endpoint: "http://localhost:8089",
		Token:    "",
	})

	var state *models.StatisticQueryState
	state, err = client.QueryStatistic("match-y")
	if err != nil {
		return
	}
	assert.NotNil(t, state)
}
