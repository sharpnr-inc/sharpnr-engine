package accounting_repo

import (
	"context"
	"errors"
	accounting_models "sharpnr-engine/modules/accounting/models"

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

// GetAccountingBooksByOrgID implements AccountingRepositoryInterface.
func (r *accountingRepository) GetAccountingBooksByOrgID(ctx context.Context, orgID uuid.UUID) ([]accounting_models.AccountingBook, error) {
	var books []accounting_models.AccountingBook

	err := r.db.WithContext(ctx).
		Where("org_id = ?", orgID).
		Order("created_at ASC").
		Find(&books).
		Error

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *accountingRepository) GetAccountingBookByOrgAndSubModule(
	ctx context.Context,
	orgID uuid.UUID,
	subModule string,
) (*accounting_models.AccountingBook, error) {

	var book accounting_models.AccountingBook

	err := r.db.WithContext(ctx).
		Where(
			"org_id = ? AND sub_module = ?",
			orgID,
			subModule,
		).
		First(&book).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &book, nil
}

func (r *accountingRepository) GetAccountingBookByOrgAndModule(
	ctx context.Context,
	orgID uuid.UUID,
	module string,
) (*accounting_models.AccountingBook, error) {

	var book accounting_models.AccountingBook

	err := r.db.WithContext(ctx).
		Where(
			"org_id = ? AND module = ?",
			orgID,
			module,
		).
		First(&book).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &book, nil
}
