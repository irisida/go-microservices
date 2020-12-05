package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

// in the project to start the application
func StartApp() {
	mapUrls()

	// ListenAndServe
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
