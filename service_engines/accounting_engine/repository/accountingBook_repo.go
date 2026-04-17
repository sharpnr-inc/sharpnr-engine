package accounting_repo

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	accounting_models "sharpnr-engine/service_engines/accounting_engine/models"

	"github.com/google/uuid"
)

func (r *accountingRepository) CreateAccountingBook(ctx context.Context, book *accounting_models.AccountingBook) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	return db.WithContext(ctx).Create(book).Error
}

/* ---------------------------------- <***> --------------------------------- */

func (r *accountingRepository) GetAccountingBookByID(ctx context.Context, accounting_bookID uuid.UUID) (*accounting_models.AccountingBook, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var books *accounting_models.AccountingBook

	// .WithContext(ctx) ensures the query respects timeouts/cancellations
	// .Where() filters by the ID
	// .Find() populates the slice
	err := db.WithContext(ctx).
		Where("id = ?", accounting_bookID).
		Find(&books).Error

	if err != nil {
		return nil, err
	}

	return books, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *accountingRepository) GetAccountingBookByLegalEntityID(
	ctx context.Context,
	legal_entityId uuid.UUID,
) (*accounting_models.AccountingBook, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var book accounting_models.AccountingBook

	// Use .WithContext for trace/timeout propagation
	// Use .Where to filter by the Legal Entity ID
	// Use .First to retrieve a single record
	err := db.WithContext(ctx).
		Where("legal_entity_id = ?", legal_entityId).
		First(&book).Error

	if err != nil {
		// GORM returns gorm.ErrRecordNotFound if no match is found.
		// You can return the error directly or wrap it.
		return nil, err
	}

	return &book, nil
}

/* ---------------------------------- <***> --------------------------------- */
