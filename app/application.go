package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nanoTitan/analytics-users-api/logger"
)

var (
	router = gin.Default()
)

/*
StartApplication starts the Go application
*/
func StartApplication() {
	mapUrls()

	logger.Info("starting the application...")
	router.Run(":8080")
}
