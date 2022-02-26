package http

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
)

func doRequest(method, url string, body []byte, header map[string]string, serviceName *string, reqUrl *url.URL) (*http.Response, error) {
	cb := getCircuitBreaker(reqUrl.Host)

	res, err := cb.Execute(func() (interface{}, error) {
		return doRequestHelper(method, url, body, header, serviceName, reqUrl)
	})
	if err != nil {
		return nil, err
	}
	return res.(*http.Response), nil
}

func doRequestHelper(method, url string, body []byte, header map[string]string, serviceName *string, reqUrl *url.URL) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Failed to create new request: %v\n", err)
		return nil, err
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}
	// Inject header metadata
	if reqUrl != nil && serviceName != nil {
		req.Header.Set(requiringServiceHeader, *serviceName)
		req.Header.Set(currentEndpointHeader, reqUrl.Path)
		req.Header.Set(currentMethodHeader, method)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to execute request: %v\n", err)
		return nil, err
	}
	return res, nil
}
