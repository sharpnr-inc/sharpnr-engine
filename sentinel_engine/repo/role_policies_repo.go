package sentinel_repo

import (
	"context"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"

	"github.com/google/uuid"
)

// This file implements the role<->policy attachment edges. The engine compiles
// every enabled org policy regardless of attachment (see eval_repo.go), so these
// rows are the organizational record of which policies an admin grouped under a
// role — they drive the "role privileges" view in the UI.

// CreateRolePolicy persists one role<->policy attachment. The ID is
// application-generated; if unset we assign one here. Joins an ambient
// transaction when one is present in the context.
func (r *sentinelRepo) CreateRolePolicy(ctx context.Context, rp *sentinelengine_models.RolePolicies) error {
	if rp.ID == uuid.Nil {
		rp.ID = uuid.New()
	}
	return r.conn(ctx).Create(rp).Error
}

// ListPoliciesByRoleID returns the policies attached to a role within an org, by
// joining the role_policies edge to the policy rows. Used to render a role's
// privileges.
func (r *sentinelRepo) ListPoliciesByRoleID(ctx context.Context, orgID, roleID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
	var rows []*sentinelengine_models.SentinelPolicy
	err := r.conn(ctx).
		Table("sentinel_engine.sentinel_policies AS p").
		Joins("JOIN sentinel_engine.role_policies AS rp ON rp.sentinel_policy_id = p.id").
		Where("rp.organization_id = ? AND rp.role_id = ? AND rp.deleted_at IS NULL", orgID, roleID).
		Order("p.created_at DESC").
		Find(&rows).Error
	return rows, err
}

// DeleteRolePolicy removes one role<->policy attachment (soft delete via gorm).
// It does not delete the policy itself, only the grouping under the role.
func (r *sentinelRepo) DeleteRolePolicy(ctx context.Context, orgID, roleID, policyID uuid.UUID) error {
	return r.conn(ctx).
		Where("organization_id = ? AND role_id = ? AND sentinel_policy_id = ?", orgID, roleID, policyID).
		Delete(&sentinelengine_models.RolePolicies{}).Error
}
