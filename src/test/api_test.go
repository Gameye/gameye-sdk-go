package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListenAndServeApiTestServer(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	cancelChannel := make(chan struct{})
	go ListenAndServeApiTestServer(
		nil, nil, cancelChannel,
	)
	defer close(cancelChannel)

	var res *http.Response
	res, err = http.Get("http://localhost:8081/noop")
	if err != nil {
		return
	}
	assert.Equal(t, 204, res.StatusCode)
}
