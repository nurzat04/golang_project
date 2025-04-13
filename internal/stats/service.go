package stats

import (
	"context"
	"shorty/pkg/config"
	"shorty/pkg/logger"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func Run() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.GetEnv("REDIS_ADDR", "localhost:6379"),
		Password: config.GetEnv("REDIS_PASSWORD", ""),
		DB:       0,
	})

	r := SetupRouter()
	port := config.GetEnv("PORT", "8082")
	logger.InfoLogger.Println("Stats service running on port:", port)
	return r.Run(":" + port)
}
