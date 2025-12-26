package accounting_repo

import (
	"context"
	"fmt"
	"sharpnr-engine/modules/accounting/accounting_models"

	"github.com/google/uuid"
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

// Return array of the Account mapping if they have
func (d *accountingRepository) GetAccountMappingsByOrg(ctx context.Context, orgID uuid.UUID) ([]accounting_models.AccountMapping, error) {
	var mappings []accounting_models.AccountMapping

	err := d.db.WithContext(ctx).
		Preload("Account").Where("org_id = ?", orgID).Order("source_type ASC, source_code ASC").
		Find(&mappings).
		Error

	if err != nil {
		return nil, err
	}

	return mappings, nil
}
