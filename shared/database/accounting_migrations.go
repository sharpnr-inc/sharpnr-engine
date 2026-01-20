package database

import (
	"sharpnr-engine/modules/accounting/accounting_models"

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
	)
}
