package repositories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irisida/go-microservices/06-githubapi/src/api/domain/repositories"
	"github.com/irisida/go-microservices/06-githubapi/src/api/services"
	"github.com/irisida/go-microservices/06-githubapi/src/api/utils/errors"
)

// CreateRepo main controller
func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		// cannot create valid json as the body of the request.
		apiError := errors.NewBadRequestApiError("invalid json body")
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
