package app

import (
	"github.com/martinyonathann/github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	router.Run(":8081")
}
