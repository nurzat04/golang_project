package shorten

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/api/v1/shorten", shortenHandler)
	return r
}
