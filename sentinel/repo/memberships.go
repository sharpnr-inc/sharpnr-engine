package sentinel_repo

import (
	"context"
	"fmt"

	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
)

// The membership registers: which containers a user belongs to.
//
// These are the OTHER half of policy distribution. The attachment registers
// (attachments.go) say "this policy reaches role X"; these say "this user is in
// role X". The engine composes the two — see engine_store.go.
//
// Two differences from the attachment registers, both structural:
//
//   - No sentinel_policy_id. A membership never mentions a policy.
//   - HARD delete. user_roles and user_groups carry no deleted_at column, so
//     removing someone deletes the row outright — there is no detach/re-attach
//     history the way there is for attachments.
//
// ⚠ Neither table has a uniqueness constraint. 000027 created them with a
// composite PK (user_id, role_id), but 000038 dropped it to add a surrogate id
// and never added a replacement index. Assign therefore guards with a
// WHERE NOT EXISTS rather than ON CONFLICT — there is nothing to conflict on.
// That narrows the duplicate window to a single statement but does not close
// it; only a unique index on (organization_id, user_id, <container>) does.

// membershipTable is the physical register behind one container kind, plus the
// entity table its ids point at so reads can resolve to entities in one query.
type membershipTable struct {
	name         string // sentinel_engine.user_roles
	containerCol string // role_id
	entityTable  string // sentinel_engine.roles
}

var (
	roleMemberships  = membershipTable{"sentinel_engine.user_roles", "role_id", "sentinel_engine.roles"}
	groupMemberships = membershipTable{"sentinel_engine.user_groups", "group_id", "sentinel_engine.groups"}
)

/* -------------------------------------------------------------------------- */
/*                            Shared implementation                           */
/* -------------------------------------------------------------------------- */

// assignMembership puts a user in a container, skipping the insert when the
// membership already exists. See the uniqueness caveat above: this is one
// statement, not a read followed by a write, but without a unique index two
// concurrent callers can still both insert.
//
// assigned_by is left NULL. Nothing currently plumbs the acting employee down
// to the repo; when role assignment gets a real endpoint it should.
func (r *sentinelRepo) assignMembership(ctx context.Context, m membershipTable, orgID, userID, containerID uuid.UUID) error {
	stmt := fmt.Sprintf(`
		INSERT INTO %s (id, organization_id, user_id, %s, created_at)
		SELECT ?, ?, ?, ?, NOW()
		WHERE NOT EXISTS (
			SELECT 1 FROM %s
			WHERE organization_id = ? AND user_id = ? AND %s = ?
		)`, m.name, m.containerCol, m.name, m.containerCol)

	return r.conn(ctx).Exec(stmt,
		uuid.New(), orgID, userID, containerID, // the row
		orgID, userID, containerID, // the guard
	).Error
}

/* ---------------------------------- <***> --------------------------------- */

// removeMembership deletes the membership row. Hard delete — these tables have
// no deleted_at, so the same pair can simply be assigned again later.
func (r *sentinelRepo) removeMembership(ctx context.Context, m membershipTable, orgID, userID, containerID uuid.UUID) error {
	stmt := fmt.Sprintf(
		`DELETE FROM %s WHERE organization_id = ? AND user_id = ? AND %s = ?`,
		m.name, m.containerCol,
	)
	return r.conn(ctx).Exec(stmt, orgID, userID, containerID).Error
}

/* ---------------------------------- <***> --------------------------------- */

// isMember reports whether the user belongs to the container.
func (r *sentinelRepo) isMember(ctx context.Context, m membershipTable, orgID, userID, containerID uuid.UUID) (bool, error) {
	var count int64
	err := r.conn(ctx).Table(m.name).
		Where("organization_id = ? AND user_id = ?", orgID, userID).
		Where(m.containerCol+" = ?", containerID).
		Count(&count).Error
	return count > 0, err
}

/* ---------------------------------- <***> --------------------------------- */

// usersOfContainer returns the employee ids in a container.
func (r *sentinelRepo) usersOfContainer(ctx context.Context, m membershipTable, orgID, containerID uuid.UUID) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	err := r.conn(ctx).Table(m.name).
		Where("organization_id = ?", orgID).
		Where(m.containerCol+" = ?", containerID).
		Pluck("user_id", &ids).Error
	return ids, err
}

/* ---------------------------------- <***> --------------------------------- */

// membershipEntities resolves a user's memberships to the container ENTITIES in
// one JOIN, rather than returning raw edge rows for the caller to re-fetch.
// That is what keeps GetEmployeeRole / GetEmployeeGroups from doing N+1 lookups.
//
// Soft-deleted containers are excluded: a user left pointing at a deleted role
// is not in it, and must not receive its policies.
//
// A free function rather than a method because Go methods cannot take type
// parameters.
func membershipEntities[T any](ctx context.Context, r *sentinelRepo, m membershipTable, orgID, userID uuid.UUID) ([]*T, error) {
	var out []*T
	err := r.conn(ctx).
		Table(m.entityTable+" AS c").
		Select("c.*").
		Joins(fmt.Sprintf("JOIN %s AS mem ON mem.%s = c.id", m.name, m.containerCol)).
		Where("mem.organization_id = ? AND mem.user_id = ?", orgID, userID).
		Where("c.organization_id = ? AND c.deleted_at IS NULL", orgID).
		Order("c.created_at DESC").
		Find(&out).Error
	return out, err
}

/* -------------------------------------------------------------------------- */
/*                                User ↔ Roles                                */
/* -------------------------------------------------------------------------- */

func (r *sentinelRepo) AssignRoleToUser(ctx context.Context, orgID, userID, roleID uuid.UUID) error {
	return r.assignMembership(ctx, roleMemberships, orgID, userID, roleID)
}

func (r *sentinelRepo) RemoveRoleFromUser(ctx context.Context, orgID, userID, roleID uuid.UUID) error {
	return r.removeMembership(ctx, roleMemberships, orgID, userID, roleID)
}

func (r *sentinelRepo) IsRoleAssignedToUser(ctx context.Context, orgID, userID, roleID uuid.UUID) (bool, error) {
	return r.isMember(ctx, roleMemberships, orgID, userID, roleID)
}

// GetRolesOfUser returns the roles a user holds, as entities.
func (r *sentinelRepo) GetRolesOfUser(ctx context.Context, orgID, userID uuid.UUID) ([]*sentinelengine_models.Role, error) {
	return membershipEntities[sentinelengine_models.Role](ctx, r, roleMemberships, orgID, userID)
}

// GetUsersOfRole returns the employee ids holding a role.
func (r *sentinelRepo) GetUsersOfRole(ctx context.Context, orgID, roleID uuid.UUID) ([]uuid.UUID, error) {
	return r.usersOfContainer(ctx, roleMemberships, orgID, roleID)
}

/* -------------------------------------------------------------------------- */
/*                               User ↔ Groups                                */
/* -------------------------------------------------------------------------- */

func (r *sentinelRepo) AssignUserToGroup(ctx context.Context, orgID, userID, groupID uuid.UUID) error {
	return r.assignMembership(ctx, groupMemberships, orgID, userID, groupID)
}

func (r *sentinelRepo) RemoveUserFromGroup(ctx context.Context, orgID, userID, groupID uuid.UUID) error {
	return r.removeMembership(ctx, groupMemberships, orgID, userID, groupID)
}

func (r *sentinelRepo) IsUserInGroup(ctx context.Context, orgID, userID, groupID uuid.UUID) (bool, error) {
	return r.isMember(ctx, groupMemberships, orgID, userID, groupID)
}

// GetGroupsOfUser returns the groups a user belongs to, as entities.
//
// Direct memberships only — group nesting is not walked, so belonging to a
// child group does not imply membership of its parent.
func (r *sentinelRepo) GetGroupsOfUser(ctx context.Context, orgID, userID uuid.UUID) ([]*sentinelengine_models.Group, error) {
	return membershipEntities[sentinelengine_models.Group](ctx, r, groupMemberships, orgID, userID)
}

// GetUsersOfGroup returns the employee ids in a group.
func (r *sentinelRepo) GetUsersOfGroup(ctx context.Context, orgID, groupID uuid.UUID) ([]uuid.UUID, error) {
	return r.usersOfContainer(ctx, groupMemberships, orgID, groupID)
}
