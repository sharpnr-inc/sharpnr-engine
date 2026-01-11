package database

import (
	"context"
	"log"
	"sharpnr-engine/config"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis(cfg config.RedisDBConfig) *redis.Client {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Host,
		Password: cfg.Password,
		DB:       0,
	})

	// check the connection
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("✅ Redis connected successfully.")
	return RedisClient
}
