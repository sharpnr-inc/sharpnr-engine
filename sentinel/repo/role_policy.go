package sentinel_repo

import (
	"context"

	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AssignPolicyToRole attaches a policy to a role. Idempotent: if the (role,
// policy) edge is already active it is a no-op. ID is app-generated if unset.
func (r *sentinelRepo) AssignPolicyToRole(ctx context.Context, rolePolicy *sentinelengine_models.RolePolicies) error {
	assigned, err := r.IsPolicyAssignedToRole(ctx, rolePolicy.RoleID, rolePolicy.SentinelPolicyID)
	if err != nil {
		return err
	}
	if assigned {
		return nil
	}
	if rolePolicy.ID == uuid.Nil {
		rolePolicy.ID = uuid.New()
	}
	return r.conn(ctx).Create(rolePolicy).Error
}

// RemovePolicyFromRole detaches a policy from a role (soft delete of the edge).
func (r *sentinelRepo) RemovePolicyFromRole(ctx context.Context, organizationID, roleID, policyID uuid.UUID) error {
	return r.conn(ctx).
		Where("organization_id = ? AND role_id = ? AND sentinel_policy_id = ?", organizationID, roleID, policyID).
		Delete(&sentinelengine_models.RolePolicies{}).Error
}

// ListPoliciesAssignedToRole returns the policies attached to a role.
func (r *sentinelRepo) ListPoliciesAssignedToRole(ctx context.Context, organizationID, roleID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
	var policies []*sentinelengine_models.SentinelPolicy
	err := r.conn(ctx).
		Table("sentinel_engine.sentinel_policies AS p").
		Joins("JOIN sentinel_engine.role_policies AS rp ON rp.sentinel_policy_id = p.id").
		Where("rp.organization_id = ? AND rp.role_id = ? AND rp.deleted_at IS NULL AND p.deleted_at IS NULL", organizationID, roleID).
		Order("p.created_at DESC").
		Find(&policies).Error
	return policies, err
}

// ListRolesAssignedToPolicy returns the roles a policy is attached to.
func (r *sentinelRepo) ListRolesAssignedToPolicy(ctx context.Context, organizationID, policyID uuid.UUID) ([]*sentinelengine_models.Role, error) {
	var roles []*sentinelengine_models.Role
	err := r.conn(ctx).
		Table("sentinel_engine.roles AS ro").
		Joins("JOIN sentinel_engine.role_policies AS rp ON rp.role_id = ro.id").
		Where("rp.organization_id = ? AND rp.sentinel_policy_id = ? AND rp.deleted_at IS NULL AND ro.deleted_at IS NULL", organizationID, policyID).
		Order("ro.created_at DESC").
		Find(&roles).Error
	return roles, err
}

// IsPolicyAssignedToRole reports whether an active (role, policy) edge exists.
func (r *sentinelRepo) IsPolicyAssignedToRole(ctx context.Context, roleID, policyID uuid.UUID) (bool, error) {
	var count int64
	err := r.conn(ctx).
		Model(&sentinelengine_models.RolePolicies{}).
		Where("role_id = ? AND sentinel_policy_id = ?", roleID, policyID).
		Count(&count).Error
	return count > 0, err
}

// ReplacePoliciesAssignedToRole sets a role's attached policies to exactly
// policyIDs: detach all current edges, then attach the given set, in one tx.
func (r *sentinelRepo) ReplacePoliciesAssignedToRole(ctx context.Context, organizationID, roleID uuid.UUID, policyIDs []uuid.UUID) error {
	return r.conn(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Where("organization_id = ? AND role_id = ?", organizationID, roleID).
			Delete(&sentinelengine_models.RolePolicies{}).Error; err != nil {
			return err
		}
		for _, policyID := range policyIDs {
			edge := &sentinelengine_models.RolePolicies{
				ID:               uuid.New(),
				OrganizationID:   organizationID,
				RoleID:           roleID,
				SentinelPolicyID: policyID,
			}
			if err := tx.Create(edge).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
