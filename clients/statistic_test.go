package clients

import (
	"testing"

	"github.com/Gameye/gameye-sdk-go/models"
	"github.com/Gameye/gameye-sdk-go/selectors"

	"github.com/stretchr/testify/assert"
)

func TestGameyeClient_SubscribeStatistic(t *testing.T) {
	runInTestContext(t, func(ctx *testContext) (err error) {
		var sub *StatisticQuerySubscription
		sub, err = ctx.Client.SubscribeStatistic("match-y")
		if err != nil {
			return
		}
		defer sub.Cancel()

		{
			ctx.Response <- `[{"path":[],"value":{"statistic":{"start":null,"stop":null,"player":{},"startedRounds":0,"finishedRounds":0,"team":{}}}}]`
			var state *models.StatisticQueryState
			state, err = sub.NextState()
			if err != nil {
				return
			}

			assert.NotNil(t, state)
			// assert.Equal(t, &StatisticStateMock, state)
		}

		return
	})
}

func TestGameyeClient_QueryStatistic(t *testing.T) {
	runInTestContext(t, func(ctx *testContext) (err error) {
		ctx.Response <- `{"statistic":{"start":null,"stop":null,"player":{},"startedRounds":0,"finishedRounds":0,"team":{}}}`

		var state *models.StatisticQueryState
		state, err = ctx.Client.QueryStatistic("match-y")
		if err != nil {
			return
		}

		assert.NotNil(t, state)
		// assert.Equal(t, &StatisticStateMock, state)

		return
	})
}

func TestGameyeClient_killTotal(t *testing.T) {
	runInTestContext(t, func(ctx *testContext) (err error) {
		var sub *StatisticQuerySubscription
		sub, err = ctx.Client.SubscribeStatistic("match-x")
		if err != nil {
			return
		}
		defer sub.Cancel()

		{
			ctx.Response <- `[{"path":[],"value":{"statistic":{"start":1535018473000,"stop":null,"startedRounds":1,"finishedRounds":1,"player":{"3":{"connected":true,"playerKey":"3","uid":"BOT","name":"Reed","statistic":{"assist":0,"death":1,"kill":0}},"4":{"connected":true,"playerKey":"4","uid":"BOT","name":"Cory","statistic":{"assist":0,"death":0,"kill":1}},"5":{"connected":true,"playerKey":"5","uid":"BOT","name":"Quinn","statistic":{"assist":0,"death":1,"kill":0}},"6":{"connected":true,"playerKey":"6","uid":"BOT","name":"Wayne","statistic":{"assist":0,"death":0,"kill":1}},"7":{"connected":true,"playerKey":"7","uid":"BOT","name":"Tyler","statistic":{"assist":2,"death":0,"kill":0}},"8":{"connected":true,"playerKey":"8","uid":"BOT","name":"Chad","statistic":{"assist":0,"death":1,"kill":0}},"9":{"connected":true,"playerKey":"9","uid":"BOT","name":"Seth","statistic":{"assist":1,"death":0,"kill":0}},"10":{"connected":true,"playerKey":"10","uid":"BOT","name":"Joe","statistic":{"assist":0,"death":1,"kill":0}},"11":{"connected":true,"playerKey":"11","uid":"BOT","name":"Xander","statistic":{"assist":0,"death":0,"kill":3}},"12":{"connected":true,"playerKey":"12","uid":"BOT","name":"Gabe","statistic":{"assist":0,"death":1,"kill":0}}},"team":{"1":{"teamKey":"1","name":"Counter Terrorists","statistic":{"score":0},"player":{"3":true,"5":true,"8":true,"10":true,"12":true}},"2":{"teamKey":"2","name":"Terrorists","statistic":{"score":1},"player":{"4":true,"6":true,"7":true,"9":true,"11":true}}}}}}]`
			var state *models.StatisticQueryState
			state, err = sub.NextState()
			if err != nil {
				return
			}

			assert.Equal(t, 5, selectKillTotal(state))
		}

		{
			ctx.Response <- `[{"path":["statistic","startedRounds"],"value":2},{"path":["statistic","team"],"value":{"1":{"teamKey":"1","name":"Counter Terrorists","statistic":{"score":0},"player":{"3":true,"5":true,"8":true,"10":true,"12":true}},"2":{"teamKey":"2","name":"Terrorists","statistic":{"score":1},"player":{"4":true,"6":true,"7":true,"9":true,"11":true}}}}]`
			var state *models.StatisticQueryState
			state, err = sub.NextState()
			if err != nil {
				return
			}

			assert.Equal(t, 5, selectKillTotal(state))
		}

		return
	})
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
