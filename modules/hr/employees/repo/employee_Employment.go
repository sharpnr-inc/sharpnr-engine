package hr_employee_repo

import (
	"context"
	sharphr_models "sharpnr-engine/modules/hr/hrModels"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// This file implements all method from EmployeeEmploymentRepositoryInterface

/*
type EmployeeEmploymentRepositoryInterface interface {
	Create(ctx context.Context, emp *sharphr_models.EmployeeEmployment) error
	GetByEmployeeID(ctx context.Context, empID uuid.UUID) (*sharphr_models.EmployeeEmployment, error)
}
*/
//
type employeeEmploymentRepo struct {
	db *gorm.DB
}

func NewEmployeeEmploymentRepo(db *gorm.DB) *employeeEmploymentRepo {
	return &employeeEmploymentRepo{
		db: db,
	}
}

func (r *employeeEmploymentRepo) Create(
	ctx context.Context,
	emp *sharphr_models.EmployeeEmployment,
) error {

	if emp.ID == uuid.Nil {
		emp.ID = uuid.New()
	}

	return r.db.WithContext(ctx).Create(emp).Error
}

func (r *employeeEmploymentRepo) GetByEmployeeID(
	ctx context.Context,
	empID uuid.UUID,
) (*sharphr_models.EmployeeEmployment, error) {

	var employment sharphr_models.EmployeeEmployment

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
