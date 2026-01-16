package employee_repo

import (
	"context"
	employee_models "sharpnr-engine/modules/hr/employees/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// This file Implements all the Methods of EmployeeIdentityRepositoryInterface

type hrEmployeeRepo struct {
	db *gorm.DB
}

func NewEmployeeIdentityRepository(db *gorm.DB) *hrEmployeeRepo {
	return &hrEmployeeRepo{
		db: db,
	}
}

func (r *hrEmployeeRepo) Create(
	ctx context.Context,
	emp *employee_models.Employee,
) error {

	if emp.ID == uuid.Nil {
		emp.ID = uuid.New()
	}

	return r.db.WithContext(ctx).Create(emp).Error
}

func (r *hrEmployeeRepo) GetByID(
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

func (r *hrEmployeeRepo) GetByEmail(
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

func (r *hrEmployeeRepo) ListByOrg(
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
