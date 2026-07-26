package coreEmployeeRepo

import (
	"context"
	coreEmployeeModels "sharpnr-engine/core/employees/models"

	"github.com/google/uuid"
)

func (r *employeeRepo) CreateEmployeeProfile(
	ctx context.Context,
	profile *coreEmployeeModels.EmployeeProfile,
) error {

	if profile.ID == uuid.Nil {
		profile.ID = uuid.New()
	}

	return r.db.WithContext(ctx).Create(profile).Error
}

func (r *employeeRepo) GetProfileByEmployeeID(
	ctx context.Context,
	empID uuid.UUID,
) (*coreEmployeeModels.EmployeeProfile, error) {

	var profile coreEmployeeModels.EmployeeProfile

	err := r.db.
		WithContext(ctx).
		Where("employee_id = ?", empID).
		First(&profile).
		Error

	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *employeeRepo) UpdateEmployeeProfile(
	ctx context.Context,
	profile *coreEmployeeModels.EmployeeProfile,
) error {

	return r.db.
		WithContext(ctx).
		Save(profile).
		Error
}
