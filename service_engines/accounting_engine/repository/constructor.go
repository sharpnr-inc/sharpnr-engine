package accounting_repo

import (
	"context"

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

func (r *accountingRepository) WithTx(
	ctx context.Context,
	fn func(txRepo AccountingRepositoryInterface) error,
) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txRepo := &accountingRepository{db: tx}
		return fn(txRepo)
	})
}
