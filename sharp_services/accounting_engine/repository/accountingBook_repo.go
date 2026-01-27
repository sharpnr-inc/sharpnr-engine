package accounting_repo

import (
	"context"
	"errors"
	accounting_models "sharpnr-engine/sharp_services/accounting_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateAccountingBook implements AccountingRepositoryInterface.
func (r *accountingRepository) CreateAccountingBook(ctx context.Context, book *accounting_models.AccountingBook) error {
	return r.db.WithContext(ctx).Create(book).Error
}

// GetAccountingBookByID implements AccountingRepositoryInterface.
func (r *accountingRepository) GetAccountingBookByID(ctx context.Context, bookID uuid.UUID) (*accounting_models.AccountingBook, error) {
	var book accounting_models.AccountingBook

	err := r.db.WithContext(ctx).
		Where("id = ?", bookID).
		First(&book).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &book, nil
}
