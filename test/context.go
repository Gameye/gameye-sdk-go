package test

import (
	"context"
	"net/http"
	"testing"

	"github.com/Gameye/gameye-sdk-go/clients"
	"github.com/stretchr/testify/assert"
)

/*
Context provides a handy context for testing
*/
type Context struct {
	Server   *http.Server
	Client   clients.GameyeClient
	Response chan string
}

/*
RunInContext runs a test in a context
*/
func RunInContext(t *testing.T, job func(*Context) error) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	responseChannel := make(chan string, 1)
	defer close(responseChannel)

	server := CreateAPITestServer(responseChannel)
	go server.ListenAndServe()
	defer server.Shutdown(context.Background())

	client := clients.NewGameyeClient(clients.GameyeClientConfig{
		Token:    "",
		Endpoint: "http://localhost" + server.Addr + "",
	})

	testContext := &Context{
		Client:   client,
		Server:   server,
		Response: responseChannel,
	}

	err = job(testContext)
}
