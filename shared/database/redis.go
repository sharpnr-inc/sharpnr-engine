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

		// Without these, a Redis that accepts the TCP connection but never
		// answers leaves every command blocked indefinitely, and the request
		// holding it blocked with them.
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,

		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	// Bounded so an unreachable Redis fails startup instead of hanging it —
	// context.Background() here would wait forever.
	ctx, cancel := context.WithTimeout(context.Background(), cfg.DialTimeout)
	defer cancel()

	if err := RedisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("❌ Failed to connect to Redis at %s: %v", addr, err)
	}

	log.Println("✅ Redis connected successfully.")
	return RedisClient
}
