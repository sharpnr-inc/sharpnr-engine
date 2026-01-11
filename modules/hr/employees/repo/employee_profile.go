package hr_employee_repo

import (
	"context"
	sharphr_models "sharpnr-engine/modules/hr/hrModels"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// This file implements the all methods of EmployeeProfileRepositoryInterface

/*
type EmployeeProfileRepositoryInterface interface {
	Create(ctx context.Context, profile *sharphr_models.EmployeeProfile) error
	GetByEmployeeID(ctx context.Context, empID uuid.UUID) (*sharphr_models.EmployeeProfile, error)
	Update(ctx context.Context, profile *sharphr_models.EmployeeProfile) error
}
*/
//
type employeeProfileRepo struct {
	db *gorm.DB
}

func NewEmployeeProfileRepo(db *gorm.DB) *employeeProfileRepo {
	return &employeeProfileRepo{
		db: db,
	}
}

func (r *employeeProfileRepo) Create(
	ctx context.Context,
	profile *sharphr_models.EmployeeProfile,
) error {

	if profile.ID == uuid.Nil {
		profile.ID = uuid.New()
	}

	return r.db.WithContext(ctx).Create(profile).Error
}

func (r *employeeProfileRepo) GetByEmployeeID(
	ctx context.Context,
	empID uuid.UUID,
) (*sharphr_models.EmployeeProfile, error) {

	var profile sharphr_models.EmployeeProfile

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

func (r *employeeProfileRepo) Update(
	ctx context.Context,
	profile *sharphr_models.EmployeeProfile,
) error {

	return r.db.
		WithContext(ctx).
		Save(profile).
		Error
}
