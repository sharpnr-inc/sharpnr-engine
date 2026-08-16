package coreEmployeeRepo

import (
	"context"
	"fmt"
	coreEmployeeModels "sharpnr-engine/core/employees/models"
	"time"

	"github.com/google/uuid"
)

func (d *employeeRepo) CreateEmployeeProvision(ctx context.Context, orgID uuid.UUID, provisions []coreEmployeeModels.EmployeeProvisioning) error {
	if len(provisions) == 0 {
		return nil
	}

	// Uses TableName() defined on coreEmployeeModels.EmployeeProvisioning automatically
	if err := d.db.WithContext(ctx).Create(&provisions).Error; err != nil {
		return fmt.Errorf("failed to create employee provisioning records for org %s: %w", orgID, err)
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (d *employeeRepo) GetEmployeeProvioionByEmployeeID(ctx context.Context, orgID uuid.UUID, employeeID uuid.UUID) ([]coreEmployeeModels.EmployeeProvisioning, error) {
	var provisions []coreEmployeeModels.EmployeeProvisioning

	err := d.db.WithContext(ctx).
		Where("employee_id = ?", employeeID).
		Order("created_at DESC").
		Find(&provisions).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get employee provisioning records for employee %s (org %s): %w", employeeID, orgID, err)
	}

	return provisions, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (d *employeeRepo) AddConsumeTimeInProvisioningToken(ctx context.Context, orgID uuid.UUID, employeeID uuid.UUID, consumptionTime time.Time) error {
	result := d.db.WithContext(ctx).
		Model(&coreEmployeeModels.EmployeeProvisioning{}).
		Where("employee_id = ? AND consumed_at IS NULL", employeeID).
		Update("consumed_at", consumptionTime)

	if result.Error != nil {
		return fmt.Errorf("failed to mark provisioning token as consumed for employee %s (org %s): %w", employeeID, orgID, result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no unconsumed provisioning token found for employee %s", employeeID)
	}

	return nil
}
