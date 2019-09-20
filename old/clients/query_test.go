package clients

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameyeClient_query(t *testing.T) {
	runInTestContext(t, func(ctx *testContext) (err error) {
		ctx.Response <- `{"a":{"b":"c"}}`

		var actual map[string]interface{}
		actual, err = ctx.Client.query("noop", nil)
		if err != nil {
			return
		}

		expected := map[string]interface{}{
			"a": map[string]interface{}{
				"b": "c",
			},
		}
		assert.Equal(t, expected, actual)

		return
	})
}
