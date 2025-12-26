package accounting_repo

import (
	"context"
	"fmt"
	"sharpnr-engine/modules/accounting/accounting_models"
)

// TODO :-> Need Documentation
func (d *accountingRepository) CreateAccountMapping(ctx context.Context, accountMapping *accounting_models.AccountMapping) error {
	result := d.db.WithContext(ctx).Create(accountMapping)

	if result.Error != nil {
		// It's helpful to wrap the error to know exactly where it failed
		return fmt.Errorf("failed to create account mapping: %w", result.Error)
	}
	return nil
}

// TODO :-> Need Documentation
func (d *accountingRepository) GetAccountingMapping(ctx context.Context, str string) (*accounting_models.AccountMapping, error) {
	panic("error")
}
