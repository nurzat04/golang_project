package stats

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/api/v1/stats/:code", statsHandler)
	return r
}
