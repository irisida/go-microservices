package happywave

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	happywave = "happywave"
)

// HappyWave lets a cloud provoder know we're ready for traffic
func HappyWave(c *gin.Context) {
	c.String(http.StatusOK, happywave)
}
