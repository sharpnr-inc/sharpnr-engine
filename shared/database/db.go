package database

import (
	"fmt"
	"log"
	"os"

	"sharpnr-engine/core/models"
	accountingmodels "sharpnr-engine/modules/accounting/accounting_models"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

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

		&models.Organization{},
		&models.OrganizationPolicy{},
		&models.OrgUIRendererTree{},
		&models.OrganizationBankAccount{},
		&models.Employee{},
		&hospitalitymodels.Hotel{},
		&hospitalitymodels.HotelFloor{},
		&hospitalitymodels.Room{},
		&hospitalitymodels.Bar{},
		&hospitalitymodels.BarFloor{},
		&hospitalitymodels.BarTable{},
		&hospitalitymodels.Restaurant{},
		&hospitalitymodels.RestaurantFloor{},
		&hospitalitymodels.RestaurantTable{},
		&hospitalitymodels.Kitchen{},
		&hospitalitymodels.KitchenTicket{},
		&hospitalitymodels.KitchenTicketItem{},
		&hospitalitymodels.HospitalityOrder{},
		&hospitalitymodels.HospitalityOrderItem{},
		&hospitalitymodels.HospitalityOrderAddon{},
		&accountingmodels.AccountingBills{},
	)
}
