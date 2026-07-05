package coreEmployeeRepo

import (
	"context"
	coreEmployeeModels "sharpnr-engine/core/employees/models"

	"time"

	"github.com/google/uuid"
)

func (r *employeeRepo) CreateEmployeeAuth(
	ctx context.Context,
	auth *coreEmployeeModels.EmployeeAuth,
) error {

	if auth.ID == uuid.Nil {
		auth.ID = uuid.New()
	}

	return r.conn(ctx).WithContext(ctx).Create(auth).Error
}

/* ---------------------------------- <***> --------------------------------- */

func (r *employeeRepo) GetAuthByEmployeeID(
	ctx context.Context,
	empID uuid.UUID,
) (*coreEmployeeModels.EmployeeAuth, error) {

	var auth coreEmployeeModels.EmployeeAuth
	if err := r.db.WithContext(ctx).
		Where("employee_id = ?", empID).
		First(&auth).Error; err != nil {
		return nil, err
	}

	return &auth, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *employeeRepo) GetByUsername(
	ctx context.Context,
	username string,
) (*coreEmployeeModels.EmployeeAuth, error) {

	var auth coreEmployeeModels.EmployeeAuth
	if err := r.db.WithContext(ctx).
		Where("username = ?", username).
		First(&auth).Error; err != nil {
		return nil, err
	}

	return &auth, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *employeeRepo) UpdateLastLogin(
	ctx context.Context,
	empID uuid.UUID,
	at time.Time,
) error {

	return r.db.WithContext(ctx).
		Model(&coreEmployeeModels.EmployeeAuth{}).
		Where("employee_id = ?", empID).
		Update("last_login_at", at).
		Error
}

/* ---------------------------------- <***> --------------------------------- */
