package clients

import (
	"testing"

	"github.com/Gameye/gameye-sdk-go/models"
	"github.com/stretchr/testify/assert"
)

func TestGameyeClient_SubscribeTemplate(t *testing.T) {
	runInTestContext(t, func(ctx *testContext) (err error) {

		var sub *TemplateQuerySubscription
		sub, err = ctx.Client.SubscribeTemplate("game-x")
		if err != nil {
			return
		}
		defer sub.Cancel()

		{
			ctx.Response <- `[{"path":[],"value":{"template":{}}}]`
			var state *models.TemplateQueryState
			state, err = sub.NextState()
			if err != nil {
				return
			}

			assert.NotNil(t, state)
			// assert.Equal(t, &TemplateStateMock, state)
		}

		return
	})
}

func TestGameyeClient_QueryTemplate(t *testing.T) {
	runInTestContext(t, func(ctx *testContext) (err error) {
		ctx.Response <- `{"template":{}}`

		var state *models.TemplateQueryState
		state, err = ctx.Client.QueryTemplate("game-x")
		if err != nil {
			return
		}

		assert.NotNil(t, state)
		// assert.Equal(t, &TemplateStateMock, state)

		return
	})
}
