package invoice_engine_repos

import "gorm.io/gorm"

type invoiceEngineRepo struct {
	db *gorm.DB
}
