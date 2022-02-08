package httpserver

import (
	"bytes"
	"net/http"
)

func doRequest(method, url string, body []byte, header map[string]string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
