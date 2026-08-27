package sentinelservice

import (
	"context"

	sentinelDtos "sharpnr-engine/sentinel/dtos"

	"github.com/google/uuid"
)

// SentinelServiceInterface is the authorization AUTHORING service, injected into
// other modules and exposed to the frontend through the orchestrator. It owns
// provisioning, the policy builder's principal palette, and CRUD over policies,
// roles, groups and the registers that connect them.
//
// Evaluation lives elsewhere: sentinel_engine reads the same tables through its
// own one-method PolicyStore and never goes through this interface.
type SentinelServiceInterface interface {
	// ProvisionOrganization sets up a new org's authorization defaults: the
	// full-access "Owner" role + its policy, the founding employee's owner
	// assignment, and the default "Stakeholders" group. Runs in one transaction
	// so an org is never left half-provisioned.
	ProvisionOrganization(ctx context.Context, orgID, firstUserID uuid.UUID) error

	// GetEmployeeRole resolves the role an employee holds within an org;
	// GetEmployeeGroups the groups they belong to.
	GetEmployeeRole(ctx context.Context, orgID, emplID uuid.UUID) (sentinelDtos.RoleView, error)
	GetEmployeeGroups(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelDtos.GroupView, error)

	// GetPrincipals returns the policy builder's principal palette: the static
	// entity types plus the org's live roles and groups.
	GetPrincipals(ctx context.Context, orgID uuid.UUID) (sentinelDtos.PrincipalsView, error)

	PolicyService
	RoleService
	GroupService
}

// Compile-time guard: the concrete service must satisfy the interface.
var _ SentinelServiceInterface = (*sentinelService)(nil)

/* -------------------------------------------------------------------------- */
/*                                  Policies                                  */
/* -------------------------------------------------------------------------- */

// PolicyService owns the policy rows and the registrations that decide who they
// reach. A policy created here is attached to nothing and reaches nobody until
// it is registered against a principal.
type PolicyService interface {
	// CreatePolicy persists a raw Cedar policy (authored and validated in the
	// FE), revalidating that it compiles, and invalidates the org's engine cache.
	CreatePolicy(ctx context.Context, orgID uuid.UUID, in sentinelDtos.AddPolicyRequestDto) (uuid.UUID, error)

	// ListPolicies returns every policy in the org. There are no buckets — what
	// a policy reaches is read from the registers, not from a column on it.
	ListPolicies(ctx context.Context, orgID uuid.UUID) ([]sentinelDtos.PolicyView, error)
	GetPolicy(ctx context.Context, orgID, policyID uuid.UUID) (sentinelDtos.PolicyView, error)

	// AttachPolicies registers existing policies against a role; DetachPolicy
	// lifts one registration. Both leave the policy rows themselves untouched.
	AttachPolicies(ctx context.Context, orgID, emplID, roleID uuid.UUID, policyIDs []uuid.UUID) error
	DetachPolicy(ctx context.Context, orgID, roleID, policyID uuid.UUID) error
}

/* -------------------------------------------------------------------------- */
/*                                    Roles                                   */
/* -------------------------------------------------------------------------- */

type RoleService interface {
	CreateRole(ctx context.Context, orgID, emplID uuid.UUID, in sentinelDtos.CreateRoleInput) (uuid.UUID, error)
	ListRoles(ctx context.Context, orgID uuid.UUID) ([]sentinelDtos.RoleView, error)

	// GetRolePrivileges returns the role plus every policy registered against it.
	GetRolePrivileges(ctx context.Context, orgID, roleID uuid.UUID) (sentinelDtos.RolePrivilegesView, error)
}

/* -------------------------------------------------------------------------- */
/*                                   Groups                                   */
/* -------------------------------------------------------------------------- */

type GroupService interface {
	CreateGroup(ctx context.Context, orgID uuid.UUID, in sentinelDtos.CreateGroupsRequestDto) (uuid.UUID, error)
	ListGroups(ctx context.Context, orgID uuid.UUID) ([]sentinelDtos.GroupView, error)
}
