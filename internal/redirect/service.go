package redirect

import (
	"context"
	"fmt"
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
	port := config.GetEnv("PORT", "8081")
	logger.InfoLogger.Println("Redirect service running on port:", port)
	return r.Run(":" + port)
}

func getOriginalURL(code string) (string, error) {
	val, err := rdb.Get(ctx, code).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("code not found")
	} else if err != nil {
		return "", err
	}
	return val, nil
}
