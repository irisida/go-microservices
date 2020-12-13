package repositories

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/irisida/go-microservices/07-concurrentapi/src/api/clients/restclient"
	"github.com/irisida/go-microservices/07-concurrentapi/src/api/domain/repositories"
	"github.com/irisida/go-microservices/07-concurrentapi/src/api/utils/errors"
	"github.com/irisida/go-microservices/07-concurrentapi/src/api/utils/testutils"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidJsonRequest(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/respositories", strings.NewReader(``))
	res := httptest.NewRecorder()
	c := testutils.GetMockContext(request, res)

	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, res.Code)

	apiErr, err := errors.NewAPIErrorFromBytes(res.Body.Bytes())

	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid json body", apiErr.Message())

}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url":"https://developer.github.com/docs"}`)),
		},
	})

	request, _ := http.NewRequest(http.MethodPost, "/respositories", strings.NewReader(`{"name": "testing"}`))
	res := httptest.NewRecorder()
	c := testutils.GetMockContext(request, res)

	CreateRepo(c)

	assert.EqualValues(t, http.StatusUnauthorized, res.Code)

	apiErr, err := errors.NewAPIErrorFromBytes(res.Body.Bytes())

	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusUnauthorized, apiErr.Status())
	assert.EqualValues(t, "Requires authentication", apiErr.Message())
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 12345}`)),
		},
	})

	request, _ := http.NewRequest(http.MethodPost, "/respositories", strings.NewReader(`{"name": "testing"}`))
	res := httptest.NewRecorder()
	c := testutils.GetMockContext(request, res)

	CreateRepo(c)

	assert.EqualValues(t, http.StatusCreated, res.Code)

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(res.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, 12345, result.ID)
	assert.EqualValues(t, "", result.Owner)
}
