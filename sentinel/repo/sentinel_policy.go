package sentinel_repo

import (
	"context"

	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
)

// CreatePolicy persists a Cedar policy row. ID is application-generated if unset.
func (r *sentinelRepo) CreatePolicy(ctx context.Context, policy *sentinelengine_models.SentinelPolicy) error {
	if policy.ID == uuid.Nil {
		policy.ID = uuid.New()
	}
	return r.conn(ctx).Create(policy).Error
}

// GetPolicyByID loads a policy by its primary key.
func (r *sentinelRepo) GetPolicyByID(ctx context.Context, policyID uuid.UUID) (sentinelengine_models.SentinelPolicy, error) {
	var policy sentinelengine_models.SentinelPolicy
	err := r.conn(ctx).Where("id = ?", policyID).First(&policy).Error
	return policy, err
}

// GetPolicyByName loads a policy by its (org, name).
func (r *sentinelRepo) GetPolicyByName(ctx context.Context, organizationID uuid.UUID, policyName string) (sentinelengine_models.SentinelPolicy, error) {
	var policy sentinelengine_models.SentinelPolicy
	err := r.conn(ctx).
		Where("organization_id = ? AND name = ?", organizationID, policyName).
		First(&policy).Error
	return policy, err
}

// ListPoliciesByOrganizationID returns an org's policies, newest first. An empty
// scope returns every bucket; otherwise it filters to that scope.
func (r *sentinelRepo) ListPoliciesByOrganizationID(ctx context.Context, organizationID uuid.UUID, scope sentinelengine_models.PolicyScope) ([]*sentinelengine_models.SentinelPolicy, error) {
	q := r.conn(ctx).Where("organization_id = ?", organizationID)
	if scope != "" {
		q = q.Where("scope = ?", scope)
	}

	var rows []*sentinelengine_models.SentinelPolicy
	err := q.Order("created_at DESC").Find(&rows).Error
	return rows, err
}

// UpdatePolicy persists all columns of the given policy by its primary key.
func (r *sentinelRepo) UpdatePolicy(ctx context.Context, policy *sentinelengine_models.SentinelPolicy) error {
	return r.conn(ctx).Save(policy).Error
}

// EnablePolicy flips a policy's enabled flag on.
func (r *sentinelRepo) EnablePolicy(ctx context.Context, policyID uuid.UUID) error {
	return r.conn(ctx).
		Model(&sentinelengine_models.SentinelPolicy{}).
		Where("id = ?", policyID).
		Update("enabled", true).Error
}

// DisablePolicy flips a policy's enabled flag off (disable without deleting).
func (r *sentinelRepo) DisablePolicy(ctx context.Context, policyID uuid.UUID) error {
	return r.conn(ctx).
		Model(&sentinelengine_models.SentinelPolicy{}).
		Where("id = ?", policyID).
		Update("enabled", false).Error
}

// DeletePolicy soft-deletes a policy.
func (r *sentinelRepo) DeletePolicy(ctx context.Context, policyID uuid.UUID) error {
	return r.conn(ctx).Where("id = ?", policyID).Delete(&sentinelengine_models.SentinelPolicy{}).Error
}

// PolicyExists reports whether a policy with that name already exists in the org.
func (r *sentinelRepo) PolicyExists(ctx context.Context, organizationID uuid.UUID, policyName string) (bool, error) {
	var count int64
	err := r.conn(ctx).
		Model(&sentinelengine_models.SentinelPolicy{}).
		Where("organization_id = ? AND name = ?", organizationID, policyName).
		Count(&count).Error
	return count > 0, err
}
