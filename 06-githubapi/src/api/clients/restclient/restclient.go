package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enabledMocks = false
	mocks        = make(map[string]*Mock)
)

// Mock struct to hold the mock values
type Mock struct {
	URL        string
	HTTPMethod string
	Response   *http.Response
	Err        error
}

func GetMockID(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

// StartMockups - throw a bool to as a mode controller
func StartMockups() {
	enabledMocks = true
}

// StopMockups - throw a bool reset as mode controller
func StopMockups() {
	enabledMocks = false
}

// AddMockup - load the mock values
func AddMockup(mock Mock) {
	mocks[GetMockID(mock.HTTPMethod, mock.URL)] = &mock
}

func FlushMocks() {
	mocks = make(map[string]*Mock)
}

// Post method function
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	if enabledMocks {
		// return local mock without calling external resource
		mock := mocks[GetMockID(http.MethodPost, url)]
		if mock == nil {
			return nil, errors.New("No mockup found for this request")
		}
		return mock.Response, mock.Err
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	req.Header = headers

	client := http.Client{}
	return client.Do(req)
}
