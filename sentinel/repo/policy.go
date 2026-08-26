package sentinel_repo

import (
	"context"

	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
)

// Policies live in ONE table; who they reach is decided by the attachment
// registers (see attachments.go). Nothing here knows about distribution.
//
// Every method takes organizationID and filters on it, so a caller can never
// reach another tenant's policy by id alone — the service layer no longer needs
// to re-check OrganizationID after a fetch.

// CreatePolicy persists a Cedar policy row. ID is application-generated if unset.
func (r *sentinelRepo) CreatePolicy(ctx context.Context, policy *sentinelengine_models.SentinelPolicy) error {
	if policy.ID == uuid.Nil {
		policy.ID = uuid.New()
	}
	return r.conn(ctx).Create(policy).Error
}

/* ---------------------------------- <***> --------------------------------- */

// GetPolicy loads one of the org's policies. A policy belonging to another org
// is reported as not found, never returned.
func (r *sentinelRepo) GetPolicy(ctx context.Context, organizationID, policyID uuid.UUID) (sentinelengine_models.SentinelPolicy, error) {
	var policy sentinelengine_models.SentinelPolicy
	err := r.conn(ctx).
		Where("organization_id = ? AND id = ?", organizationID, policyID).
		First(&policy).Error
	return policy, err
}

/* ---------------------------------- <***> --------------------------------- */

// GetPolicyByName loads a policy by its (org, name) pair.
func (r *sentinelRepo) GetPolicyByName(ctx context.Context, organizationID uuid.UUID, name string) (sentinelengine_models.SentinelPolicy, error) {
	var policy sentinelengine_models.SentinelPolicy
	err := r.conn(ctx).
		Where("organization_id = ? AND name = ?", organizationID, name).
		First(&policy).Error
	return policy, err
}

/* ---------------------------------- <***> --------------------------------- */

// ListPolicies returns every policy in the org, newest first. Policies are no
// longer bucketed by scope — what a policy reaches is read from the attachment
// registers, not from a column on the policy itself.
func (r *sentinelRepo) ListPolicies(ctx context.Context, organizationID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
	var rows []*sentinelengine_models.SentinelPolicy
	err := r.conn(ctx).
		Where("organization_id = ?", organizationID).
		Order("created_at DESC").
		Find(&rows).Error
	return rows, err
}

/* ---------------------------------- <***> --------------------------------- */

// UpdatePolicy persists the given policy's columns, scoped to the org so a
// mismatched OrganizationID updates nothing instead of writing across tenants.
func (r *sentinelRepo) UpdatePolicy(ctx context.Context, organizationID uuid.UUID, policy *sentinelengine_models.SentinelPolicy) error {
	return r.conn(ctx).
		Model(&sentinelengine_models.SentinelPolicy{}).
		Where("organization_id = ? AND id = ?", organizationID, policy.ID).
		Updates(map[string]any{
			"name":        policy.Name,
			"description": policy.Description,
			"policy":      policy.Policy,
			"policy_json": policy.PolicyJson,
			"enabled":     policy.Enabled,
		}).Error
}

/* ---------------------------------- <***> --------------------------------- */

// SetPolicyEnabled turns a policy on or off without deleting it. `enabled` is a
// switch independent of attachment: a disabled policy keeps its registers and
// starts applying again the moment it is re-enabled.
func (r *sentinelRepo) SetPolicyEnabled(ctx context.Context, organizationID, policyID uuid.UUID, enabled bool) error {
	return r.conn(ctx).
		Model(&sentinelengine_models.SentinelPolicy{}).
		Where("organization_id = ? AND id = ?", organizationID, policyID).
		Update("enabled", enabled).Error
}

/* ---------------------------------- <***> --------------------------------- */

// DeletePolicy soft-deletes a policy. The attachment registers carry a real FK
// with ON DELETE CASCADE, but that fires only on a hard delete — a soft-deleted
// policy simply stops being collected by GetAllPoliciesOfUser.
func (r *sentinelRepo) DeletePolicy(ctx context.Context, organizationID, policyID uuid.UUID) error {
	return r.conn(ctx).
		Where("organization_id = ? AND id = ?", organizationID, policyID).
		Delete(&sentinelengine_models.SentinelPolicy{}).Error
}

/* ---------------------------------- <***> --------------------------------- */

// PolicyExists reports whether a policy with that name already exists in the org.
func (r *sentinelRepo) PolicyExists(ctx context.Context, organizationID uuid.UUID, name string) (bool, error) {
	var count int64
	err := r.conn(ctx).
		Model(&sentinelengine_models.SentinelPolicy{}).
		Where("organization_id = ? AND name = ?", organizationID, name).
		Count(&count).Error
	return count > 0, err
}

/* ---------------------------------- <***> --------------------------------- */
