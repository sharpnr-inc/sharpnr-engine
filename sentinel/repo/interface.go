package sentinel_repo

import (
	"context"
	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
)

// SentinelRepository is the data-access layer for the authorization tables.
// Each method participates in an ambient transaction when one is present in the
// context (see dbtx), so a service can run several writes atomically.
type SentinelRepositoryInterface interface {
	RolesRepository
	PolicyRepository
	RolePoliciesRepository
	GroupRepository
	UserRoleRepository
	UserGroupRepository
	UserPolicyRepository
}

/* ---------------------------------- <***> --------------------------------- */

type RolesRepository interface {

	// Create the Roles in The DB.
	CreateRole(ctx context.Context, role *sentinelengine_models.Role) error
	// Get Role By Role.ID returns the Role
	GetRoleByID(ctx context.Context, id uuid.UUID) (sentinelengine_models.Role, error)
	// this will list all the Roles in the organization
	ListRolesByOrgID(ctx context.Context, orgID uuid.UUID) ([]*sentinelengine_models.Role, error)
	// this will update the Role General Details
	UpdateRoleDetails(ctx context.Context, roleID uuid.UUID, updates *sentinelDtos.UpdateRoleDetailsRequestDto) error
}

/* ---------------------------------- <***> --------------------------------- */

// RolePoliciesRepository manages the role<->policy attachment edges. The engine
// compiles every enabled org policy regardless of attachment; this join is the
// organizational/UI record of which policies an admin grouped under a role.
type RolePoliciesRepository interface {
	CreateRolePolicy(ctx context.Context, rp *sentinelengine_models.RolePolicies) error
	ListPoliciesByRoleID(ctx context.Context, orgID, roleID uuid.UUID) ([]*sentinelengine_models.SentinelPolicy, error)
	DeleteRolePolicy(ctx context.Context, orgID, roleID, policyID uuid.UUID) error
}

/* ---------------------------------- <***> --------------------------------- */

type PolicyRepository interface {
	// this will create Policy it create one Policy
	CreatePolicy(ctx context.Context, policy *sentinelengine_models.SentinelPolicy) error
	// This will return a Policy by the Policy.ID
	GetPolicyByID(ctx context.Context, id uuid.UUID) (sentinelengine_models.SentinelPolicy, error)
	// ListPoliciesByOrgID returns an org's policies; an empty scope means all
	// scopes, otherwise only policies in that bucket.
	ListPoliciesByOrgID(ctx context.Context, orgID uuid.UUID, scope sentinelengine_models.PolicyScope) ([]*sentinelengine_models.SentinelPolicy, error)
}

/* ---------------------------------- <***> --------------------------------- */

type GroupRepository interface {
	// This create a Group in the Organzitaion
	CreateGroup(ctx context.Context, group *sentinelengine_models.Group) error
	// List all Groups of the organization by the orgID
	ListGroupsByOrgID(ctx context.Context, orgID uuid.UUID) ([]*sentinelengine_models.Group, error)
}

type UserRoleRepository interface {
	// this will assign the Role to the Employee
	AssignUserRole(ctx context.Context, orgID, emplID, roleID uuid.UUID) error
	GetUserRole(ctx context.Context, orgID, emplID uuid.UUID) (sentinelengine_models.UserRole, error) // role-assignment edge for an employee

	CreateUserRole(ctx context.Context, userRole *sentinelengine_models.UserRole) error
}

type UserGroupRepository interface {
}

type UserPolicyRepository interface {
}

/* ---------------------------------- <***> --------------------------------- */
