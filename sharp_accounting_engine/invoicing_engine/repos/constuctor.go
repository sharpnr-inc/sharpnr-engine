package invoice_engine_repos

import "gorm.io/gorm"

type invoiceEngineRepo struct {
	db *gorm.DB
}

func NewInvoiceEngineRepo(db *gorm.DB) *invoiceEngineRepo {
	return &invoiceEngineRepo{
		db: db,
	}
}
