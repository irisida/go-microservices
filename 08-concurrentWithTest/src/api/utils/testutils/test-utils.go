package testutils

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// GetMockContext - return the context of a Mock for testing
func GetMockContext(request *http.Request, response *httptest.ResponseRecorder) *gin.Context {
	c, _ := gin.CreateTestContext(response)
	c.Request = request
	return c
}
