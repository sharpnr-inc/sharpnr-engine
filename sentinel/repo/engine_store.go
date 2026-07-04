package sentinel_repo

import (
	"context"

	sentinelengine_models "sharpnr-engine/sentinel/models"
	sentinelEngine "sharpnr-engine/sentinel_engine"

	"github.com/google/uuid"
)

// The Sentinel repo doubles as the engine's read surface: the engine depends
// only on sentinelEngine.PolicyStore, and this file satisfies it.
var _ sentinelEngine.PolicyStore = (*sentinelRepo)(nil)

// ListEnabledPolicies implements [sentinelEngine.PolicyStore]: every enabled,
// live policy of an org, in the text form the engine compiles.
func (r *sentinelRepo) ListEnabledPolicies(ctx context.Context, orgID uuid.UUID) ([]sentinelEngine.PolicyRecord, error) {
	var rows []*sentinelengine_models.SentinelPolicy
	err := r.conn(ctx).
		Select("id", "policy").
		Where("organization_id = ? AND enabled = ?", orgID, true).
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

// GetPrincipalParents implements [sentinelEngine.PolicyStore]: the role and
// group names a user belongs to within an org (the Cedar parent edges).
func (r *sentinelRepo) GetPrincipalParents(ctx context.Context, orgID, userID uuid.UUID) (sentinelEngine.PrincipalParents, error) {
	var parents sentinelEngine.PrincipalParents

	err := r.conn(ctx).
		Table("sentinel_engine.roles AS ro").
		Joins("JOIN sentinel_engine.user_roles AS ur ON ur.role_id = ro.id").
		Where("ur.organization_id = ? AND ur.user_id = ? AND ro.deleted_at IS NULL", orgID, userID).
		Pluck("ro.name", &parents.Roles).Error
	if err != nil {
		return parents, err
	}

	err = r.conn(ctx).
		Table("sentinel_engine.groups AS g").
		Joins("JOIN sentinel_engine.user_groups AS ug ON ug.group_id = g.id").
		Where("ug.organization_id = ? AND ug.user_id = ? AND g.deleted_at IS NULL", orgID, userID).
		Pluck("g.name", &parents.Groups).Error
	return parents, err
}
