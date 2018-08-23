package clients

import (
	"context"
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

	responseChannel := make(chan string, 1)
	mux := test.CreateAPITestServerMux(responseChannel)
	server := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	defer server.Shutdown(context.Background())
	go server.ListenAndServe()

	client := NewGameyeClient(GameyeClientConfig{
		Endpoint: "http://localhost:8080",
		Token:    "",
	})

	responseChannel <- `{"a":{"b":"c"}}`
	var actual map[string]interface{}
	actual, err = client.query("noop", nil)
	if err != nil {
		return
	}

	expected := map[string]interface{}{
		"a": map[string]interface{}{
			"b": "c",
		},
	}

	assert.Equal(t, expected, actual)
}
