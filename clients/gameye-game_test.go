package clients

import (
	"context"
	"fmt"
	"net/http"
	"strings"
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

	responseChannel := make(chan string, 1)
	mux := test.CreateAPITestServerMux(responseChannel)
	server := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	defer server.Shutdown(context.Background())
	go server.ListenAndServe()

	client := NewGameyeClient(GameyeClientConfig{
		Endpoint: "http://localhost:8080",
		Token:    "",
	})

	var sub *GameQuerySubscription
	sub, err = client.SubscribeGame()
	if err != nil {
		return
	}
	defer sub.Cancel()

	{
		responseChannel <- fmt.Sprintf(`[{"path":[],"value":%s}]`, strings.Replace(models.GameStateJSONMock, "\n", "", -1))
		var state *models.GameQueryState
		state, err = sub.NextState()
		if err != nil {
			return
		}

		assert.Equal(t, &models.GameStateMock, state)
	}
}

func TestGameyeClient_QueryGame(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	responseChannel := make(chan string, 1)
	mux := test.CreateAPITestServerMux(responseChannel)
	server := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	defer server.Shutdown(context.Background())
	go server.ListenAndServe()

	client := NewGameyeClient(GameyeClientConfig{
		Endpoint: "http://localhost:8080",
		Token:    "",
	})

	responseChannel <- models.GameStateJSONMock
	var state *models.GameQueryState
	state, err = client.QueryGame()
	if err != nil {
		return
	}
	assert.Equal(t, &models.GameStateMock, state)
}
