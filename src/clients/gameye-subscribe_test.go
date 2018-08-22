package clients

import (
	"testing"

	"github.com/Gameye/gameye-sdk-go/src/test"
	"github.com/stretchr/testify/assert"
)

func TestGameyeClient_subscribe(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	cancelChannel := make(chan struct{})
	patchChannel := make(chan test.QueryPatch, 1)
	go test.ListenAndServeApiTestServer(
		nil, patchChannel, cancelChannel,
	)
	defer close(cancelChannel)

	client := NewGameyeClient(GameyeClientConfig{
		Endpoint: "http://localhost:8081",
		Token:    "",
	})

	var qs QuerySubscription
	qs, err = client.subscribe("noop", nil)
	if err != nil {
		return
	}

	{
		expected := map[string]interface{}{
			"a": map[string]interface{}{
				"b": "c",
			},
		}
		patchChannel <- test.QueryPatch{
			Path:  []string{},
			Value: expected,
		}
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
		patchChannel <- test.QueryPatch{
			Path:  []string{"a", "b"},
			Value: "d",
		}
		var actual map[string]interface{}
		actual, err = qs.NextState()
		if err != nil {
			return
		}

		assert.Equal(t, expected, actual)
	}

	qs.Cancel()
}
