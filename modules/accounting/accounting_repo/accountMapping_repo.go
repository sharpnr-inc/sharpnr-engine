package accounting_repo

import (
	"context"
	"errors"
	"fmt"
	accounting_models "sharpnr-engine/modules/accounting/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (r *accountingRepository) ResolveAccountMapping(
	ctx context.Context,
	orgID uuid.UUID,
	module string,
	sourceType string,
	sourceCode string,
) (*accounting_models.Account, error) {

	var mapping accounting_models.AccountMapping

	err := r.db.
		WithContext(ctx).
		Preload("Account").
		Where(`
			org_id = ? AND
			sub_module = ? AND
			source_type = ? AND
			source_code = ?
		`,
			orgID,
			module,
			sourceType,
			sourceCode,
		).
		First(&mapping).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf(
				"no account mapping found for module=%s source_type=%s source_code=%s",
				module, sourceType, sourceCode,
			)
		}
		return nil, err
	}

	return &mapping.Account, nil
}
