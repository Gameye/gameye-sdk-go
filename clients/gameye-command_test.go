package clients

import (
	"context"
	"testing"

	"github.com/Gameye/gameye-sdk-go/test"
	"github.com/stretchr/testify/assert"
)

func TestGameyeClient_command(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	server := test.CreateAPITestServer(nil)
	defer server.Shutdown(context.Background())
	go server.ListenAndServe()

	client := NewGameyeClient(GameyeClientConfig{
		Endpoint: "http://localhost:8080",
		Token:    "",
	})

	err = client.command("noop", struct{}{})
	if err != nil {
		return
	}
}
