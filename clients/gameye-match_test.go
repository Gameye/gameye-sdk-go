package clients

import (
	"context"
	"fmt"
	"strings"
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

	responseChannel := make(chan string, 1)
	server := test.CreateAPITestServer(responseChannel)
	defer server.Shutdown(context.Background())
	go server.ListenAndServe()

	client := NewGameyeClient(GameyeClientConfig{
		Endpoint: "http://localhost:8080",
		Token:    "",
	})

	var sub *MatchQuerySubscription
	sub, err = client.SubscribeMatch()
	if err != nil {
		return
	}
	defer sub.Cancel()

	{
		responseChannel <- fmt.Sprintf(`[{"path":[],"value":%s}]`, strings.Replace(models.MatchStateJSONMock, "\n", "", -1))
		var state *models.MatchQueryState
		state, err = sub.NextState()
		if err != nil {
			return
		}

		assert.Equal(t, &models.MatchStateMock, state)
	}
}

func TestGameyeClient_QueryMatch(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	responseChannel := make(chan string, 1)
	server := test.CreateAPITestServer(responseChannel)
	defer server.Shutdown(context.Background())
	go server.ListenAndServe()

	client := NewGameyeClient(GameyeClientConfig{
		Endpoint: "http://localhost:8080",
		Token:    "",
	})

	responseChannel <- models.MatchStateJSONMock
	var state *models.MatchQueryState
	state, err = client.QueryMatch()
	if err != nil {
		return
	}
	assert.Equal(t, &models.MatchStateMock, state)
}
