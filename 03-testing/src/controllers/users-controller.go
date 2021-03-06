package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/irisida/go-microservices/03-testing/src/services"
	"github.com/irisida/go-microservices/03-testing/src/utils"
)

// GetUser will return all users
func GetUser(res http.ResponseWriter, req *http.Request) {
	userID, err := strconv.ParseInt(req.URL.Query().Get("id"), 10, 64)

	if err != nil {
		// return bad request to the client
		APIError := &utils.ApplicationError{
			Message:    "user-id is not in the correct format",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}

		jsonValue, _ := json.Marshal(APIError)
		res.WriteHeader(APIError.StatusCode)
		res.Write(jsonValue)
		return
	}

	user, APIError := services.GetUser(userID)
	if APIError != nil {
		jsonValue, _ := json.Marshal(APIError)
		res.WriteHeader(APIError.StatusCode)
		res.Write([]byte(jsonValue))
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
