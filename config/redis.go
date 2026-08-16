package config

import (
	"os"
	"time"
)

// RedisDBConfig holds connection settings for the shared Redis client.
//
// There is deliberately no global "TTL" here: every key written to Redis has
// its own lifetime driven by what it stores (see TTLConfig), and a single
// blanket TTL would either be ignored — as REDIS_TTL was — or silently wrong
// for most keys.
type RedisDBConfig struct {
	Host     string
	Port     string
	Password string
	DB       int

	// DialTimeout bounds how long a new connection may take to establish.
	// Also used to bound the startup PING so an unreachable Redis fails fast
	// instead of hanging the boot sequence forever.
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PoolSize     int
	MinIdleConns int
}

func loadRedis() RedisDBConfig {
	return RedisDBConfig{
		Host:     getEnv("REDIS_HOST", "127.0.0.1"),
		Port:     getEnv("REDIS_PORT", "6379"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       getEnvAsInt("REDIS_DB", 0),

		DialTimeout:  getEnvAsDuration("REDIS_DIAL_TIMEOUT", "5s"),
		ReadTimeout:  getEnvAsDuration("REDIS_READ_TIMEOUT", "3s"),
		WriteTimeout: getEnvAsDuration("REDIS_WRITE_TIMEOUT", "3s"),

		PoolSize:     getEnvAsInt("REDIS_POOL_SIZE", 10),
		MinIdleConns: getEnvAsInt("REDIS_MIN_IDLE_CONNS", 2),
	}
}
