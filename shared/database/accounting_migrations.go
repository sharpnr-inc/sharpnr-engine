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
	)
}
