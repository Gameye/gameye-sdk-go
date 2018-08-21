package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gameye/gameye-sdk-go/src/errors"
)

type GameyeClient struct {
	config     GameyeClientConfig
	httpClient *http.Client
}

func NewGameyeClient(config GameyeClientConfig) (
	client GameyeClient,
) {
	config.validate()
	httpClient := &http.Client{}

	client = GameyeClient{
		config,
		httpClient,
	}

	return
}

func (client GameyeClient) command(
	name string,
	payload interface{},
) (
	err error,
) {
	var (
		body []byte
		req  *http.Request
		res  *http.Response
	)

	url := fmt.Sprintf("%s/action/%s", client.config.Endpoint, name)
	authorization := fmt.Sprintf("Bearer %s", client.config.Token)

	body, err = json.Marshal(payload)
	if err != nil {
		return
	}

	req, err = http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return
	}
	req.Header.Add("Authorization", authorization)
	res, err = client.httpClient.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode != 202 {
		err = errors.UnexpectedStatus
		return
	}

	return
}

func (client GameyeClient) query(
	name string,
	arg interface{},
) {
	return
}

func (client GameyeClient) subscribe(
	name string,
	arg interface{},
) {

}
