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
	OrganizationPolicesRepoInterface
	RolePoliciesRepository
	GroupRepository
	GroupPoliciesRepository
	UserRoleRepository
	UserGroupRepository
	UserPolicyRepository
}

/* -------------------------------------------------------------------------- */
/*                                   Policies                                 */
/* -------------------------------------------------------------------------- */

type PolicyRepository interface {
	// Create
	CreatePolicy(ctx context.Context, policy *sentinelengine_models.SentinelPolicy) error
	// Read
	GetPolicyByID(ctx context.Context, policyID uuid.UUID) (sentinelengine_models.SentinelPolicy, error)
	ListPoliciesByOrganizationID(ctx context.Context, organizationID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error)

	// Update
	UpdatePolicy(ctx context.Context, policy *sentinelengine_models.SentinelPolicy) error
	EnablePolicy(ctx context.Context, policyID uuid.UUID) error
	DisablePolicy(ctx context.Context, policyID uuid.UUID) error

	// Delete
	DeletePolicy(ctx context.Context, policyID uuid.UUID) error
}

/* -------------------------------------------------------------------------- */
/*                            Organization Policies                           */
/* -------------------------------------------------------------------------- */

// this is Repo Interface for the Organization Repository
// all policies is stored in the sentinel Policy Table.
// this just Register which policy belongs to which org on Top Level Wise.
type OrganizationPolicesRepoInterface interface {
	CreateOrgWidePolicy(ctx context.Context, orgPolicy sentinelengine_models.OrgSentinelPolicy) error

	// this will returns the array or slice on in which we can find the SentinelPolicyID
	// which we can use the fetch the Policy from SentielPolicy Table.
	GetAllOrgWidePolicesEntry(ctx context.Context, orgID uuid.UUID) ([]sentinelengine_models.OrgSentinelPolicy, error)
	DeleteOrgPolicy(ctx context.Context, orgID uuid.UUID, orgPolicyID uuid.UUID) error
}

/* -------------------------------------------------------------------------- */
/*                                    Roles                                   */
/* -------------------------------------------------------------------------- */

type RolesRepository interface {

	// Create
	CreateRole(ctx context.Context, role *sentinelengine_models.Role) error

	// Read
	GetRoleByID(ctx context.Context, roleID uuid.UUID) (sentinelengine_models.Role, error)
	ListRolesByOrganizationID(ctx context.Context, organizationID uuid.UUID) ([]*sentinelengine_models.Role, error)
	UpdateRoleDetails(ctx context.Context, roleID uuid.UUID, updates *sentinelDtos.UpdateRoleDetailsRequestDto) error
	DeleteRole(ctx context.Context, roleID uuid.UUID) error

	// Helpers
	RoleExists(ctx context.Context, organizationID uuid.UUID, roleName string) (bool, error)
}

/* -------------------------------------------------------------------------- */
/*                              Role ↔ Policies                               */
/* -------------------------------------------------------------------------- */

type RolePoliciesRepository interface {
}

/* -------------------------------------------------------------------------- */
/*                                   Groups                                   */
/* -------------------------------------------------------------------------- */

type GroupRepository interface {
}

/* -------------------------------------------------------------------------- */
/*                              Group ↔ Policies                              */
/* -------------------------------------------------------------------------- */

type GroupPoliciesRepository interface {
}

/* -------------------------------------------------------------------------- */
/*                                 User ↔ Roles                               */
/* -------------------------------------------------------------------------- */

type UserRoleRepository interface {
}

/* -------------------------------------------------------------------------- */
/*                                User ↔ Groups                               */
/* -------------------------------------------------------------------------- */

type UserGroupRepository interface {
}

/* -------------------------------------------------------------------------- */
/*                               User ↔ Policies                              */
/* -------------------------------------------------------------------------- */

type UserPolicyRepository interface {
}
