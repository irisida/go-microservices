package controllers

import (
	"log"
	"net/http"
)

// GetUser will return all users
func GetUser(res http.ResponseWriter, req *http.Request) {
	// some functionality to be implemented
	userID := req.URL.Query().Get("id")
	log.Printf("Processing user Id: %v", userID)
}
