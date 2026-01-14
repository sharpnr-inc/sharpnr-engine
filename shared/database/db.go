package database

import (
	"fmt"
	"log"

	"sharpnr-engine/config"
	"sharpnr-engine/core/models"
	"sharpnr-engine/modules/accounting/accounting_models"
	bar_models "sharpnr-engine/modules/hospitality/bar/models"
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"
	kitchen_models "sharpnr-engine/modules/hospitality/kitchen/models"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"
	sharphr_models "sharpnr-engine/modules/hr/hrModels"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg config.DatabaseConfig) *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port,
	)

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
		&models.OrganizationSubscription{},
		&models.OrganizationOffice{},
		&models.OrganizationSetting{},
		&models.OrganizationVerification{},
		&models.OrganizationSetting{},
		// More to come

		// Core --> Hospitality

		// -------------------------------
		// Employee Models
		// -------------------------------

		&sharphr_models.Employee{},
		&sharphr_models.EmployeeOnboarding{},
		&sharphr_models.EmployeeVerification{},
		&sharphr_models.EmployeeAuth{},
		&sharphr_models.EmployeePrivilegePolicy{},
		&sharphr_models.EmployeeSession{},
		&sharphr_models.EmployeeProfile{},
		&sharphr_models.EmployeeBankAccount{},
		&sharphr_models.EmployeeDocument{},
		&sharphr_models.EmployeeReporting{},
		&sharphr_models.EmployeeEmployment{},
		&sharphr_models.EmployeeCompensation{},
		&sharphr_models.EmployeeSalaryComponent{},
		&sharphr_models.EmployeeLeaveType{},
		&sharphr_models.EmployeeLeaveBalance{},
		&sharphr_models.EmployeeLeaveRequest{},

		&sharphr_models.EmployeePolicy{},
		&sharphr_models.EmployeePolicyRule{},
		&sharphr_models.EmployeePrivilegePolicy{},

		// -------------------------------
		// 🏨 HOSPITALITY MODELS
		// -------------------------------
		// ────────────── CORE HOTEL STRUCTURE ──────────────
		&hotel_models.Hotel{},
		&hotel_models.HotelFloor{},
		&hotel_models.Room{},
		&hotel_models.HotelRoomPrice{},

		// ────────────── GUESTS ──────────────
		&hotel_models.HotelGuest{},

		// ────────────── BOOKINGS (PARENT) ──────────────
		&hotel_models.HotelRoomBooking{},

		// ────────────── BOOKING CHILD TABLES ──────────────
		&hotel_models.BookingGuest{},
		&hotel_models.HotelRoomBookingPayment{},
		&hotel_models.HotelBookingCharge{},

		// ────────────── SERVICES ──────────────
		&hotel_models.HotelService{},
		&hotel_models.HotelServiceItem{},

		// ────────────── TAX ──────────────
		&hotel_models.HotelTaxRule{},
		&hotel_models.HotelTaxComponent{},

		// ────────────── BILLING ──────────────
		&hotel_models.HotelBill{},
		&hotel_models.HotelBillTax{},

		// Hospitality --> Bar
		&bar_models.Bar{},
		&bar_models.BarFloor{},
		&bar_models.BarTable{},
		// Bar-orders
		&bar_models.BarTableSession{},
		&bar_models.BarTablesSessionOrder{},

		// Hospitality --> Restaurant
		&restaurant_models.Restaurant{},
		&restaurant_models.RestaurantFloor{},
		&restaurant_models.RestaurantTable{},
		&restaurant_models.RestaurantTableBookingSession{},

		//Hospitality --> Kitchen
		&kitchen_models.Kitchen{},
		&kitchen_models.KitchenTicket{},
		&kitchen_models.KitchenTicketItem{},
		&kitchen_models.KitchenTicketItemAddon{},

		//---------------------------------------------------

		// ✅ MENU SYSTEM
		&hospitalityMenu_models.MenuGroup{},
		&hospitalityMenu_models.MenuCategory{},
		&hospitalityMenu_models.MenuSubCategory{},
		&hospitalityMenu_models.MenuItem{},
		&hospitalityMenu_models.MenuVariant{},
		&hospitalityMenu_models.MenuVariantGroup{},
		&hospitalityMenu_models.Addon{},

		// -------------------------------
		// 💰 Hospitality Customer & Guest MODELS
		// -------------------------------

		// -------------------------------
		// 💰 ACCOUNTING MODELS
		// -------------------------------
		&accounting_models.AccountingBook{},
		&accounting_models.AccountingBookSetting{},

		&accounting_models.FiscalYear{},
		&accounting_models.AccountGroup{},
		&accounting_models.Account{},
		&accounting_models.AccountMapping{},

		&accounting_models.JournalEntry{},
		&accounting_models.JournalLine{},

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
		&accounting_models.AuditLog{},

		// HR

	)
}
