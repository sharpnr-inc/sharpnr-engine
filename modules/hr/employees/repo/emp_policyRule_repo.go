package hr_employee_repo

import (
	"context"
	sharphr_models "sharpnr-engine/modules/hr/hrModels"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// employeePolicyRuleRepo implements EmployeePolicyRuleRepositoryInterface
type employeePolicyRuleRepo struct {
	db *gorm.DB
}

func NewEmployeePolicyRuleRepository(
	db *gorm.DB,
) *employeePolicyRuleRepo {
	return &employeePolicyRuleRepo{db: db}
}

func (r *employeePolicyRuleRepo) Create(
	ctx context.Context,
	rule *sharphr_models.EmployeePolicyRule,
) error {

	if rule.ID == uuid.Nil {
		rule.ID = uuid.New()
	}

	return r.db.WithContext(ctx).Create(rule).Error
}

func (r *employeePolicyRuleRepo) ListByPolicyID(
	ctx context.Context,
	policyID uuid.UUID,
) ([]*sharphr_models.EmployeePolicyRule, error) {

	var rules []*sharphr_models.EmployeePolicyRule

	if err := r.db.
		WithContext(ctx).
		Where("employee_policy_id = ?", policyID).
		Find(&rules).
		Error; err != nil {
		return nil, err
	}

	return rules, nil
}

func (r *employeePolicyRuleRepo) DeleteByPolicyID(
	ctx context.Context,
	policyID uuid.UUID,
) error {

	return r.db.
		WithContext(ctx).
		Where("employee_policy_id = ?", policyID).
		Delete(&sharphr_models.EmployeePolicyRule{}).
		Error
}
