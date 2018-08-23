package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Gameye/gameye-sdk-go/clients"
	"github.com/Gameye/gameye-sdk-go/models"

	"github.com/stretchr/testify/assert"
)

func TestGameyeClient_SubscribeMatch(t *testing.T) {
	RunInContext(t, func(ctx *Context) (err error) {
		var sub *clients.MatchQuerySubscription
		sub, err = ctx.Client.SubscribeMatch()
		if err != nil {
			return
		}
		defer sub.Cancel()

		{
			ctx.Response <- fmt.Sprintf(`[{"path":[],"value":%s}]`, strings.Replace(models.MatchStateJSONMock, "\n", "", -1))
			var state *models.MatchQueryState
			state, err = sub.NextState()
			if err != nil {
				return
			}

			assert.Equal(t, &models.MatchStateMock, state)
		}

		return
	})
}

func TestGameyeClient_QueryMatch(t *testing.T) {
	RunInContext(t, func(ctx *Context) (err error) {
		ctx.Response <- models.MatchStateJSONMock

		var state *models.MatchQueryState
		state, err = ctx.Client.QueryMatch()
		if err != nil {
			return
		}
		assert.Equal(t, &models.MatchStateMock, state)

		return
	})
}
