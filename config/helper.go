package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

// --------------------
// String
// --------------------

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// --------------------
// Required string
// --------------------

func mustGet(key string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	log.Fatalf("❌ missing required environment variable: %s", key)
	return ""
}

// --------------------
// Int
// --------------------

func getEnvAsInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("❌ invalid int for %s: %v", key, err)
		}
		return i
	}
	return fallback
}

// --------------------
// Duration (THIS IS WHAT YOU ASKED)
// --------------------

func getEnvAsDuration(key string, fallback string) time.Duration {
	if v := os.Getenv(key); v != "" {
		d, err := time.ParseDuration(v)
		if err != nil {
			log.Fatalf("❌ invalid duration for %s: %v", key, err)
		}
		return d
	}

	d, err := time.ParseDuration(fallback)
	if err != nil {
		log.Fatalf("❌ invalid fallback duration for %s: %v", key, err)
	}
	return d
}

// --------------------
// Config validation helper
// --------------------

func must(value string, name string) {
	if value == "" {
		log.Fatalf("❌ missing required config value: %s", name)
	}
}
