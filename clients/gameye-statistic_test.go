package clients

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"testing"

	"github.com/Gameye/gameye-sdk-go/models"
	"github.com/Gameye/gameye-sdk-go/selectors"

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

func TestGameyeClient_killTotal(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	patchChannel := make(chan string, 1)
	mux := test.CreateAPITestServerMux(
		`{}`, patchChannel,
	)

	var listener net.Listener
	listener, err = net.Listen("tcp", ":8100")
	if err != nil {
		return
	}
	defer listener.Close()
	go http.Serve(listener, mux)

	client := NewGameyeClient(GameyeClientConfig{
		Endpoint: "http://localhost:8100",
		Token:    "",
	})

	var sub *StatisticQuerySubscription
	sub, err = client.SubscribeStatistic("match-x")
	if err != nil {
		return
	}
	defer sub.Cancel()

	{
		patchChannel <- `[{"path":[],"value":{"statistic":{"start":1535018473000,"stop":null,"startedRounds":1,"finishedRounds":1,"player":{"3":{"connected":true,"playerKey":"3","uid":"BOT","name":"Reed","statistic":{"assist":0,"death":1,"kill":0}},"4":{"connected":true,"playerKey":"4","uid":"BOT","name":"Cory","statistic":{"assist":0,"death":0,"kill":1}},"5":{"connected":true,"playerKey":"5","uid":"BOT","name":"Quinn","statistic":{"assist":0,"death":1,"kill":0}},"6":{"connected":true,"playerKey":"6","uid":"BOT","name":"Wayne","statistic":{"assist":0,"death":0,"kill":1}},"7":{"connected":true,"playerKey":"7","uid":"BOT","name":"Tyler","statistic":{"assist":2,"death":0,"kill":0}},"8":{"connected":true,"playerKey":"8","uid":"BOT","name":"Chad","statistic":{"assist":0,"death":1,"kill":0}},"9":{"connected":true,"playerKey":"9","uid":"BOT","name":"Seth","statistic":{"assist":1,"death":0,"kill":0}},"10":{"connected":true,"playerKey":"10","uid":"BOT","name":"Joe","statistic":{"assist":0,"death":1,"kill":0}},"11":{"connected":true,"playerKey":"11","uid":"BOT","name":"Xander","statistic":{"assist":0,"death":0,"kill":3}},"12":{"connected":true,"playerKey":"12","uid":"BOT","name":"Gabe","statistic":{"assist":0,"death":1,"kill":0}}},"team":{"1":{"teamKey":"1","name":"Counter Terrorists","statistic":{"score":0},"player":{"3":true,"5":true,"8":true,"10":true,"12":true}},"2":{"teamKey":"2","name":"Terrorists","statistic":{"score":1},"player":{"4":true,"6":true,"7":true,"9":true,"11":true}}}}}}]`
		var state *models.StatisticQueryState
		state, err = sub.NextState()
		if err != nil {
			return
		}

		assert.Equal(t, 5, selectKillTotal(state))
	}

	{
		patchChannel <- `[{"path":["statistic","startedRounds"],"value":2},{"path":["statistic","team"],"value":{"1":{"teamKey":"1","name":"Counter Terrorists","statistic":{"score":0},"player":{"3":true,"5":true,"8":true,"10":true,"12":true}},"2":{"teamKey":"2","name":"Terrorists","statistic":{"score":1},"player":{"4":true,"6":true,"7":true,"9":true,"11":true}}}}]`
		var state *models.StatisticQueryState
		state, err = sub.NextState()
		if err != nil {
			return
		}

		assert.Equal(t, 5, selectKillTotal(state))
	}
}

func selectKillTotal(state *models.StatisticQueryState) (
	total int,
) {
	playerList := selectors.SelectPlayerList(state)
	for _, playerItem := range playerList {
		if value, exists := playerItem.Statistic["kill"]; exists {
			total += value
		}
	}
	return
}
