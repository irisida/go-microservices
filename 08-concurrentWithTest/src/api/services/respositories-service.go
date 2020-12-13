package services

import (
	"net/http"
	"sync"

	"github.com/irisida/go-microservices/08-concurrentWithTest/src/api/config"
	"github.com/irisida/go-microservices/08-concurrentWithTest/src/api/domain/github"
	"github.com/irisida/go-microservices/08-concurrentWithTest/src/api/domain/repositories"
	"github.com/irisida/go-microservices/08-concurrentWithTest/src/api/providers/githubprovider"
	"github.com/irisida/go-microservices/08-concurrentWithTest/src/api/utils/errors"
)

type repoService struct{}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError)
	CreateMultipleRepos(request []repositories.CreateRepoRequest) (repositories.CreateMultipleReposResponse, errors.APIError)
}

var (
	// RepositoryService is the exported service
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	res, err := githubprovider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewAPIError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		ID:    res.Id,
		Name:  res.Name,
		Owner: res.Owner.Login,
	}
	return &result, nil
}

func (s *repoService) CreateMultipleRepos(requests []repositories.CreateRepoRequest) (repositories.CreateMultipleReposResponse, errors.APIError) {
	input := make(chan repositories.CreateMultipleReposResult)
	output := make(chan repositories.CreateMultipleReposResponse)
	defer close(output)

	var wg sync.WaitGroup
	go s.handleRepoResults(&wg, input, output)

	for _, current := range requests {
		wg.Add(1)
		go s.createRepoConcurrent(current, input)
	}

	wg.Wait()
	close(input)

	result := <-output

	successCount := 0
	for _, current := range result.Results {
		if current.Response != nil {
			successCount++
		}
	}

	if successCount == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else {

		if successCount == len(requests) {
			result.StatusCode = http.StatusCreated
		} else {
			result.StatusCode = http.StatusPartialContent
		}
	}

	return result, nil
}

func (s *repoService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateMultipleReposResult, output chan repositories.CreateMultipleReposResponse) {
	var results repositories.CreateMultipleReposResponse

	for incomingEvent := range input {
		repoResult := repositories.CreateMultipleReposResult{
			Response: incomingEvent.Response,
			Error:    incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)

		wg.Done()
	}

	output <- results
}

func (s *repoService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateMultipleReposResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateMultipleReposResult{Error: err}
		return
	}

	result, err := s.CreateRepo(input)

	if err != nil {
		output <- repositories.CreateMultipleReposResult{Error: err}
		return
	}

	output <- repositories.CreateMultipleReposResult{Response: result}
}
