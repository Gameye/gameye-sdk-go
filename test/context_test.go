package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunInContext(t *testing.T) {
	RunInContext(t, func(ctx *Context) (err error) {
		var res *http.Response
		res, err = http.Get("http://localhost" + ctx.Server.Addr + "/noop")
		if err != nil {
			return
		}
		assert.Equal(t, 204, res.StatusCode)
		return
	})
}
