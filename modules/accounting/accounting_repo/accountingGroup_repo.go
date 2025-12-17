package accounting_repo

import (
	"context"
	"sharpnr-engine/modules/accounting/accounting_models"

	"github.com/google/uuid"
)

// CreateAccountGroup implements AccountingRepositoryInterface.
func (r *accountingRepository) CreateAccountGroup(ctx context.Context, group *accounting_models.AccountGroup) error {
	panic("unimplemented")
}

// GetAccountGroupByID implements AccountingRepositoryInterface.
func (r *accountingRepository) GetAccountGroupByID(ctx context.Context, groupID uuid.UUID) (*accounting_models.AccountGroup, error) {
	panic("unimplemented")
}

// GetAccountGroupsByBookID implements AccountingRepositoryInterface.
func (r *accountingRepository) GetAccountGroupsByBookID(ctx context.Context, bookID uuid.UUID) ([]accounting_models.AccountGroup, error) {
	panic("unimplemented")
}
