package hr_employee_repo

import (
	"context"
	sharphr_models "sharpnr-engine/modules/hr/hrModels"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type employeePrivilegePolicyRepo struct {
	db *gorm.DB
}

func NewEmployeePrivilegePolicyRepository(
	db *gorm.DB,
) *employeePrivilegePolicyRepo {
	return &employeePrivilegePolicyRepo{db: db}
}

func (r *employeePrivilegePolicyRepo) Assign(
	ctx context.Context,
	assignment *sharphr_models.EmployeePrivilegePolicy,
) error {

	if assignment.ID == uuid.Nil {
		assignment.ID = uuid.New()
	}

	if assignment.AssignedAt.IsZero() {
		assignment.AssignedAt = time.Now()
	}

	return r.db.WithContext(ctx).Create(assignment).Error
}

func (r *employeePrivilegePolicyRepo) ListByEmployeeID(
	ctx context.Context,
	employeeID uuid.UUID,
) ([]*sharphr_models.EmployeePrivilegePolicy, error) {

	var assignments []*sharphr_models.EmployeePrivilegePolicy

	if err := r.db.
		WithContext(ctx).
		Where("employee_id = ?", employeeID).
		Find(&assignments).
		Error; err != nil {
		return nil, err
	}

	return assignments, nil
}

func (r *employeePrivilegePolicyRepo) RevokeByEmployeeID(
	ctx context.Context,
	employeeID uuid.UUID,
) error {

	return r.db.
		WithContext(ctx).
		Where("employee_id = ?", employeeID).
		Delete(&sharphr_models.EmployeePrivilegePolicy{}).
		Error
}
