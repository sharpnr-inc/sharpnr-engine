package database

import (
	"fmt"
	"log"
	"os"

	"sharpnr-engine/core/models"
	"sharpnr-engine/modules/accounting/accounting_models"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	sharphr_models "sharpnr-engine/modules/hr/hrModels"

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

	// dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	DB = db
	log.Println("✅ PostgreSQL connected successfully.")
	return DB
}

// -------------------------------
// 🧱 Run Migration to Create Tables
// -------------------------------
func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		// -------------------------------
		// 🌐 CORE MODELS
		// -------------------------------
		&models.Organization{},
		&models.OrganizationPolicy{},
		&models.OrgUIRendererTree{},
		&models.OrganizationBrandingDetails{},
		&models.OrganizationBankAccount{},
		&sharphr_models.Employee{},

		// -------------------------------
		// 🏨 HOSPITALITY MODELS
		// -------------------------------
		&hospitalitymodels.Hotel{},
		&hospitalitymodels.HotelFloor{},
		&hospitalitymodels.Room{},
		&hospitalitymodels.HotelGuest{},
		&hospitalitymodels.HotelRoomBooking{},
		&hospitalitymodels.HotelBookingOrder{},
		&hospitalitymodels.HotelBookingOrderItem{},
		&hospitalitymodels.HotelService{},
		&hospitalitymodels.HotelServiceItem{},
		&hospitalitymodels.BookingCharge{},
		&hospitalitymodels.BookingPayment{},

		// Hospitality --> Bar
		&hospitalitymodels.Bar{},
		&hospitalitymodels.BarFloor{},
		&hospitalitymodels.BarTable{},

		// Hospitality --> Restaurant
		&hospitalitymodels.Restaurant{},
		&hospitalitymodels.RestaurantFloor{},
		&hospitalitymodels.RestaurantTable{},
		&hospitalitymodels.RestaurantTableBookingSession{},
		&hospitalitymodels.RestaurantTableOrder{},

		//Hospitality --> Kitchen
		&hospitalitymodels.Kitchen{},
		&hospitalitymodels.KitchenTicket{},
		&hospitalitymodels.KitchenTicketItem{},

		//---------------------------------------------------

		// ✅ MENU SYSTEM
		&hospitalitymodels.MenuGroup{},
		&hospitalitymodels.MenuCategory{},
		&hospitalitymodels.MenuSubCategory{},
		&hospitalitymodels.MenuItem{},
		&hospitalitymodels.MenuVariant{},
		&hospitalitymodels.MenuVariantGroup{},
		&hospitalitymodels.Addon{},

		// -------------------------------
		// 💰 Hospitality Customer & Guest MODELS
		// -------------------------------

		// -------------------------------
		// 💰 ACCOUNTING MODELS
		// -------------------------------
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

		// HR

	)
}
