package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

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
	
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authorization)

	var res *http.Response
	res, err = client.httpClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 204 {
		err = ErrUnexpectedStatus
		return
	}

	return
}
