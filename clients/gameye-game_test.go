package clients

import (
	"fmt"
	"strings"
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
			ctx.Response <- fmt.Sprintf(`[{"path":[],"value":%s}]`, strings.Replace(models.GameStateJSONMock, "\n", "", -1))
			var state *models.GameQueryState
			state, err = sub.NextState()
			if err != nil {
				return
			}

			assert.Equal(t, &models.GameStateMock, state)
		}

		return
	})
}

func TestGameyeClient_QueryGame(t *testing.T) {
	runInTestContext(t, func(ctx *testContext) (err error) {
		ctx.Response <- models.GameStateJSONMock

		var state *models.GameQueryState
		state, err = ctx.Client.QueryGame()
		if err != nil {
			return
		}

		assert.Equal(t, &models.GameStateMock, state)
		return
	})

}
