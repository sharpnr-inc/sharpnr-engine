package hr_employee_repo

import (
	"context"
	sharphr_models "sharpnr-engine/modules/hr/hrModels"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// This file implements all the methods of EmployeeVerificationRepositoryInterface
/*
type EmployeeVerificationRepositoryInterface interface {
	Create(ctx context.Context, v *sharphr_models.EmployeeVerification) error
	GetByTokenHash(ctx context.Context, hash string) (*sharphr_models.EmployeeVerification, error)
	MarkVerified(ctx context.Context, id uuid.UUID, at time.Time) error
}
*/
// employeeVerificationRepo implements EmployeeVerificationRepositoryInterface
type employeeVerificationRepo struct {
	db *gorm.DB
}

func NewEmployeeVerificationRepo(db *gorm.DB) *employeeVerificationRepo {
	return &employeeVerificationRepo{
		db: db,
	}
}

func (r *employeeVerificationRepo) Create(
	ctx context.Context,
	v *sharphr_models.EmployeeVerification,
) error {

	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}

	return r.db.WithContext(ctx).Create(v).Error
}

func (r *employeeVerificationRepo) GetByTokenHash(
	ctx context.Context,
	hash string,
) (*sharphr_models.EmployeeVerification, error) {

	var verification sharphr_models.EmployeeVerification

	err := r.db.
		WithContext(ctx).
		Where("token_hash = ?", hash).
		First(&verification).
		Error

	if err != nil {
		return nil, err
	}

	return &verification, nil
}

func (r *employeeVerificationRepo) MarkVerified(
	ctx context.Context,
	id uuid.UUID,
	at time.Time,
) error {

	return r.db.
		WithContext(ctx).
		Model(&sharphr_models.EmployeeVerification{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"verified_at": at,
		}).
		Error
}

var _ EmployeeVerificationRepositoryInterface = (*employeeVerificationRepo)(nil)
