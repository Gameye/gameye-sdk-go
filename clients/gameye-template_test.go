package clients

import (
	"fmt"
	"net"
	"net/http"
	"strings"
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

	patchChannel := make(chan string, 1)
	mux := test.CreateAPITestServerMux(
		`{}`, patchChannel,
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

	{
		patchChannel <- fmt.Sprintf(`[{"path":[],"value":%s}]`, strings.Replace(models.TemplateStateJSONMock, "\n", "", -1))
		var state *models.TemplateQueryState
		state, err = sub.NextState()
		if err != nil {
			return
		}

		assert.Equal(t, &models.TemplateStateMock, state)
	}
}

func TestGameyeClient_QueryTemplate(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	mux := test.CreateAPITestServerMux(
		models.TemplateStateJSONMock, nil,
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
	assert.Equal(t, &models.TemplateStateMock, state)
}
