package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanoTitan/analytics-users-api/domain/users"
	"github.com/nanoTitan/analytics-users-api/services"
	"github.com/nanoTitan/analytics-users-api/utils/errors"
)

// CreateUser - create a new user given a user_id
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// SearchUser - find a user given a user_id
func SearchUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "pong",
	})
}

// GetUser - return a user
func GetUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "pong",
	})
}
