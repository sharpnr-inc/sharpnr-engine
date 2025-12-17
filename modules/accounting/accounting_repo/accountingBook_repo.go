package accounting_repo

import (
	"context"
	"sharpnr-engine/modules/accounting/accounting_models"

	"github.com/google/uuid"
)

// CreateAccountingBook implements AccountingRepositoryInterface.
func (r *accountingRepository) CreateAccountingBook(ctx context.Context, book *accounting_models.AccountingBook) error {
	panic("unimplemented")
}

// GetAccountingBookByID implements AccountingRepositoryInterface.
func (r *accountingRepository) GetAccountingBookByID(ctx context.Context, bookID uuid.UUID) (*accounting_models.AccountingBook, error) {
	panic("unimplemented")
}

// GetAccountingBooksByOrgID implements AccountingRepositoryInterface.
func (r *accountingRepository) GetAccountingBooksByOrgID(ctx context.Context, orgID uuid.UUID) ([]accounting_models.AccountingBook, error) {
	panic("unimplemented")
}
