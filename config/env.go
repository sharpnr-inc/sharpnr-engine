package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	if os.Getenv("APP_ENV") == "development" {
		if err := godotenv.Load(".env.development"); err != nil {
			log.Println("⚠️ .env.development not found")
		}
	}
}
