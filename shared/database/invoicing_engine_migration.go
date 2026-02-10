package database

import (
	invoice_engine_models "sharpnr-engine/service_engines/invoicing_engine/models"

	"gorm.io/gorm"
)

func migrateInvoicingEngine(db *gorm.DB) error {
	return db.AutoMigrate(
		&invoice_engine_models.Invoice{},
		&invoice_engine_models.InvoiceItem{},
		&invoice_engine_models.InvoicePaymentLink{},
		invoice_engine_models.InvoiceTaxBreakup{},
	)
}
