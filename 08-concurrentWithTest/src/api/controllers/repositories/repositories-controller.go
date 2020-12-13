package repositories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irisida/go-microservices/08-concurrentWithTest/src/api/domain/repositories"
	"github.com/irisida/go-microservices/08-concurrentWithTest/src/api/services"
	"github.com/irisida/go-microservices/08-concurrentWithTest/src/api/utils/errors"
)

// CreateRepo main controller
func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		// cannot create valid json as the body of the request.
		apiError := errors.NewBadRequestAPIError("invalid json body - controller create repo")
		c.JSON(apiError.Status(), apiError)
		return
	}

	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	// no error scenario, send successful request
	c.JSON(http.StatusCreated, result)
}

// CreateMultipleRepos main controller
func CreateMultipleRepos(c *gin.Context) {
	var request []repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		// cannot create valid json as the body of the request.
		apiError := errors.NewBadRequestAPIError("invalid json body multiple repos")
		c.JSON(apiError.Status(), apiError)
		return
	}

	result, err := services.RepositoryService.CreateMultipleRepos(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	// no error scenario, send successful request
	c.JSON(result.StatusCode, result)
}
