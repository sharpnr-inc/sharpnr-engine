package database

import (
	"fmt"
	"log"
	"os"

	"sharpnr-engine/core/models"
	"sharpnr-engine/modules/accounting/accounting_models"
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
		&accounting_models.AccountingBook{},
		&accounting_models.Account{},
		&accounting_models.AccountGroup{},
		&accounting_models.FiscalYear{},
		&accounting_models.Voucher{},
		&accounting_models.VoucherEntry{},
		&accounting_models.Transaction{},
		&accounting_models.Invoice{},
		&accounting_models.InvoiceItem{},
		&accounting_models.Bill{},
		&accounting_models.BillItem{},
		&accounting_models.TaxRate{},
		&accounting_models.GSTProfile{},
		&accounting_models.Customer{},
		&accounting_models.Supplier{},
		&accounting_models.LedgerBalance{},
		&accounting_models.ProfitLossSummary{},
		&accounting_models.BalanceSheetSummary{},
		&accounting_models.AccountingSetting{},
		&accounting_models.AuditLog{},
	)
}
