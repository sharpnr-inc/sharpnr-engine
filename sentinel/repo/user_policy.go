package sentinel_repo

import (
	"context"

	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Note: the UserPolicy model names the user column employee_id; the userID
// arguments below map to it.

// AssignPolicyToUser grants a policy directly to a user. Idempotent.
func (r *sentinelRepo) AssignPolicyToUser(ctx context.Context, organizationID, userID, policyID uuid.UUID) error {
	assigned, err := r.IsPolicyAssignedToUser(ctx, userID, policyID)
	if err != nil {
		return err
	}
	if assigned {
		return nil
	}
	edge := &sentinelengine_models.UserPolicy{
		ID:               uuid.New(),
		OrganizationID:   organizationID,
		EmployeeID:       userID,
		SentinelPolicyID: policyID,
	}
	return r.conn(ctx).Create(edge).Error
}

// RemovePolicyFromUser revokes a directly-granted policy (soft delete of the edge).
func (r *sentinelRepo) RemovePolicyFromUser(ctx context.Context, organizationID, userID, policyID uuid.UUID) error {
	return r.conn(ctx).
		Where("organization_id = ? AND employee_id = ? AND sentinel_policy_id = ?", organizationID, userID, policyID).
		Delete(&sentinelengine_models.UserPolicy{}).Error
}

// GetPoliciesAssignedToUser returns the direct-grant edges for a user.
func (r *sentinelRepo) GetPoliciesAssignedToUser(ctx context.Context, userID uuid.UUID) ([]*sentinelengine_models.UserPolicy, error) {
	var edges []*sentinelengine_models.UserPolicy
	err := r.conn(ctx).Where("employee_id = ?", userID).Find(&edges).Error
	return edges, err
}

// GetUsersAssignedToPolicy returns the direct-grant edges for a policy in an org.
func (r *sentinelRepo) GetUsersAssignedToPolicy(ctx context.Context, organizationID, policyID uuid.UUID) ([]*sentinelengine_models.UserPolicy, error) {
	var edges []*sentinelengine_models.UserPolicy
	err := r.conn(ctx).
		Where("organization_id = ? AND sentinel_policy_id = ?", organizationID, policyID).
		Find(&edges).Error
	return edges, err
}

// IsPolicyAssignedToUser reports whether a user has a direct grant of a policy.
func (r *sentinelRepo) IsPolicyAssignedToUser(ctx context.Context, userID, policyID uuid.UUID) (bool, error) {
	var count int64
	err := r.conn(ctx).
		Model(&sentinelengine_models.UserPolicy{}).
		Where("employee_id = ? AND sentinel_policy_id = ?", userID, policyID).
		Count(&count).Error
	return count > 0, err
}

// ReplacePoliciesAssignedToUser sets a user's direct policies to exactly
// policyIDs, in one tx.
func (r *sentinelRepo) ReplacePoliciesAssignedToUser(ctx context.Context, organizationID, userID uuid.UUID, policyIDs []uuid.UUID) error {
	return r.conn(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Where("organization_id = ? AND employee_id = ?", organizationID, userID).
			Delete(&sentinelengine_models.UserPolicy{}).Error; err != nil {
			return err
		}
		for _, policyID := range policyIDs {
			edge := &sentinelengine_models.UserPolicy{
				ID:               uuid.New(),
				OrganizationID:   organizationID,
				EmployeeID:       userID,
				SentinelPolicyID: policyID,
			}
			if err := tx.Create(edge).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
