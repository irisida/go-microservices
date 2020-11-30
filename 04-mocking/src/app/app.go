package app

import (
	"net/http"

	"github.com/irisida/go-microservices/04-mocking/src/controllers"
)

// StartApp is the ignition key called from the main.go
// in the project to start the application
func StartApp() {
	// getUser
	http.HandleFunc("/users", controllers.GetUser)

	// ListenAndServe
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
