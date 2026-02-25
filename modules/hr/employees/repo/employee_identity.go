package employee_repo

import (
	"context"
	employee_models "sharpnr-engine/modules/hr/employees/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// This file Implements all the Methods of EmployeeIdentityRepositoryInterface

type employeeRepo struct {
	db *gorm.DB
}

func NewEmployeeIdentityRepository(db *gorm.DB) *employeeRepo {
	return &employeeRepo{
		db: db,
	}
}

func (r *employeeRepo) Create(
	ctx context.Context,
	emp *employee_models.Employee,
) error {

	if emp.ID == uuid.Nil {
		emp.ID = uuid.New()
	}

	return r.db.WithContext(ctx).Create(emp).Error
}

func (r *employeeRepo) GetByID(
	ctx context.Context,
	empID uuid.UUID,
) (*employee_models.Employee, error) {

	var emp employee_models.Employee

	err := r.db.
		WithContext(ctx).
		Where("id = ?", empID).
		First(&emp).
		Error

	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func (r *employeeRepo) GetByEmail(
	ctx context.Context,
	email string,
) (*employee_models.Employee, error) {

	var emp employee_models.Employee

	err := r.db.
		WithContext(ctx).
		Where("email = ?", email).
		First(&emp).
		Error

	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func (r *employeeRepo) ExistsByEmail(
	ctx context.Context,
	email string,
) (bool, error) {

	var count int64

	err := r.db.WithContext(ctx).
		Model(&employee_models.Employee{}).
		Where("email = ?", email).
		Count(&count).
		Error

	return count > 0, err
}

func (r *employeeRepo) ListByOrg(
	ctx context.Context,
	orgID uuid.UUID,
) ([]*employee_models.Employee, error) {

	var employees []*employee_models.Employee

	err := r.db.
		WithContext(ctx).
		Where("organization_id = ?", orgID).
		Find(&employees).
		Error

	if err != nil {
		return nil, err
	}

	return employees, nil
}
