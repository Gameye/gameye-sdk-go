package clients

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameyeClient_subscribe(t *testing.T) {
	runInTestContext(t, func(ctx *testContext) (err error) {
		var qs *querySubscription
		qs, err = ctx.Client.subscribe("noop", nil)
		if err != nil {
			return
		}
		defer qs.Cancel()

		{
			expected := map[string]interface{}{
				"a": map[string]interface{}{
					"b": "c",
				},
			}
			ctx.Response <- `[{"path":[],"value":{"a":{"b":"c"}}}]`
			var actual map[string]interface{}
			actual, err = qs.NextState()
			if err != nil {
				return
			}

			assert.Equal(t, expected, actual)
		}

		{
			expected := map[string]interface{}{
				"a": map[string]interface{}{
					"b": "d",
				},
			}
			ctx.Response <- `[{"path":["a","b"],"value":"d"}]`
			var actual map[string]interface{}
			actual, err = qs.NextState()
			if err != nil {
				return
			}

			assert.Equal(t, expected, actual)
		}

		{
			expected := map[string]interface{}{
				"a": float64(1),
			}
			ctx.Response <- `[{"path":["a"],"value":1}]`
			var actual map[string]interface{}
			actual, err = qs.NextState()
			if err != nil {
				return
			}

			assert.Equal(t, expected, actual)
		}

		return
	})
}
