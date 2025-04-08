package handler

import (
	"net/http"
	"stats-service/config"
	"stats-service/model"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func GetStats(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.StatsRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		key := "stats:" + req.Code
		count, err := rdb.Get(config.Ctx, key).Int64()
		if err == redis.Nil {
			c.JSON(http.StatusOK, model.StatsResponse{Code: req.Code, Count: 0})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis error"})
			return
		}

		c.JSON(http.StatusOK, model.StatsResponse{Code: req.Code, Count: count})
	}
}
