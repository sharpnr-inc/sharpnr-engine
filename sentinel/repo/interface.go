package sentinel_repo

import (
	"context"

	"github.com/google/uuid"

	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinelengine_models "sharpnr-engine/sentinel/models"
)

// SentinelRepositoryInterface is the aggregate repository for the Sentinel
// authorization subsystem.
//
// The schema has nine tables in three kinds, and this interface mirrors that
// exactly — one sub-interface per kind, not one per table:
//
//	entities     sentinel_policies, roles, groups
//	attachments  org_sentinel_policies, role_policies,
//	             groups_sentinel_policies, user_policies   (policy -> principal)
//	memberships  user_roles, user_groups                   (user -> container)
//
// Convention: every method takes organizationID and filters on it. Tenant
// scoping is part of the signature, not something the caller re-checks after a
// fetch.
type SentinelRepositoryInterface interface {
	PolicyRepository
	RoleRepository
	GroupRepository

	AttachmentRepository
	MembershipRepository
}

/* -------------------------------------------------------------------------- */
/*                            Entities: Policies                              */
/* -------------------------------------------------------------------------- */

// PolicyRepository owns sentinel_engine.sentinel_policies — the one table
// policies live in. It knows nothing about who a policy reaches; that is the
// attachment registers' job.
type PolicyRepository interface {
	// Create
	CreatePolicy(ctx context.Context, policy *sentinelengine_models.SentinelPolicy) error

	// Read
	GetPolicy(ctx context.Context, organizationID, policyID uuid.UUID) (sentinelengine_models.SentinelPolicy, error)
	GetPolicyByName(ctx context.Context, organizationID uuid.UUID, name string) (sentinelengine_models.SentinelPolicy, error)
	ListPolicies(ctx context.Context, organizationID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error)

	// Update
	UpdatePolicy(ctx context.Context, organizationID uuid.UUID, policy *sentinelengine_models.SentinelPolicy) error
	SetPolicyEnabled(ctx context.Context, organizationID, policyID uuid.UUID, enabled bool) error

	// Delete
	DeletePolicy(ctx context.Context, organizationID, policyID uuid.UUID) error

	// Helpers
	PolicyExists(ctx context.Context, organizationID uuid.UUID, name string) (bool, error)
}

/* -------------------------------------------------------------------------- */
/*                             Entities: Roles                                */
/* -------------------------------------------------------------------------- */

// RoleRepository owns sentinel_engine.roles. Names are unique per org
// (idx_role_org_name).
type RoleRepository interface {
	// Create
	CreateRole(ctx context.Context, role *sentinelengine_models.Role) error

	// Read
	GetRole(ctx context.Context, organizationID, roleID uuid.UUID) (sentinelengine_models.Role, error)
	GetRoleByName(ctx context.Context, organizationID uuid.UUID, name string) (sentinelengine_models.Role, error)
	ListRoles(ctx context.Context, organizationID uuid.UUID) ([]*sentinelengine_models.Role, error)

	// Update
	UpdateRole(ctx context.Context, organizationID, roleID uuid.UUID, updates *sentinelDtos.UpdateRoleDetailsRequestDto) error

	// Delete
	DeleteRole(ctx context.Context, organizationID, roleID uuid.UUID) error

	// Helpers
	RoleExists(ctx context.Context, organizationID uuid.UUID, name string) (bool, error)
}

/* -------------------------------------------------------------------------- */
/*                             Entities: Groups                               */
/* -------------------------------------------------------------------------- */

// GroupRepository owns sentinel_engine.groups. Unlike roles, group names carry
// no unique index — GroupExists is the check before creating.
type GroupRepository interface {
	// Create
	CreateGroup(ctx context.Context, group *sentinelengine_models.Group) error

	// Read
	GetGroup(ctx context.Context, organizationID, groupID uuid.UUID) (sentinelengine_models.Group, error)
	GetGroupByName(ctx context.Context, organizationID uuid.UUID, name string) (sentinelengine_models.Group, error)
	ListGroups(ctx context.Context, organizationID uuid.UUID) ([]*sentinelengine_models.Group, error)
	ListChildGroups(ctx context.Context, organizationID, parentGroupID uuid.UUID) ([]*sentinelengine_models.Group, error)

	// Update
	UpdateGroup(ctx context.Context, organizationID uuid.UUID, group *sentinelengine_models.Group) error

	// Delete
	DeleteGroup(ctx context.Context, organizationID, groupID uuid.UUID) error

	// Helpers
	GroupExists(ctx context.Context, organizationID uuid.UUID, name string) (bool, error)
}

/* -------------------------------------------------------------------------- */
/*                                Attachments                                 */
/* -------------------------------------------------------------------------- */

// AttachmentRepository owns the four registers that decide which policies reach
// which principals: org_sentinel_policies, role_policies,
// groups_sentinel_policies and user_policies.
//
// They are structurally identical, so one implementation serves all four
// (attachments.go). Attach is idempotent; Detach is a soft delete, so a
// registration can be lifted and restored with a history. Table names never
// leave the repo — callers speak in PrincipalKind.
type AttachmentRepository interface {
	// Organization-wide: reaches every employee. No principal id — the org is
	// the principal. Replaces the old scope = 'organization' flag.
	AttachPolicyToOrg(ctx context.Context, orgID, policyID uuid.UUID) error
	DetachPolicyFromOrg(ctx context.Context, orgID, policyID uuid.UUID) error
	IsPolicyAttachedToOrg(ctx context.Context, orgID, policyID uuid.UUID) (bool, error)
	ListPoliciesOfOrg(ctx context.Context, orgID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error)

	// Role
	AttachPolicyToRole(ctx context.Context, orgID, roleID, policyID uuid.UUID) error
	DetachPolicyFromRole(ctx context.Context, orgID, roleID, policyID uuid.UUID) error
	IsPolicyAttachedToRole(ctx context.Context, orgID, roleID, policyID uuid.UUID) (bool, error)
	ListPoliciesOfRole(ctx context.Context, orgID, roleID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error)

	// Group
	AttachPolicyToGroup(ctx context.Context, orgID, groupID, policyID uuid.UUID) error
	DetachPolicyFromGroup(ctx context.Context, orgID, groupID, policyID uuid.UUID) error
	IsPolicyAttachedToGroup(ctx context.Context, orgID, groupID, policyID uuid.UUID) (bool, error)
	ListPoliciesOfGroup(ctx context.Context, orgID, groupID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error)

	// User — a direct grant to one employee.
	AttachPolicyToUser(ctx context.Context, orgID, userID, policyID uuid.UUID) error
	DetachPolicyFromUser(ctx context.Context, orgID, userID, policyID uuid.UUID) error
	IsPolicyAttachedToUser(ctx context.Context, orgID, userID, policyID uuid.UUID) (bool, error)
	ListPoliciesOfUser(ctx context.Context, orgID, userID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error)

	// Reverse lookup: every principal a policy reaches, grouped by kind. This
	// is what the UI shows in place of the removed `scope` label.
	ListPolicyAttachments(ctx context.Context, orgID, policyID uuid.UUID) (map[PrincipalKind][]uuid.UUID, error)
}

/* -------------------------------------------------------------------------- */
/*                                Memberships                                 */
/* -------------------------------------------------------------------------- */

// MembershipRepository owns user_roles and user_groups — which containers a
// user belongs to. The attachment registers say which policies reach a
// container; these say who is in it.
//
// Unlike attachments these are HARD deletes (no deleted_at column), so removing
// a membership leaves no history. Reads resolve to entities in one JOIN rather
// than returning edge rows, and exclude soft-deleted containers.
type MembershipRepository interface {
	// User ↔ Roles
	AssignRoleToUser(ctx context.Context, orgID, userID, roleID uuid.UUID) error
	RemoveRoleFromUser(ctx context.Context, orgID, userID, roleID uuid.UUID) error
	IsRoleAssignedToUser(ctx context.Context, orgID, userID, roleID uuid.UUID) (bool, error)
	GetRolesOfUser(ctx context.Context, orgID, userID uuid.UUID) ([]*sentinelengine_models.Role, error)
	GetUsersOfRole(ctx context.Context, orgID, roleID uuid.UUID) ([]uuid.UUID, error)

	// User ↔ Groups
	AssignUserToGroup(ctx context.Context, orgID, userID, groupID uuid.UUID) error
	RemoveUserFromGroup(ctx context.Context, orgID, userID, groupID uuid.UUID) error
	IsUserInGroup(ctx context.Context, orgID, userID, groupID uuid.UUID) (bool, error)
	GetGroupsOfUser(ctx context.Context, orgID, userID uuid.UUID) ([]*sentinelengine_models.Group, error)
	GetUsersOfGroup(ctx context.Context, orgID, groupID uuid.UUID) ([]uuid.UUID, error)
}
