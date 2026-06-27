package sentinel_repo

import (
	"context"
	sentinelengine_models "sharpnr-engine/sentinel/models"

	"github.com/google/uuid"
)

// SentinelRepository is the data-access layer for the authorization tables.
// Each method participates in an ambient transaction when one is present in the
// context (see dbtx), so a service can run several writes atomically.
type SentinelRepositoryInterface interface {
	RolesRepository
	PolicyRepository
	GroupRepository
	RolePoliciesRepository

	AssignUserRole(ctx context.Context, orgID, userID, roleID uuid.UUID) error
	GetUserRole(ctx context.Context, orgID, emplID uuid.UUID) (sentinelengine_models.UserRole, error) // role-assignment edge for an employee
}

/* ---------------------------------- <***> --------------------------------- */

type RolesRepository interface {
	CreateRole(ctx context.Context, role *sentinelengine_models.Role) error
	GetRoleByID(ctx context.Context, id uuid.UUID) (sentinelengine_models.Role, error)
	ListRolesByOrgID(ctx context.Context, orgID uuid.UUID) ([]*sentinelengine_models.Role, error)
}

/* ---------------------------------- <***> --------------------------------- */

type PolicyRepository interface {
	CreatePolicy(ctx context.Context, policy *sentinelengine_models.SentinelPolicy) error
	GetPolicyByID(ctx context.Context, id uuid.UUID) (sentinelengine_models.SentinelPolicy, error)
	// ListPoliciesByOrgID returns an org's policies; an empty scope means all
	// scopes, otherwise only policies in that bucket.
	ListPoliciesByOrgID(ctx context.Context, orgID uuid.UUID, scope sentinelengine_models.PolicyScope) ([]*sentinelengine_models.SentinelPolicy, error)
}

/* ---------------------------------- <***> --------------------------------- */

type GroupRepository interface {
	CreateGroup(ctx context.Context, group *sentinelengine_models.Group) error
	ListGroupsByOrgID(ctx context.Context, orgID uuid.UUID) ([]*sentinelengine_models.Group, error)
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
