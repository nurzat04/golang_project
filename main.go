package main

import (
	"stats-service/config"
	"stats-service/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	rdb := config.InitRedis()

	r.POST("/stats", handler.GetStats(rdb))

	r.Run(":8082") // Stats Service 监听 8082 端口
}
