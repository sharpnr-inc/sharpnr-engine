package employee_repo

import (
	"context"
	employee_models "sharpnr-engine/hrms_engine/employees/models"
	"time"

	"github.com/google/uuid"
)

func (r *employeeRepo) CreateEmplVerification(
	ctx context.Context,
	v *employee_models.EmployeeVerification,
) error {

	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}

	return r.db.WithContext(ctx).Create(v).Error
}

func (r *employeeRepo) GetByTokenHash(
	ctx context.Context,
	hash string,
) (*employee_models.EmployeeVerification, error) {

	var verification employee_models.EmployeeVerification

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
		Model(&employee_models.EmployeeVerification{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"verified_at": at,
		}).
		Error
}
