package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	gin.SetMode(gin.DebugMode)
	mapUrls()
	router.Run(":8081")
}
