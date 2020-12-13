package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/irisida/go-microservices/08-concurrentWithTest/src/api/clients/restclient"
	"github.com/irisida/go-microservices/08-concurrentWithTest/src/api/domain/repositories"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidInputName(t *testing.T) {
	restclient.FlushMocks()
	request := repositories.CreateRepoRequest{}

	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "invalid repository name", err.Message())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url":"https://developer.github.com/docs"}`)),
		},
	})

	request := repositories.CreateRepoRequest{Name: "TestingValue"}

	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "Requires authentication", err.Message())
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 12345, "name":"TestingRepoName", "owner": {"login": "irisida"}}`)),
		},
	})

	request := repositories.CreateRepoRequest{Name: "TestingRepoName"}

	result, err := RepositoryService.CreateRepo(request)
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.EqualValues(t, 12345, result.ID)
	assert.EqualValues(t, "TestingRepoName", result.Name)
	assert.EqualValues(t, "irisida", result.Owner)
}
