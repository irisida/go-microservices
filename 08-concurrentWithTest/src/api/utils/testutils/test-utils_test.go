package testutils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMockContext(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "http://localhost:5566/something", nil)
	response := httptest.NewRecorder()
	request.Header = http.Header{"X-Mock": {"true"}}
	c := GetMockContext(request, response)

	assert.Nil(t, err)
	assert.EqualValues(t, http.MethodGet, c.Request.Method)
	assert.EqualValues(t, "5566", c.Request.URL.Port())
	assert.EqualValues(t, "/something", c.Request.URL.Path)
	assert.EqualValues(t, "http", c.Request.URL.Scheme)
	assert.EqualValues(t, 1, len(c.Request.Header))
	assert.EqualValues(t, "true", c.GetHeader("X-Mock"))
	assert.EqualValues(t, "true", c.GetHeader("x-mock"))
}
