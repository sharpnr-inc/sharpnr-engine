package accounting_repo

import (
	"context"
	accounting_models "sharpnr-engine/sharp_services/accounting_engine/models"

	"github.com/google/uuid"
)

// CreateAccountingBook implements AccountingRepositoryInterface.
func (r *accountingRepository) CreateAccountingBook(ctx context.Context, book *accounting_models.AccountingBook) error {
	return r.db.WithContext(ctx).Create(book).Error
}

func (r *accountingRepository) GetAccountingBookByID(ctx context.Context, accounting_bookID uuid.UUID) ([]*accounting_models.AccountingBook, error) {
	var books []*accounting_models.AccountingBook

	// .WithContext(ctx) ensures the query respects timeouts/cancellations
	// .Where() filters by the ID
	// .Find() populates the slice
	err := r.db.WithContext(ctx).
		Where("id = ?", accounting_bookID).
		Find(&books).Error

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *accountingRepository) GetAccountingBookByLegalEntityID(
	ctx context.Context,
	legal_entityId uuid.UUID,
) (*accounting_models.AccountingBook, error) {

	var book accounting_models.AccountingBook

	// Use .WithContext for trace/timeout propagation
	// Use .Where to filter by the Legal Entity ID
	// Use .First to retrieve a single record
	err := r.db.WithContext(ctx).
		Where("organization_legal_entity_id = ?", legal_entityId).
		First(&book).Error

	if err != nil {
		// GORM returns gorm.ErrRecordNotFound if no match is found.
		// You can return the error directly or wrap it.
		return nil, err
	}

	return &book, nil
}
