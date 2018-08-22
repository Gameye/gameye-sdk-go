package clients

import (
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
	cancelChannel := make(chan struct{})
	go test.ListenAndServeApiTestServer(
		expected, nil, cancelChannel,
	)
	defer close(cancelChannel)

	client := NewGameyeClient(GameyeClientConfig{
		Endpoint: "http://localhost:8081",
		Token:    "",
	})

	var actual map[string]interface{}
	actual, err = client.query("noop", nil)
	if err != nil {
		return
	}

	assert.Equal(t, expected, actual)
}
