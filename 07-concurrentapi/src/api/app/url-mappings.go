package app

import (
	"github.com/irisida/go-microservices/06-githubapi/src/api/controllers/checkyerself"
	"github.com/irisida/go-microservices/06-githubapi/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/checkyerself", checkyerself.WreckYerself)
	router.POST("/repositories", repositories.CreateRepo)
}
