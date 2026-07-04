package sentinel_repo

import (
	"context"

	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AssignPolicyToGroup attaches a policy to a group. Idempotent: if the (group,
// policy) edge is already active it is a no-op. ID is app-generated if unset.
func (r *sentinelRepo) AssignPolicyToGroup(ctx context.Context, groupPolicy *sentinelengine_models.GroupsSentinelPolicies) error {
	assigned, err := r.IsPolicyAssignedToGroup(ctx, groupPolicy.GroupID, groupPolicy.SentinelPolicyID)
	if err != nil {
		return err
	}
	if assigned {
		return nil
	}
	if groupPolicy.ID == uuid.Nil {
		groupPolicy.ID = uuid.New()
	}
	return r.conn(ctx).Create(groupPolicy).Error
}

// RemovePolicyFromGroup detaches a policy from a group (soft delete of the edge).
func (r *sentinelRepo) RemovePolicyFromGroup(ctx context.Context, organizationID, groupID, policyID uuid.UUID) error {
	return r.conn(ctx).
		Where("organization_id = ? AND group_id = ? AND sentinel_policy_id = ?", organizationID, groupID, policyID).
		Delete(&sentinelengine_models.GroupsSentinelPolicies{}).Error
}

// ListPoliciesAssignedToGroup returns the policies attached to a group.
func (r *sentinelRepo) ListPoliciesAssignedToGroup(ctx context.Context, organizationID, groupID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
	var policies []*sentinelengine_models.SentinelPolicy
	err := r.conn(ctx).
		Table("sentinel_engine.sentinel_policies AS p").
		Joins("JOIN sentinel_engine.groups_sentinel_policies AS gp ON gp.sentinel_policy_id = p.id").
		Where("gp.organization_id = ? AND gp.group_id = ? AND gp.deleted_at IS NULL AND p.deleted_at IS NULL", organizationID, groupID).
		Order("p.created_at DESC").
		Find(&policies).Error
	return policies, err
}

// ListGroupsAssignedToPolicy returns the groups a policy is attached to.
func (r *sentinelRepo) ListGroupsAssignedToPolicy(ctx context.Context, organizationID, policyID uuid.UUID) ([]*sentinelengine_models.Group, error) {
	var groups []*sentinelengine_models.Group
	err := r.conn(ctx).
		Table("sentinel_engine.groups AS g").
		Joins("JOIN sentinel_engine.groups_sentinel_policies AS gp ON gp.group_id = g.id").
		Where("gp.organization_id = ? AND gp.sentinel_policy_id = ? AND gp.deleted_at IS NULL AND g.deleted_at IS NULL", organizationID, policyID).
		Order("g.created_at DESC").
		Find(&groups).Error
	return groups, err
}

// IsPolicyAssignedToGroup reports whether an active (group, policy) edge exists.
func (r *sentinelRepo) IsPolicyAssignedToGroup(ctx context.Context, groupID, policyID uuid.UUID) (bool, error) {
	var count int64
	err := r.conn(ctx).
		Model(&sentinelengine_models.GroupsSentinelPolicies{}).
		Where("group_id = ? AND sentinel_policy_id = ?", groupID, policyID).
		Count(&count).Error
	return count > 0, err
}

// ReplacePoliciesAssignedToGroup sets a group's attached policies to exactly
// policyIDs: detach all current edges, then attach the given set, in one tx.
func (r *sentinelRepo) ReplacePoliciesAssignedToGroup(ctx context.Context, organizationID, groupID uuid.UUID, policyIDs []uuid.UUID) error {
	return r.conn(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Where("organization_id = ? AND group_id = ?", organizationID, groupID).
			Delete(&sentinelengine_models.GroupsSentinelPolicies{}).Error; err != nil {
			return err
		}
		for _, policyID := range policyIDs {
			edge := &sentinelengine_models.GroupsSentinelPolicies{
				ID:               uuid.New(),
				OrganizationID:   organizationID,
				GroupID:          groupID,
				SentinelPolicyID: policyID,
			}
			if err := tx.Create(edge).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
