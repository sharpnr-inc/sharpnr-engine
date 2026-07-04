package sentinel_repo

import (
	"context"

	"github.com/google/uuid"

	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinelengine_models "sharpnr-engine/sentinel/models"
)

// SentinelRepositoryInterface is the aggregate repository for the Sentinel
// authorization subsystem.
type SentinelRepositoryInterface interface {
	RolesRepository
	PolicyRepository
	RolePoliciesRepository
	GroupRepository
	GroupPoliciesRepository
	UserRoleRepository
	UserGroupRepository
	UserPolicyRepository
}

/* -------------------------------------------------------------------------- */
/*                                    Roles                                   */
/* -------------------------------------------------------------------------- */

type RolesRepository interface {
	// Create
	CreateRole(ctx context.Context, role *sentinelengine_models.Role) error

	// Read
	GetRoleByID(ctx context.Context, roleID uuid.UUID) (sentinelengine_models.Role, error)
	GetRoleByName(ctx context.Context, organizationID uuid.UUID, roleName string) (sentinelengine_models.Role, error)
	ListRolesByOrganizationID(ctx context.Context, organizationID uuid.UUID) ([]*sentinelengine_models.Role, error)

	// Update
	UpdateRoleDetails(
		ctx context.Context,
		roleID uuid.UUID,
		updates *sentinelDtos.UpdateRoleDetailsRequestDto,
	) error

	// Delete
	DeleteRole(ctx context.Context, roleID uuid.UUID) error

	// Helpers
	RoleExists(ctx context.Context, organizationID uuid.UUID, roleName string) (bool, error)
}

/* -------------------------------------------------------------------------- */
/*                                   Policies                                 */
/* -------------------------------------------------------------------------- */

type PolicyRepository interface {
	// Create
	CreatePolicy(ctx context.Context, policy *sentinelengine_models.SentinelPolicy) error

	// Read
	GetPolicyByID(ctx context.Context, policyID uuid.UUID) (sentinelengine_models.SentinelPolicy, error)

	// will fetch the all policies that are scoped to the organization,
	GetPolciesScopedToOrganization(ctx context.Context, organizationID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error)

	GetPolicyByName(ctx context.Context, organizationID uuid.UUID, policyName string) (sentinelengine_models.SentinelPolicy, error)

	ListPoliciesByOrganizationID(
		ctx context.Context,
		organizationID uuid.UUID,
		scope sentinelengine_models.PolicyScope,
	) ([]*sentinelengine_models.SentinelPolicy, error)

	// Update
	UpdatePolicy(ctx context.Context, policy *sentinelengine_models.SentinelPolicy) error
	EnablePolicy(ctx context.Context, policyID uuid.UUID) error
	DisablePolicy(ctx context.Context, policyID uuid.UUID) error

	// Delete
	DeletePolicy(ctx context.Context, policyID uuid.UUID) error

	// Helpers
	PolicyExists(ctx context.Context, organizationID uuid.UUID, policyName string) (bool, error)
}

/* -------------------------------------------------------------------------- */
/*                              Role ↔ Policies                               */
/* -------------------------------------------------------------------------- */

type RolePoliciesRepository interface {
	// Assignment
	AssignPolicyToRole(ctx context.Context, rolePolicy *sentinelengine_models.RolePolicies) error
	RemovePolicyFromRole(ctx context.Context, organizationID, roleID, policyID uuid.UUID) error

	// Queries
	ListPoliciesAssignedToRole(ctx context.Context, organizationID, roleID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error)
	ListRolesAssignedToPolicy(ctx context.Context, organizationID, policyID uuid.UUID) ([]*sentinelengine_models.Role, error)

	IsPolicyAssignedToRole(ctx context.Context, roleID, policyID uuid.UUID) (bool, error)

	// Bulk
	ReplacePoliciesAssignedToRole(ctx context.Context, organizationID, roleID uuid.UUID, policyIDs []uuid.UUID) error
}

/* -------------------------------------------------------------------------- */
/*                                   Groups                                   */
/* -------------------------------------------------------------------------- */

type GroupRepository interface {
	// Create
	CreateGroup(ctx context.Context, group *sentinelengine_models.Group) error

	// Read
	GetGroupByID(ctx context.Context, groupID uuid.UUID) (sentinelengine_models.Group, error)

	GetGroupByName(ctx context.Context, organizationID uuid.UUID, name string) (sentinelengine_models.Group, error)

	ListGroupsByOrganizationID(ctx context.Context, organizationID uuid.UUID) ([]*sentinelengine_models.Group, error)
	ListChildGroups(ctx context.Context, parentGroupID uuid.UUID) ([]*sentinelengine_models.Group, error)

	// Update
	UpdateGroup(ctx context.Context, group *sentinelengine_models.Group) error

	// Delete
	DeleteGroup(ctx context.Context, groupID uuid.UUID) error

	// Helpers
	GroupExists(ctx context.Context, organizationID uuid.UUID, name string) (bool, error)
}

/* -------------------------------------------------------------------------- */
/*                              Group ↔ Policies                              */
/* -------------------------------------------------------------------------- */

type GroupPoliciesRepository interface {
	// Assignment
	AssignPolicyToGroup(ctx context.Context, groupPolicy *sentinelengine_models.GroupsSentinelPolicies) error
	RemovePolicyFromGroup(ctx context.Context, organizationID, groupID, policyID uuid.UUID) error

	// Queries
	ListPoliciesAssignedToGroup(ctx context.Context, organizationID, groupID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error)
	ListGroupsAssignedToPolicy(ctx context.Context, organizationID, policyID uuid.UUID) ([]*sentinelengine_models.Group, error)

	IsPolicyAssignedToGroup(ctx context.Context, groupID, policyID uuid.UUID) (bool, error)

	// Bulk
	ReplacePoliciesAssignedToGroup(ctx context.Context, organizationID, groupID uuid.UUID, policyIDs []uuid.UUID) error
}

/* -------------------------------------------------------------------------- */
/*                                 User ↔ Roles                               */
/* -------------------------------------------------------------------------- */

type UserRoleRepository interface {
	// Assignment
	AssignRoleToUser(ctx context.Context, organizationID, userID, roleID uuid.UUID) error
	RemoveRoleFromUser(ctx context.Context, organizationID, userID, roleID uuid.UUID) error

	// Queries
	GetRolesAssignedToUser(ctx context.Context, userID uuid.UUID) ([]sentinelengine_models.UserRole, error)
	GetUsersAssignedToRole(ctx context.Context, organizationID, roleID uuid.UUID) ([]sentinelengine_models.UserRole, error)

	IsRoleAssignedToUser(ctx context.Context, userID, roleID uuid.UUID) (bool, error)

	// Bulk
	ReplaceRolesAssignedToUser(ctx context.Context, organizationID, userID uuid.UUID, roleIDs []uuid.UUID) error
}

/* -------------------------------------------------------------------------- */
/*                                User ↔ Groups                               */
/* -------------------------------------------------------------------------- */

type UserGroupRepository interface {
	// Assignment
	AssignUserToGroup(ctx context.Context, organizationID, userID, groupID uuid.UUID) error
	RemoveUserFromGroup(ctx context.Context, organizationID, userID, groupID uuid.UUID) error

	// Queries
	GetGroupsAssignedToUser(ctx context.Context, userID uuid.UUID) ([]sentinelengine_models.UserGroup, error)
	GetUsersAssignedToGroup(ctx context.Context, organizationID, groupID uuid.UUID) ([]sentinelengine_models.UserGroup, error)

	IsUserAssignedToGroup(ctx context.Context, userID, groupID uuid.UUID) (bool, error)

	// Bulk
	ReplaceGroupsAssignedToUser(ctx context.Context, organizationID, userID uuid.UUID, groupIDs []uuid.UUID) error
}

/* -------------------------------------------------------------------------- */
/*                               User ↔ Policies                              */
/* -------------------------------------------------------------------------- */

type UserPolicyRepository interface {
	// Assignment
	AssignPolicyToUser(ctx context.Context, organizationID, userID, policyID uuid.UUID) error
	RemovePolicyFromUser(ctx context.Context, organizationID, userID, policyID uuid.UUID) error

	// Queries
	GetPoliciesAssignedToUser(ctx context.Context, userID uuid.UUID) ([]*sentinelengine_models.UserPolicy, error)
	GetUsersAssignedToPolicy(ctx context.Context, organizationID, policyID uuid.UUID) ([]*sentinelengine_models.UserPolicy, error)

	IsPolicyAssignedToUser(ctx context.Context, userID, policyID uuid.UUID) (bool, error)

	// Bulk
	ReplacePoliciesAssignedToUser(ctx context.Context, organizationID, userID uuid.UUID, policyIDs []uuid.UUID) error
}
