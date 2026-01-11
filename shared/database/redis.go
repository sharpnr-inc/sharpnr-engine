package database

import (
	"context"
	"log"
	"os"
	"sharpnr-engine/config"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis(cfg config.RedisDBConfig) *redis.Client {

	redisHost := os.Getenv("REDIS_HOST")         // 127.0.0.1:6379
	redisPassword := os.Getenv("REDIS_PASSWORD") // "" if no password

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       0,
	})

	// check the connection
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("✅ Redis connected successfully.")
	return RedisClient
}
