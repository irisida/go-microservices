package checkyerself

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/irisida/go-microservices/06-githubapi/src/api/utils/testutils"
	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "before you wreck yerself!", verification)
}

func TestCheckYerself(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/checkyerself", nil)
	c := testutils.GetMockContext(request, response)

	WreckYerself(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "before you wreck yerself!", response.Body.String())
}
