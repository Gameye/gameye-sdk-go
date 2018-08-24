package clients

import "testing"

func TestGameyeClient_command(t *testing.T) {
	runInTestContext(t, func(ctx *testContext) (err error) {
		err = ctx.Client.command("noop", struct{}{})
		if err != nil {
			return
		}
		return
	})
}
