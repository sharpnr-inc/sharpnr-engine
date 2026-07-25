package sentinel_repo

import (
	"context"

	sentinelEngine "sharpnr-engine/sentinel_engine"

	"github.com/google/uuid"
)

// The Sentinel repo doubles as the engine's read surface: the engine depends
// only on sentinelEngine.PolicyStore, and this file satisfies it.
var _ sentinelEngine.PolicyStore = (*sentinelRepo)(nil)

// GetAllPoliciesOfUser implements [sentinelEngine.PolicyStore]: every enabled,
// live policy that applies to the user, collected from all four sources —
//
//  1. policies scoped to the organization (scope = 'organization'; the only
//     scope with engine meaning — role/group/user scopes are FE grouping labels)
//  2. policies attached to the user's roles  (user_roles -> role_policies)
//  3. policies attached to the user's groups (user_groups -> groups_sentinel_policies)
//  4. policies attached directly to the user (user_policies)
//
// One SELECT with OR-EXISTS branches instead of four queries: the DB does the
// combining, and a policy reachable through several attachments (e.g. two of
// the user's roles) still comes back exactly once.
func (r *sentinelRepo) GetAllPoliciesOfUser(ctx context.Context, orgID, userID uuid.UUID) ([]sentinelEngine.PolicyRecord, error) {
	var rows []struct {
		ID     uuid.UUID
		Policy string
	}

	err := r.conn(ctx).
		Table("sentinel_engine.sentinel_policies AS p").
		Select("p.id", "p.policy").
		Where("p.organization_id = ? AND p.enabled = TRUE AND p.deleted_at IS NULL", orgID).
		Where(`(
			p.scope = 'organization'
			OR EXISTS (
				SELECT 1 FROM sentinel_engine.role_policies rp
				JOIN sentinel_engine.user_roles ur ON ur.role_id = rp.role_id
				WHERE rp.sentinel_policy_id = p.id
				  AND rp.deleted_at IS NULL
				  AND ur.organization_id = p.organization_id
				  AND ur.user_id = ?
			)
			OR EXISTS (
				SELECT 1 FROM sentinel_engine.groups_sentinel_policies gp
				JOIN sentinel_engine.user_groups ug ON ug.group_id = gp.group_id
				WHERE gp.sentinel_policy_id = p.id
				  AND gp.deleted_at IS NULL
				  AND ug.organization_id = p.organization_id
				  AND ug.user_id = ?
			)
			OR EXISTS (
				SELECT 1 FROM sentinel_engine.user_policies up
				WHERE up.sentinel_policy_id = p.id
				  AND up.deleted_at IS NULL
				  AND up.organization_id = p.organization_id
				  AND up.employee_id = ?
			)
		)`, userID, userID, userID).
		Find(&rows).Error
	if err != nil {
		return nil, err
	}

	records := make([]sentinelEngine.PolicyRecord, 0, len(rows))
	for _, row := range rows {
		records = append(records, sentinelEngine.PolicyRecord{ID: row.ID, Text: row.Policy})
	}
	return records, nil
}
