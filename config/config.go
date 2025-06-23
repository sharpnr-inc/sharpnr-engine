package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from a specific .env file based on the `APP_ENV` value.
// Supported environments:
//   - "alpha": loads variables from `.env.alpha`
//   - "beta": loads variables from `.env.beta`
//   - default: loads variables from `.env`
//
// The function uses the `github.com/joho/godotenv` package to load environment variables from files.
// If the file cannot be loaded, the function logs a fatal error and stops the application.
//
// Example usage:
//
//	config.LoadEnv()
//
// Environment variable to set before running:
//
//	export APP_ENV=alpha
//
// Typical use case:
//
//	Separate environment configurations for different deployment stages like development, staging (beta), and production.
func LoadEnv() {
	env := os.Getenv("APP_ENV")
	var envFile string

	switch env {
	case "alpha":
		envFile = ".env.alpha"
	case "beta":
		envFile = ".env.beta"
	case "production":
		envFile = ".env.production"
	default:
		envFile = ".env"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("❌ Failed to load %s file: %v", envFile, err)
	}

	log.Printf("✅ Loaded environment: %s", envFile)
}
