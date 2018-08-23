package clients

import (
	"context"
	"testing"

	"github.com/Gameye/gameye-sdk-go/test"
	"github.com/stretchr/testify/assert"
)

func TestGameyeClient_subscribe(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	responseChannel := make(chan string, 1)
	server := test.CreateAPITestServer(responseChannel)
	defer server.Shutdown(context.Background())
	go server.ListenAndServe()

	client := NewGameyeClient(GameyeClientConfig{
		Endpoint: "http://localhost:8080",
		Token:    "",
	})

	var qs *querySubscription
	qs, err = client.subscribe("noop", nil)
	if err != nil {
		return
	}
	defer qs.Cancel()

	{
		expected := map[string]interface{}{
			"a": map[string]interface{}{
				"b": "c",
			},
		}
		responseChannel <- `[{"path":[],"value":{"a":{"b":"c"}}}]`
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
		responseChannel <- `[{"path":["a","b"],"value":"d"}]`
		var actual map[string]interface{}
		actual, err = qs.NextState()
		if err != nil {
			return
		}

		assert.Equal(t, expected, actual)
	}

	{
		expected := map[string]interface{}{
			"a": float64(1),
		}
		responseChannel <- `[{"path":["a"],"value":1}]`
		var actual map[string]interface{}
		actual, err = qs.NextState()
		if err != nil {
			return
		}

		assert.Equal(t, expected, actual)
	}

}
