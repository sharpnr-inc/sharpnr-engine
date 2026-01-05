package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadJwtSecret() string {
	// Load .env only for local/dev
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system env")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("❌ JWT_SECRET is not set")
	}

	return secret
}
