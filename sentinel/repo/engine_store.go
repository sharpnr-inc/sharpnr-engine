package sentinel_repo

import (
	"context"
	"fmt"

	sentinelengine_models "sharpnr-engine/sentinel/models"
	sentinelEngine "sharpnr-engine/sentinel_engine"

	"github.com/google/uuid"
)

// The Sentinel repo doubles as the engine's read surface: the engine depends
// only on sentinelEngine.PolicyStore, and this file satisfies it.
var _ sentinelEngine.PolicyStore = (*sentinelRepo)(nil)

// policySource is one path by which a policy can reach a user. The four
// sources below are the whole of policy distribution — there is no fifth
// mechanism and no special case, which is what dropping `scope` bought.
type policySource struct {
	kind PrincipalKind
	load func(ctx context.Context, orgID, userID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error)
}

// policySources lists the paths in reach order, broadest first. Adding a
// principal kind (department, cost centre, ...) is one entry here.
func (r *sentinelRepo) policySources() []policySource {
	return []policySource{
		// Registered against the org itself: reaches every employee.
		{PrincipalOrg, func(ctx context.Context, orgID, _ uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			return r.policiesAttachedTo(ctx, PrincipalOrg, orgID, orgID)
		}},
		// Registered against a role the user holds.
		{PrincipalRole, func(ctx context.Context, orgID, userID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			return r.policiesViaMembership(ctx, roleMemberships, PrincipalRole, orgID, userID)
		}},
		// Registered against a group the user belongs to.
		{PrincipalGroup, func(ctx context.Context, orgID, userID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			return r.policiesViaMembership(ctx, groupMemberships, PrincipalGroup, orgID, userID)
		}},
		// Granted directly to the employee.
		{PrincipalUser, func(ctx context.Context, orgID, userID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error) {
			return r.policiesAttachedTo(ctx, PrincipalUser, orgID, userID)
		}},
	}
}

/* ---------------------------------- <***> --------------------------------- */

// GetAllPoliciesOfUser implements [sentinelEngine.PolicyStore]: every enabled,
// live policy that applies to the user, unioned across all four distribution
// paths.
//
// One query per source rather than one query with four OR-EXISTS branches. The
// sources are the same functions the admin screens use, so "which policies does
// this role grant" has exactly one implementation. The cost is four round trips
// instead of one, on a path the engine caches per (org, user) — see
// policyCacheTTL in sentinel_engine.
//
// The `seen` map does what OR-EXISTS did implicitly: a policy reachable by
// several paths (two of the user's roles, or a role AND directly) is compiled
// once.
//
// Fails closed by construction — any source erroring aborts the whole load, so
// the engine never evaluates against a partial policy set.
func (r *sentinelRepo) GetAllPoliciesOfUser(ctx context.Context, orgID, userID uuid.UUID) ([]sentinelEngine.PolicyRecord, error) {
	seen := make(map[uuid.UUID]struct{})
	records := make([]sentinelEngine.PolicyRecord, 0, 16)

	for _, src := range r.policySources() {
		policies, err := src.load(ctx, orgID, userID)
		if err != nil {
			return nil, fmt.Errorf("sentinel: loading %s policies for user %s: %w", src.kind, userID, err)
		}
		for _, p := range policies {
			// The registers deliberately return disabled policies — an admin
			// screen wants to see them. The engine must not compile them.
			if !p.Enabled {
				continue
			}
			if _, dup := seen[p.ID]; dup {
				continue
			}
			seen[p.ID] = struct{}{}
			records = append(records, sentinelEngine.PolicyRecord{ID: p.ID, Text: p.Policy})
		}
	}
	return records, nil
}

/* ---------------------------------- <***> --------------------------------- */

// policiesViaMembership resolves the two-hop paths: the policies registered
// against any container (role or group) the user belongs to.
//
//	sentinel_policies ⋈ <attachment register> ⋈ <membership register> ⋈ <container>
//
// The final join to the container table is what excludes SOFT-DELETED roles and
// groups. This is a deliberate tightening over the pre-rework query, which
// joined only the registers: deleting a role used to leave its policies
// reaching every member, because neither the membership row nor the
// role_policies row cascades from roles(id). Deleting a role now revokes what
// it granted, and it matches what GetRolesOfUser reports.
func (r *sentinelRepo) policiesViaMembership(
	ctx context.Context,
	m membershipTable,
	kind PrincipalKind,
	orgID, userID uuid.UUID,
) ([]*sentinelengine_models.SentinelPolicy, error) {
	att, err := kind.table()
	if err != nil {
		return nil, err
	}

	var policies []*sentinelengine_models.SentinelPolicy
	err = r.conn(ctx).
		Table("sentinel_engine.sentinel_policies AS p").
		Select("p.*").
		Joins(fmt.Sprintf(
			"JOIN %s AS att ON att.sentinel_policy_id = p.id AND att.deleted_at IS NULL", att.name)).
		Joins(fmt.Sprintf(
			"JOIN %s AS mem ON mem.%s = att.%s", m.name, m.containerCol, att.principalCol)).
		Joins(fmt.Sprintf(
			"JOIN %s AS c ON c.id = mem.%s AND c.deleted_at IS NULL", m.entityTable, m.containerCol)).
		Where("p.organization_id = ? AND p.deleted_at IS NULL", orgID).
		Where("att.organization_id = ?", orgID).
		Where("mem.organization_id = ? AND mem.user_id = ?", orgID, userID).
		Where("c.organization_id = ?", orgID).
		Order("p.created_at DESC").
		Find(&policies).Error
	return policies, err
}
