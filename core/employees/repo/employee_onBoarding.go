package coreEmployeeRepo

import (
	"context"
	coreEmployeeModels "sharpnr-engine/core/employees/models"

	"github.com/google/uuid"
)

func (r *employeeRepo) CreateEmployeeOnBoarding(
	ctx context.Context,
	onboarding *coreEmployeeModels.EmployeeOnboarding,
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
) (*coreEmployeeModels.EmployeeOnboarding, error) {

	var onboarding coreEmployeeModels.EmployeeOnboarding

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
		Model(&coreEmployeeModels.EmployeeOnboarding{}).
		Where("employee_id = ?", empID).
		Update("status", status).
		Error
}

/* ---------------------------------- <***> --------------------------------- */
