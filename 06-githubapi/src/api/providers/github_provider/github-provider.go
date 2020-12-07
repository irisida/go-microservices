package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/irisida/go-microservices/06-githubapi/src/api/clients/restclient"
	"github.com/irisida/go-microservices/06-githubapi/src/api/domain/github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, req github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	res, err := restclient.Post(urlCreateRepo, req, headers)
	if err != nil {
		log.Println("Error trying to create the %s repository.", err.Error())
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	// read the body and throw any error if cannot be read
	// defer the closing of the response body so that we
	// have a clean exit.
	bytes, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "invalid response",
		}
	}

	// blown the success code range. We have an error
	if res.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Invalid response body",
			}
		}
		// get the statuscode of the response and set in errResponse
		errResponse.StatusCode = res.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &errResponse); err != nil {
		log.Println("Error trying to unmarshal the create %s repository response.", err.Error())
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error unmarshaling the create repository response",
		}
	}

	return &github.CreateRepoResponse{}, nil
}
