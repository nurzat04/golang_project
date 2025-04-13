package stats

import (
	"net/http"
	"shorty/pkg/logger"

	"github.com/gin-gonic/gin"
)

func statsHandler(c *gin.Context) {
	code := c.Param("code")
	key := "stats:" + code
	count, err := rdb.Get(ctx, key).Int()
	if err != nil {
		logger.ErrorLogger.Println("Failed to get stats:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "No stats found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": code, "clicks": count})
}
