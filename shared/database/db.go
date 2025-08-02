package database

import (
	"fmt"
	"log"
	"os"

	"sharpnr-engine/core/models"
	accountingmodels "sharpnr-engine/modules/base/accounting/accounting_models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	DB = db
	log.Println("✅ PostgreSQL connected successfully.")
	return DB
}

// Run migration to create tables
// This Function will Create Database
func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Organization{}, // Parent of everything
		&models.Department{},   // Belongs to Organization
		&models.Employee{},     // Belongs to Department + Organization
		&models.Product{},      // Belongs to Organization
		&models.Service{},      // Belongs to Organization
		&models.Operations{},   // Belongs to Department + Organization
		&models.Terminal{},     // Belongs to Organization + has many-to-many with Employee & Operations
		&accountingmodels.AccountingBills{},
	)
}
