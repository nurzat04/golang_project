package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func Run() error {
	_ = godotenv.Load()

	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	r := SetupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return r.Run(":" + port)
}

func saveURL(shortCode, url string) error {
	return rdb.Set(ctx, shortCode, url, 0).Err()
}

func generateShortCode(url string) string {
	return fmt.Sprintf("%x", len(url)+int(url[0]))[:6]
}
