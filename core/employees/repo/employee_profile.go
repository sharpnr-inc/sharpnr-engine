package employee_repo

import (
	"context"
	employee_models "sharpnr-engine/core/employees/models"

	"github.com/google/uuid"
)

// CreateEmployeeSystemAccess implements [EmployeeProfileRepositoryInterface].
func (r *employeeRepo) CreateEmployeeSystemAccess(ctx context.Context, access *employee_models.EmployeeSystemAccess) (*employee_models.EmployeeSystemAccess, error) {
	panic("unimplemented")
}

// GetEmployeeSystemAccessByEmployeeID implements [EmployeeProfileRepositoryInterface].
func (r *employeeRepo) GetEmployeeSystemAccessByEmployeeID(ctx context.Context, empID uuid.UUID) ([]*employee_models.EmployeeSystemAccess, error) {
	panic("unimplemented")
}

// IsEmployeeHaveOnlineAccess implements [EmployeeProfileRepositoryInterface].
func (r *employeeRepo) IsEmployeeHaveOnlineAccess(ctx context.Context, empID uuid.UUID) (bool, error) {
	panic("unimplemented")
}

func (r *employeeRepo) CreateEmployeeProfile(
	ctx context.Context,
	profile *employee_models.EmployeeProfile,
) error {

	if profile.ID == uuid.Nil {
		profile.ID = uuid.New()
	}

	return r.db.WithContext(ctx).Create(profile).Error
}

func (r *employeeRepo) GetProfileByEmployeeID(
	ctx context.Context,
	empID uuid.UUID,
) (*employee_models.EmployeeProfile, error) {

	var profile employee_models.EmployeeProfile

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
	profile *employee_models.EmployeeProfile,
) error {

	return r.db.
		WithContext(ctx).
		Save(profile).
		Error
}
