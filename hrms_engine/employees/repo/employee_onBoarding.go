package employee_repo

import (
	"context"
	employee_models "sharpnr-engine/hrms_engine/employees/models"

	"github.com/google/uuid"
)

func (r *employeeRepo) CreateEmployeeOnBoarding(
	ctx context.Context,
	onboarding *employee_models.EmployeeOnboarding,
) error {

	if onboarding.ID == uuid.Nil {
		onboarding.ID = uuid.New()
	}

	return r.conn(ctx).WithContext(ctx).Create(onboarding).Error
}

/* ---------------------------------- <***> --------------------------------- */

func (r *employeeRepo) GetEmplOnBoardingByEmployeeID(
	ctx context.Context,
	empID uuid.UUID,
) (*employee_models.EmployeeOnboarding, error) {

	var onboarding employee_models.EmployeeOnboarding

	err := r.db.
		WithContext(ctx).
		Where("employee_id = ?", empID).
		First(&onboarding).
		Error

	if err != nil {
		return nil, err
	}

	return &onboarding, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *employeeRepo) UpdateStatus(
	ctx context.Context,
	empID uuid.UUID,
	status string,
) error {

	return r.db.
		WithContext(ctx).
		Model(&employee_models.EmployeeOnboarding{}).
		Where("employee_id = ?", empID).
		Update("status", status).
		Error
}

/* ---------------------------------- <***> --------------------------------- */
