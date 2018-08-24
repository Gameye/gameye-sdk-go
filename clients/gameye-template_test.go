package clients

import (
	"fmt"
	"strings"
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
			ctx.Response <- fmt.Sprintf(`[{"path":[],"value":%s}]`, strings.Replace(models.TemplateStateJSONMock, "\n", "", -1))
			var state *models.TemplateQueryState
			state, err = sub.NextState()
			if err != nil {
				return
			}

			assert.Equal(t, &models.TemplateStateMock, state)
		}

		return
	})
}

func TestGameyeClient_QueryTemplate(t *testing.T) {
	runInTestContext(t, func(ctx *testContext) (err error) {
		ctx.Response <- models.TemplateStateJSONMock

		var state *models.TemplateQueryState
		state, err = ctx.Client.QueryTemplate("game-x")
		if err != nil {
			return
		}
		assert.Equal(t, &models.TemplateStateMock, state)

		return
	})
}
