package test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiTestServer(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	server := CreateAPITestServer(nil)
	defer server.Shutdown(context.Background())
	go server.ListenAndServe()

	var res *http.Response
	res, err = http.Get("http://localhost:8080/noop")
	if err != nil {
		return
	}
	assert.Equal(t, 204, res.StatusCode)
}
