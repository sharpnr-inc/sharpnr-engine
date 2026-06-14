package accounting_repo

import (
	"gorm.io/gorm"
)

type accountingRepository struct {
	db *gorm.DB
}

func NewAccountingRepository(db *gorm.DB) *accountingRepository {
	return &accountingRepository{
		db: db,
	}
}
