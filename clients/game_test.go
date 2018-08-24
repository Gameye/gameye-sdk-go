package clients

import (
	"testing"

	"github.com/Gameye/gameye-sdk-go/models"
	"github.com/stretchr/testify/assert"
)

func TestGameyeClient_SubscribeGame(t *testing.T) {
	runInTestContext(t, func(ctx *testContext) (err error) {
		var sub *GameQuerySubscription
		sub, err = ctx.Client.SubscribeGame()
		if err != nil {
			return
		}
		defer sub.Cancel()

		{
			ctx.Response <- `[{"path":[],"value":{"game":{},"location":{}}}]`
			var state *models.GameQueryState
			state, err = sub.NextState()
			if err != nil {
				return
			}

			assert.NotNil(t, state)
			// assert.Equal(t, &GameStateMock, state)
		}

		return
	})
}

func TestGameyeClient_QueryGame(t *testing.T) {
	runInTestContext(t, func(ctx *testContext) (err error) {
		ctx.Response <- `{"game":{},"location":{}}`

		var state *models.GameQueryState
		state, err = ctx.Client.QueryGame()
		if err != nil {
			return
		}

		assert.NotNil(t, state)
		// assert.Equal(t, &GameStateMock, state)

		return
	})

}
