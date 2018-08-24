package clients

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
testContext provides a handy context for testing
*/
type testContext struct {
	Server   *http.Server
	Client   GameyeClient
	Response chan string
}

/*
runInTestContext runs a test in a context
*/
func runInTestContext(t *testing.T, job func(*testContext) error) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	responseChannel := make(chan string, 1)
	defer close(responseChannel)

	server := createAPITestServer(responseChannel)
	go server.ListenAndServe()
	defer server.Shutdown(context.Background())

	client := NewGameyeClient(GameyeClientConfig{
		Token:    "",
		Endpoint: "http://localhost" + server.Addr + "",
	})

	testContext := &testContext{
		Client:   client,
		Server:   server,
		Response: responseChannel,
	}

	err = job(testContext)
}
