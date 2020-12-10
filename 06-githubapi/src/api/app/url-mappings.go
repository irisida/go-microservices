package app

import (
	"github.com/irisida/go-microservices/06-githubapi/src/api/controllers/happywave"
	"github.com/irisida/go-microservices/06-githubapi/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/arewegood", happywave.HappyWave)
	router.POST("/repositories", repositories.CreateRepo)
}
