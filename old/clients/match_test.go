package clients

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Gameye/gameye-sdk-go/models"
)

func TestGameyeClient_SubscribeMatch(t *testing.T) {
	runInTestContext(t, func(ctx *testContext) (err error) {
		var sub *MatchQuerySubscription
		sub, err = ctx.Client.SubscribeMatch()
		if err != nil {
			return
		}
		defer sub.Cancel()

		{
			ctx.Response <- `[{"path":[],"value":{"match":{}}}]`
			var state *models.MatchQueryState
			state, err = sub.NextState()
			if err != nil {
				return
			}

			assert.NotNil(t, state)
			// assert.Equal(t, &MatchStateMock, state)
		}

		return
	})
}

func TestGameyeClient_QueryMatch(t *testing.T) {
	runInTestContext(t, func(ctx *testContext) (err error) {
		ctx.Response <- `{"match":{}}`

		var state *models.MatchQueryState
		state, err = ctx.Client.QueryMatch()
		if err != nil {
			return
		}

		assert.NotNil(t, state)
		// assert.Equal(t, &MatchStateMock, state)

		return
	})
}
