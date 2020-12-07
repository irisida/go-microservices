package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	req.Header = headers

	client := http.Client{}
	res, err := client.Do(req)
}
