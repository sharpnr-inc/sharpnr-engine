package hr_employee_repo

import (
	"context"
	sharphr_models "sharpnr-engine/modules/hr/hrModels"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type hrEmployeeAuthRepo struct {
	db *gorm.DB
}

func NewEmployeeAuthRepository(db *gorm.DB) *hrEmployeeAuthRepo {
	return &hrEmployeeAuthRepo{
		db: db,
	}
}

func (r *hrEmployeeAuthRepo) Create(
	ctx context.Context,
	auth *sharphr_models.EmployeeAuth,
) error {

	if auth.ID == uuid.Nil {
		auth.ID = uuid.New()
	}

	return r.db.WithContext(ctx).Create(auth).Error
}

func (r *hrEmployeeAuthRepo) GetByEmployeeID(
	ctx context.Context,
	empID uuid.UUID,
) (*sharphr_models.EmployeeAuth, error) {

	var auth sharphr_models.EmployeeAuth
	if err := r.db.WithContext(ctx).
		Where("employee_id = ?", empID).
		First(&auth).Error; err != nil {
		return nil, err
	}

	return &auth, nil
}

func (r *hrEmployeeAuthRepo) GetByUsername(
	ctx context.Context,
	username string,
) (*sharphr_models.EmployeeAuth, error) {

	var auth sharphr_models.EmployeeAuth
	if err := r.db.WithContext(ctx).
		Where("username = ?", username).
		First(&auth).Error; err != nil {
		return nil, err
	}

	return &auth, nil
}

func (r *hrEmployeeAuthRepo) UpdateLastLogin(
	ctx context.Context,
	empID uuid.UUID,
	at time.Time,
) error {

	return r.db.WithContext(ctx).
		Model(&sharphr_models.EmployeeAuth{}).
		Where("employee_id = ?", empID).
		Update("last_login_at", at).
		Error
}
