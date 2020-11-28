package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/irisida/go-microservices/02-mvc/services"
)

// GetUser will return all users
func GetUser(res http.ResponseWriter, req *http.Request) {
	userID, err := strconv.ParseInt(req.URL.Query().Get("id"), 10, 64)

	if err != nil {
		// return bad request to the client
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("user id is not the correct format"))
		return
	}

	user, err := services.GetUser(userID)
	if err != nil {
		// handle err and return to the client
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(err.Error()))
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
