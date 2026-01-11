package config

import (
	"os"
	"time"
)

type RedisDBConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
	TTL      time.Duration
}

func loadRedis() RedisDBConfig {
	return RedisDBConfig{
		Host:     getEnv("REDIS_HOST", "127.0.0.1"),
		Port:     getEnv("REDIS_PORT", "6379"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       getEnvAsInt("REDIS_DB", 0),
		TTL:      getEnvAsDuration("REDIS_TTL", "10m"),
	}
}
