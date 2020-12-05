package app

import (
	"github.com/irisida/go-microservices/05-gin/src/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
