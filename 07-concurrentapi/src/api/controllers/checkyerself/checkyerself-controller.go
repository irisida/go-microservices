package checkyerself

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	verification = "before you wreck yerself!"
)

// WreckYerself lets a cloud provider know we're ready for traffic
func WreckYerself(c *gin.Context) {
	c.String(http.StatusOK, verification)
}
