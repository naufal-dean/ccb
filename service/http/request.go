package http

import (
	"bytes"
	"log"
	"net/http"
)

func doRequest(method, url string, body []byte, header map[string]string, serviceName, targetEndpoint *string) (*http.Response, error) {
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
	if targetEndpoint != nil && serviceName != nil {
		req.Header.Set(requiringServiceHeader, *serviceName)
		req.Header.Set(currentEndpointHeader, *targetEndpoint)
		req.Header.Set(currentMethodHeader, method)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to execute request: %v\n", err)
		return nil, err
	}
	return res, nil
}
