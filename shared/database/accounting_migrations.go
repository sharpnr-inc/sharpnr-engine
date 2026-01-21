package database

import (
	accounting_models "sharpnr-engine/sharp_services/accounting_engine/models"

	"gorm.io/gorm"
)

func migrateAccounting(db *gorm.DB) error {
	return db.AutoMigrate(
		&accounting_models.AccountingBook{},
		&accounting_models.AccountingBookSetting{},

		&accounting_models.FiscalYear{},
		&accounting_models.AccountGroup{},
		&accounting_models.Account{},

		&accounting_models.JournalEntry{},
		&accounting_models.JournalLine{},
		&accounting_models.LedgerBalance{},

		// Payables & Expenses
		&accounting_models.Supplier{},
		&accounting_models.PurchaseBill{},
		&accounting_models.PurchaseBillItem{},
		&accounting_models.SupplierPayment{},
		&accounting_models.Expense{},
	)
}
