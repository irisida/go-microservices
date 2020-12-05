package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/irisida/go-microservices/05-gin/src/services"
	"github.com/irisida/go-microservices/05-gin/src/utils"
	"net/http"
	"strconv"
)

// GetUser will return all users
func GetUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		// return bad request to the client
		APIError := &utils.ApplicationError{
			Message:    "user-id is not in the correct format",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}

		utils.RespondError(c, APIError)
		return
	}

	user, APIError := services.UsersService.GetUser(userID)
	if APIError != nil {
		utils.RespondError(c, APIError)
		return
	}

	// return user to client
	utils.Respond(c, http.StatusOK, user)
}
