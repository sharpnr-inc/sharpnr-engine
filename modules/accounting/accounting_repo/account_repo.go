package accounting_repo

import (
	"context"
	"sharpnr-engine/modules/accounting/accounting_models"

	"github.com/google/uuid"
)

// CreateAccount implements AccountingRepositoryInterface.
func (r *accountingRepository) CreateAccount(ctx context.Context, account *accounting_models.Account) error {
	panic("unimplemented")
}

// GetAccountByID implements AccountingRepositoryInterface.
func (r *accountingRepository) GetAccountByID(ctx context.Context, accountID uuid.UUID) (*accounting_models.Account, error) {
	panic("unimplemented")
}

// GetAccountsByGroupID implements AccountingRepositoryInterface.
func (r *accountingRepository) GetAccountsByGroupID(ctx context.Context, groupID uuid.UUID) ([]accounting_models.Account, error) {
	panic("unimplemented")
}
