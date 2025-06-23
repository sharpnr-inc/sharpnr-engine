package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/sharpnr-inc/sharpnr-engine/config"
)

var DB *gorm.DB

func ConnectDB() {
	log.Println("📦 Starting database connection...")
	config.LoadEnv()
	log.Println("✅ Environment variables loaded.")

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	log.Printf("🔌 Connecting to DB at %s:%s (DB: %s, User: %s)\n", dbHost, dbPort, dbName, dbUser)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	DB = db
	log.Println("✅ PostgreSQL connected successfully.")
}
