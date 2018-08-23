package test

import (
	"testing"
)

func TestGameyeClient_query(t *testing.T) {
	RunInContext(t, func(ctx *Context) (err error) {
		// ctx.Response <- `{"a":{"b":"c"}}`

		// var actual map[string]interface{}
		// actual, err = ctx.Client.query("noop", nil)
		// if err != nil {
		// 	return
		// }

		// expected := map[string]interface{}{
		// 	"a": map[string]interface{}{
		// 		"b": "c",
		// 	},
		// }
		// assert.Equal(t, expected, actual)

		return
	})
}
