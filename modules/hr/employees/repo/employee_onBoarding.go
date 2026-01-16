package employee_repo

import (
	"context"
	employee_models "sharpnr-engine/modules/hr/employees/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// *This File Implements all the methods from the EmployeeOnboardingRepositoryInterface
/*
type EmployeeOnboardingRepositoryInterface interface {
	Create(ctx context.Context, onboarding *employee_models.EmployeeOnboarding) error
	GetByEmployeeID(ctx context.Context, empID uuid.UUID) (*employee_models.EmployeeOnboarding, error)
	UpdateStatus(ctx context.Context, empID uuid.UUID, status string) error
}
*/
// employeeOnboardingRepo implements EmployeeOnboardingRepositoryInterface
type employeeOnboardingRepo struct {
	db *gorm.DB
}

func NewEmployeeOnboardingRepository(db *gorm.DB) *employeeOnboardingRepo {
	return &employeeOnboardingRepo{
		db: db,
	}
}

func (r *employeeOnboardingRepo) Create(
	ctx context.Context,
	onboarding *employee_models.EmployeeOnboarding,
) error {

	if onboarding.ID == uuid.Nil {
		onboarding.ID = uuid.New()
	}

	return r.db.WithContext(ctx).Create(onboarding).Error
}

func (r *employeeOnboardingRepo) GetByEmployeeID(
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

func (r *employeeOnboardingRepo) UpdateStatus(
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

var _ EmployeeOnboardingRepositoryInterface = (*employeeOnboardingRepo)(nil)
