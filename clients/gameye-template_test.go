package clients

import (
	"net"
	"net/http"
	"testing"

	"github.com/Gameye/gameye-sdk-go/models"

	"github.com/Gameye/gameye-sdk-go/test"
	"github.com/stretchr/testify/assert"
)

func TestGameyeClient_SubscribeTemplate(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	mux := test.CreateAPITestServerMux(
		nil, nil,
	)

	var listener net.Listener
	listener, err = net.Listen("tcp", ":8090")
	if err != nil {
		return
	}
	defer listener.Close()
	go http.Serve(listener, mux)

	client := NewGameyeClient(GameyeClientConfig{
		Endpoint: "http://localhost:8090",
		Token:    "",
	})

	var sub *TemplateQuerySubscription
	sub, err = client.SubscribeTemplate("game-x")
	if err != nil {
		return
	}
	defer sub.Cancel()
}

func TestGameyeClient_QueryTemplate(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	mux := test.CreateAPITestServerMux(
		nil, nil,
	)

	var listener net.Listener
	listener, err = net.Listen("tcp", ":8091")
	if err != nil {
		return
	}
	defer listener.Close()
	go http.Serve(listener, mux)

	client := NewGameyeClient(GameyeClientConfig{
		Endpoint: "http://localhost:8091",
		Token:    "",
	})

	var state *models.TemplateQueryState
	state, err = client.QueryTemplate("game-x")
	if err != nil {
		return
	}
	assert.NotNil(t, state)
}
