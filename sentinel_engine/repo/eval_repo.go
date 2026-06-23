package sentinel_repo

import (
	"context"
	"errors"

	sentinelengine "sharpnr-engine/sentinel_engine/engine"
	sentinelengine_models "sharpnr-engine/sentinel_engine/models"

	"github.com/cedar-policy/cedar-go/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// This file implements the read side the Sentinel engine needs to make a
// decision: the enabled policies for an org, plus the principal's role and
// group membership. Together these satisfy sentinelengine.Repository, so the
// concrete *sentinelRepo can be passed straight into sentinelengine.New.
//
// Writes live in policy_repo.go / roles_repo.go; this file is read-only and is
// hit on every authorization call (the engine rebuilds the entity graph each
// time), so the queries are kept tight and indexed.

// ListEnabledPolicies returns every enabled policy row for an org. The engine
// compiles these into a per-org cedar.PolicySet (cached until InvalidateOrg).
func (r *sentinelRepo) ListEnabledPolicies(ctx context.Context, orgID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
	var rows []*sentinelengine_models.SentinelPolicy
	err := r.conn(ctx).
		Where("organization_id = ? AND enabled = ?", orgID, true).
		Find(&rows).Error
	return rows, err
}

// UserRoleSlugs returns the slugs of the roles directly assigned to the user
// within the org. Slugs (not ids) because policies reference Role::"<slug>".
func (r *sentinelRepo) UserRoleSlugs(ctx context.Context, orgID, userID uuid.UUID) ([]string, error) {
	var slugs []string
	err := r.conn(ctx).
		Table("sentinel_engine.user_roles AS ur").
		Joins("JOIN sentinel_engine.roles AS rl ON rl.id = ur.role_id").
		Where("ur.organization_id = ? AND ur.user_id = ?", orgID, userID).
		Pluck("rl.slug", &slugs).Error
	return slugs, err
}

// UserGroupClosure returns the user's direct groups (Direct=true) plus every
// ancestor group reached via parent_group_id (Direct=false). Each row also
// carries its own parent's slug so the engine can wire the Group hierarchy
// edges. A single recursive CTE walks the closure; bool_or collapses a group
// that is both a direct membership and an ancestor into Direct=true.
func (r *sentinelRepo) UserGroupClosure(ctx context.Context, orgID, userID uuid.UUID) ([]sentinelengine.GroupMembership, error) {
	const q = `
WITH RECURSIVE closure AS (
    SELECT g.id, g.slug, g.parent_group_id, TRUE AS direct
    FROM sentinel_engine.groups g
    JOIN sentinel_engine.user_groups ug ON ug.group_id = g.id
    WHERE ug.organization_id = ? AND ug.user_id = ?
  UNION
    SELECT pg.id, pg.slug, pg.parent_group_id, FALSE AS direct
    FROM sentinel_engine.groups pg
    JOIN closure c ON c.parent_group_id = pg.id
)
SELECT
    c.slug              AS slug,
    COALESCE(p.slug, '') AS parent_slug,
    bool_or(c.direct)   AS direct
FROM closure c
LEFT JOIN sentinel_engine.groups p ON p.id = c.parent_group_id
GROUP BY c.slug, p.slug;`

	var out []sentinelengine.GroupMembership
	err := r.conn(ctx).Raw(q, orgID, userID).Scan(&out).Error
	return out, err
}

// UserAttributes returns the principal's ABAC attributes for the engine to stamp
// onto the User entity, so policies can use `principal.<attr>` in `when` clauses.
// Sourced from the employee's organization profile (department, position,
// managerId). A missing profile is not an error — it just means no extra
// attributes. Queried by table name to avoid coupling the sentinel repo to the
// employee model package.
func (r *sentinelRepo) UserAttributes(ctx context.Context, orgID, userID uuid.UUID) (map[string]types.Value, error) {
	var row struct {
		Department string
		Position   string
		ManagerID  string
	}
	err := r.conn(ctx).
		Table("employee.employee_organization_profile").
		Select("department, position, manager_id").
		Where("organization_id = ? AND employee_id = ?", orgID, userID).
		Take(&row).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return map[string]types.Value{}, nil
	}
	if err != nil {
		return nil, err
	}

	attrs := map[string]types.Value{}
	if row.Department != "" {
		attrs["department"] = types.String(row.Department)
	}
	if row.Position != "" {
		attrs["position"] = types.String(row.Position)
	}
	if row.ManagerID != "" {
		attrs["managerId"] = types.String(row.ManagerID)
	}
	return attrs, nil
}
