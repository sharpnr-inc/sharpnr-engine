package sentinelservice

import (
	"context"

	sentinel_actions "sharpnr-engine/sentinel_engine/actions"
	sentinelDtos "sharpnr-engine/sentinel_engine/dtos"

	"github.com/google/uuid"
)

// SentinelService is the authorization authoring service injected into other
// modules (and exposed to the frontend through the orchestrator). It owns
// provisioning, the permission catalog, and CRUD over policies, roles, groups,
// and their attachments. The engine (evaluation) lives separately.
type SentinelService interface {
	// ProvisionOrganization sets up a new org's authorization defaults: the
	// full-access "Owner" role + its policy, the founding employee's owner
	// assignment, and the default "Stakeholders" group. Runs in one transaction
	// so an org is never left half-provisioned.
	ProvisionOrganization(ctx context.Context, orgID, firstUserID uuid.UUID) error

	// GetEmployeeRole resolves the role assigned to an employee within an org.
	GetEmployeeRole(ctx context.Context, orgID, emplID uuid.UUID) (sentinelDtos.RoleView, error)

	// Catalog returns the full permission catalog (the policy builder's picker).
	Catalog() []sentinel_actions.Permission

	/* -------------------------------- policies -------------------------------- */

	// CreatePolicy persists a raw Cedar policy (authored/validated in the FE),
	// revalidating that it compiles, and invalidates the org's engine cache.
	CreatePolicy(ctx context.Context, orgID uuid.UUID, in sentinelDtos.AddPolicyRequestDto) (uuid.UUID, error)
	ListPolicies(ctx context.Context, orgID uuid.UUID) ([]sentinelDtos.PolicyView, error)
	GetPolicy(ctx context.Context, orgID, policyID uuid.UUID) (sentinelDtos.PolicyView, error)

	/* ---------------------------------- roles --------------------------------- */

	CreateRole(ctx context.Context, orgID, emplID uuid.UUID, in sentinelDtos.CreateRoleInput) (uuid.UUID, error)
	ListRoles(ctx context.Context, orgID uuid.UUID) ([]sentinelDtos.RoleView, error)
	GetRolePrivileges(ctx context.Context, orgID, roleID uuid.UUID) (sentinelDtos.RolePrivilegesView, error)

	// AttachPolicies groups existing policies under a role; DetachPolicy removes
	// one grouping. Both leave the policy rows themselves untouched.
	AttachPolicies(ctx context.Context, orgID, emplID, roleID uuid.UUID, policyIDs []uuid.UUID) error
	DetachPolicy(ctx context.Context, orgID, roleID, policyID uuid.UUID) error

	/* --------------------------------- groups --------------------------------- */

	CreateGroup(ctx context.Context, orgID uuid.UUID, in sentinelDtos.CreateGroupsRequestDto) (uuid.UUID, error)
	ListGroups(ctx context.Context, orgID uuid.UUID) ([]sentinelDtos.GroupView, error)
}

// Compile-time guard: the concrete service must satisfy the interface.
var _ SentinelService = (*sentinelService)(nil)
