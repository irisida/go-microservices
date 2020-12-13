package app

import (
	"github.com/irisida/go-microservices/07-concurrentapi/src/api/controllers/checkyerself"
	"github.com/irisida/go-microservices/07-concurrentapi/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/checkyerself", checkyerself.WreckYerself)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateMultipleRepos)
}
