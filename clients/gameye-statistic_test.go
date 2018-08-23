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

func TestGameyeClient_SubscribeStatistic(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	patchChannel := make(chan string, 1)
	mux := test.CreateAPITestServerMux(
		`{}`, patchChannel,
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

	{
		patchChannel <- fmt.Sprintf(`[{"path":[],"value":%s}]`, strings.Replace(models.StatisticStateJSONMock, "\n", "", -1))
		var state *models.StatisticQueryState
		state, err = sub.NextState()
		if err != nil {
			return
		}

		assert.Equal(t, &models.StatisticStateMock, state)
	}
}

func TestGameyeClient_QueryStatistic(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	mux := test.CreateAPITestServerMux(
		models.StatisticStateJSONMock, nil,
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
	assert.Equal(t, &models.StatisticStateMock, state)
}
