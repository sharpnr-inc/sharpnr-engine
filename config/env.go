package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	env := os.Getenv("APP_ENV")

	file := ".env"
	switch env {
	case "development":
		file = ".env.development"
	case "production":
		file = ".env.production"
	}

	if err := godotenv.Load(file); err != nil {
		log.Printf("⚠️ could not load %s\n", file)
	}
}
