package employee_repo

import (
	"context"
	employee_models "sharpnr-engine/sharp_hrms_engine/employees/models"

	"github.com/google/uuid"
)

func (r *employeeRepo) CreateEmployement(
	ctx context.Context,
	emp *employee_models.EmployeeEmployment,
) error {

	if emp.ID == uuid.Nil {
		emp.ID = uuid.New()
	}

	return r.db.WithContext(ctx).Create(emp).Error
}

func (r *employeeRepo) GetByEmployeeID(
	ctx context.Context,
	empID uuid.UUID,
) (*employee_models.EmployeeEmployment, error) {

	var employment employee_models.EmployeeEmployment

	err := r.db.
		WithContext(ctx).
		Where("employee_id = ?", empID).
		First(&employment).
		Error

	if err != nil {
		return nil, err
	}

	return &employment, nil
}
