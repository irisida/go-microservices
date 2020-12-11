package services

import (
	"strings"

	"github.com/irisida/go-microservices/06-githubapi/src/api/config"
	"github.com/irisida/go-microservices/06-githubapi/src/api/domain/github"
	"github.com/irisida/go-microservices/06-githubapi/src/api/domain/repositories"
	"github.com/irisida/go-microservices/06-githubapi/src/api/providers/githubprovider"
	"github.com/irisida/go-microservices/06-githubapi/src/api/utils/errors"
)

type repoService struct{}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError)
}

var (
	// RepositoryService is the exported service
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {
	input.Name = strings.TrimSpace(input.Name)

	if input.Name == "" {
		return nil, errors.NewBadRequestAPIError("invalid repository name")
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
