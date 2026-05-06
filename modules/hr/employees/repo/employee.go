package employee_repo

import (
	"context"
	employee_models "sharpnr-engine/modules/hr/employees/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *employeeRepo) Update(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, updates map[string]interface{}) error {
	// 1. Ensure we are scoped to the correct Organization and Employee
	result := r.db.WithContext(ctx).
		Model(&employee_models.Employee{}).
		Where("id = ? AND organization_id = ?", emplID, orgID).
		Updates(updates)

	// 2. Handle database-level errors
	if result.Error != nil {
		return result.Error
	}

	// 3. Handle cases where the record doesn't exist or orgID mismatch
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
