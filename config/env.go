package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	file := ".env.development" // default (local dev)

	switch os.Getenv("APP_ENV") {
	case EnvDevelopment:
		file = ".env.development"
	case EnvProduction:
		file = ".env.production"
	case EnvDocker:
		file = ".env.docker"
	}

	// Not fatal: containerised deployments inject the variables directly and
	// ship no dotenv file at all. Anything genuinely required is enforced by
	// mustGet/validate below, which fail loudly on their own.
	if err := godotenv.Load(file); err != nil {
		log.Printf("⚠️  could not load %s — falling back to process environment", file)
	}
}
