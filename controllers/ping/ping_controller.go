package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping - returns a 'pong' response on success with status 200
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
