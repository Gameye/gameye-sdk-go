package test

import (
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiTestServer(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	mux := CreateAPITestServerMux(
		`{}`, nil,
	)

	var listener net.Listener
	listener, err = net.Listen("tcp", ":8080")
	if err != nil {
		return
	}
	defer listener.Close()
	go http.Serve(listener, mux)

	var res *http.Response
	res, err = http.Get("http://localhost:8080/noop")
	if err != nil {
		return
	}
	assert.Equal(t, 204, res.StatusCode)
}
