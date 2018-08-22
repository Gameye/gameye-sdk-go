package clients

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

func (client GameyeClient) subscribe(
	name string,
	arg map[string]string,
) (
	qs QuerySubscription,
	err error,
) {
	var query url.Values
	if arg != nil {
		for key, value := range arg {
			query.Add(key, value)
		}
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
	req.Header.Add("Accept", "application/x-ndjson")

	var ctx context.Context
	var cancelRequest context.CancelFunc
	ctx, cancelRequest = context.WithCancel(context.Background())
	defer func() {
		if err != nil {
			cancelRequest()
		}
	}()
	req = req.WithContext(ctx)

	var res *http.Response
	res, err = client.httpClient.Do(req)
	if err != nil {
		return
	}
	if res.StatusCode != 200 {
		err = ErrUnexpectedStatus
		return
	}

	reader := bufio.NewReader(res.Body)
	qs = NewQuerySubscription(
		reader,
		cancelRequest,
	)

	return
}
