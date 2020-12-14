package repositories

import (
	"strings"

	"github.com/irisida/go-microservices/08-concurrentWithTest/src/api/utils/errors"
)

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *CreateRepoRequest) Validate() errors.APIError {
	r.Name = strings.TrimSpace(r.Name)

	if r.Name == "" {
		return errors.NewBadRequestAPIError("invalid repository name")
	}
	return nil
}

type CreateRepoResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type CreateManyRepositoriesResponse struct {
	StatusCode int                            `json:"status"`
	Results    []CreateManyRepositoriesResult `json:"results"`
}

type CreateManyRepositoriesResult struct {
	Response *CreateRepoResponse `json:"repo"`
	Error    errors.APIError     `json:"error"`
}
