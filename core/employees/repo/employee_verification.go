package coreEmployeeRepo

import (
	"context"
	coreEmployeeModels "sharpnr-engine/core/employees/models"

	"time"

	"github.com/google/uuid"
)

func (r *employeeRepo) CreateEmplVerification(
	ctx context.Context,
	v *coreEmployeeModels.EmployeeVerification,
) error {

	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}

	return r.db.WithContext(ctx).Create(v).Error
}

func (r *employeeRepo) GetByTokenHash(
	ctx context.Context,
	hash string,
) (*coreEmployeeModels.EmployeeVerification, error) {

	var verification coreEmployeeModels.EmployeeVerification

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

func (r *employeeRepo) MarkVerified(
	ctx context.Context,
	id uuid.UUID,
	at time.Time,
) error {

	return r.db.
		WithContext(ctx).
		Model(&coreEmployeeModels.EmployeeVerification{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"verified_at": at,
		}).
		Error
}
