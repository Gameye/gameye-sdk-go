package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (client GameyeClient) query(
	name string,
	arg map[string]string,
) (
	state map[string]interface{},
	err error,
) {
	query := url.Values{}
	if arg != nil {
		for key, value := range arg {
			query.Add(key, value)
		}
	}
	querystring := query.Encode()
	if len(querystring) > 0 {
		querystring = fmt.Sprintf("?%s", querystring)
	}

	href := fmt.Sprintf("%s/query/%s%s", client.config.Endpoint, name, querystring)
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
	defer res.Body.Close()
	if res.StatusCode != 200 {
		err = ErrUnexpectedStatus
		return
	}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&state)
	if err != nil {
		return
	}

	return
}
