package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

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

	href := fmt.Sprintf("%s/action/%s", client.config.Endpoint, name)
	authorization := fmt.Sprintf("Bearer %s", client.config.Token)

	var body []byte
	body, err = json.Marshal(payload)
	if err != nil {
		return
	}

	var req *http.Request

	req, err = http.NewRequest("POST", href, bytes.NewReader(body))
	if err != nil {
		return
	}
	req.Header.Add("Authorization", authorization)

	var res *http.Response
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
	arg map[string]string,
	state interface{},
) (
	err error,
) {
	var query url.Values
	for key, value := range arg {
		query.Add(key, value)
	}
	querystring := query.Encode()
	if len(querystring) > 0 {
		querystring = fmt.Sprintf("?%s", querystring)
	}

	href := fmt.Sprintf("%s/fetch/%s%s", client.config.Endpoint, name, querystring)
	authorization := fmt.Sprintf("Bearer %s", client.config.Token)

	var req *http.Request
	req, err = http.NewRequest("GET", href, nil)
	if err != nil {
		return
	}
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Accept", "application/json")

	var res *http.Response
	res, err = client.httpClient.Do(req)
	if err != nil {
		return
	}
	if res.StatusCode != 200 {
		err = errors.UnexpectedStatus
		return
	}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(state)
	if err != nil {
		return
	}

	return
}

func (client GameyeClient) subscribe(
	name string,
	arg map[string]string,
	state interface{},
	stateChannel chan<- interface{},
	cancelChannel <-chan struct{},
) (err error) {
	var query url.Values
	for key, value := range arg {
		query.Add(key, value)
	}
	querystring := query.Encode()
	if len(querystring) > 0 {
		querystring = fmt.Sprintf("?%s", querystring)
	}

	href := fmt.Sprintf("%s/fetch/%s%s", client.config.Endpoint, name, querystring)
	authorization := fmt.Sprintf("Bearer %s", client.config.Token)

	var req *http.Request
	req, err = http.NewRequest("GET", href, nil)
	if err != nil {
		return
	}
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Accept", "application/json")

	var res *http.Response
	res, err = client.httpClient.Do(req)
	if err != nil {
		return
	}
	if res.StatusCode != 200 {
		err = errors.UnexpectedStatus
		return
	}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(state)
	if err != nil {
		return
	}

	<-cancelChannel

	return
}
