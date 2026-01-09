package hr_employee_repo

import (
	"context"
	sharphr_models "sharpnr-engine/modules/hr/hrModels"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// This file Implements the Methods for all the from EmployeePolicyRepositoryInterface

/*
type EmployeePolicyRepositoryInterface interface {
	Create(ctx context.Context, policy *sharphr_models.EmployeePolicy) error
	GetByID(ctx context.Context, id uuid.UUID) (*sharphr_models.EmployeePolicy, error)
	GetByCode(ctx context.Context, orgID uuid.UUID, code string) (*sharphr_models.EmployeePolicy, error)
	ListByOrganization(ctx context.Context, orgID uuid.UUID) ([]*sharphr_models.EmployeePolicy, error)
}
*/
// Implements the Methods for all the from EmployeePolicyRepositoryInterface
type empPolicyRepo struct {
	db *gorm.DB
}

func NewEmployeePolicyRepo(db *gorm.DB) *empPolicyRepo {
	return &empPolicyRepo{
		db: db,
	}
}

func (r *empPolicyRepo) Create(
	ctx context.Context,
	policy *sharphr_models.EmployeePolicy,
) error {

	if policy.ID == uuid.Nil {
		policy.ID = uuid.New()
	}

	return r.db.WithContext(ctx).Create(policy).Error
}

func (r *empPolicyRepo) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*sharphr_models.EmployeePolicy, error) {

	var policy sharphr_models.EmployeePolicy

	if err := r.db.
		WithContext(ctx).
		Where("id = ?", id).
		First(&policy).
		Error; err != nil {
		return nil, err
	}

	return &policy, nil
}

func (r *empPolicyRepo) GetByCode(
	ctx context.Context,
	orgID uuid.UUID,
	code string,
) (*sharphr_models.EmployeePolicy, error) {

	var policy sharphr_models.EmployeePolicy

	if err := r.db.
		WithContext(ctx).
		Where("organization_id = ? AND code = ?", orgID, code).
		First(&policy).
		Error; err != nil {
		return nil, err
	}

	return &policy, nil
}

func (r *empPolicyRepo) ListByOrganization(
	ctx context.Context,
	orgID uuid.UUID,
) ([]*sharphr_models.EmployeePolicy, error) {

	var policies []*sharphr_models.EmployeePolicy

	if err := r.db.
		WithContext(ctx).
		Where("organization_id = ?", orgID).
		Find(&policies).
		Error; err != nil {
		return nil, err
	}

	return policies, nil
}
