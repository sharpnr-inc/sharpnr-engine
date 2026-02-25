package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	env := os.Getenv("APP_ENV")

	file := ".env.development" // default (local dev)

	switch env {
	case "development":
		file = ".env.development"
	case "production":
		file = ".env.production"
	case "docker":
		file = ".env.docker"
	}

	if err := godotenv.Load(file); err != nil {
		log.Printf("⚠️ could not load %s", file)
	}
}
