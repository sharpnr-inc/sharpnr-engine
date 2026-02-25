package database

import (
	"context"
	"fmt"
	"log"
	"sharpnr-engine/config"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis(cfg config.RedisDBConfig) *redis.Client {

	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("✅ Redis connected successfully.")
	return RedisClient
}
