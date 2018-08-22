package clients

import (
	"net"
	"net/http"
	"testing"

	"github.com/Gameye/gameye-sdk-go/test"
	"github.com/stretchr/testify/assert"
)

func TestGameyeClient_query(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	expected := map[string]interface{}{
		"a": map[string]interface{}{
			"b": "c",
		},
	}
	mux := test.CreateApiTestServerMux(
		expected, nil,
	)

	var listener net.Listener
	listener, err = net.Listen("tcp", ":8082")
	if err != nil {
		return
	}
	defer listener.Close()
	go http.Serve(listener, mux)

	client := NewGameyeClient(GameyeClientConfig{
		Endpoint: "http://localhost:8082",
		Token:    "",
	})

	var actual map[string]interface{}
	actual, err = client.query("noop", nil)
	if err != nil {
		return
	}

	assert.Equal(t, expected, actual)
}
