package network_client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}

var client = &http.Client{
	Transport: tr,
}

// sendRequest send simple http request
func sendRequest(dataFinal *map[string]interface{}, headers *map[string]string, IP string) (*http.Response, error) {
	var req *http.Request
	var err error
	// log.Println("Send to: ", "http://"+IP+"/info")
	if dataFinal != nil {
		by := new(bytes.Buffer)
		json.NewEncoder(by).Encode(dataFinal)
		req, err = http.NewRequest("POST", "http://"+IP+"/info", by)
	} else {
		req, err = http.NewRequest("GET", "http://"+IP+"/info", nil)
	}

	if err != nil {
		return nil, err
	}

	if headers != nil {
		for k, v := range *headers {
			req.Header.Add(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
