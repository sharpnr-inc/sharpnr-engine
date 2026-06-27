package sentinelservice

import (
	"context"

	sentinelDtos "sharpnr-engine/sentinel/dtos"
	sentinelengine_models "sharpnr-engine/sentinel/models"
	sentinel_resources "sharpnr-engine/sentinel/resources"

	"github.com/google/uuid"
)

// SentinelService is the authorization authoring service injected into other
// modules (and exposed to the frontend through the orchestrator). It owns
// provisioning, the policy builder's palettes (principals + resources), and CRUD
// over policies, roles, groups, and their attachments. The engine (evaluation)
// lives separately.
type SentinelService interface {
	// ProvisionOrganization sets up a new org's authorization defaults: the
	// full-access "Owner" role + its policy, the founding employee's owner
	// assignment, and the default "Stakeholders" group. Runs in one transaction
	// so an org is never left half-provisioned.
	ProvisionOrganization(ctx context.Context, orgID, firstUserID uuid.UUID) error

	// GetEmployeeRole resolves the role assigned to an employee within an org.
	GetEmployeeRole(ctx context.Context, orgID, emplID uuid.UUID) (sentinelDtos.RoleView, error)

	// GetPrincipals returns the policy builder's principal palette: the static
	// Cedar entity types (User/Role/Group) plus the org's live roles and groups.
	GetPrincipals(ctx context.Context, orgID uuid.UUID) (sentinelDtos.PrincipalsView, error)

	/* ------------------------------- resources -------------------------------- */
	// The resource-centric policy builder: list resource types, then drill into a
	// type's actions, attributes, and live instances.

	ListResources() []sentinel_resources.Summary
	GetResourceActions(resourceType string) ([]sentinel_resources.Action, error)
	GetResourceAttributes(resourceType string) ([]sentinel_resources.Attribute, error)
	ListResourceInstances(ctx context.Context, orgID uuid.UUID, resourceType string) ([]sentinel_resources.Ref, error)

	/* -------------------------------- policies -------------------------------- */

	// CreatePolicy persists a raw Cedar policy (authored/validated in the FE),
	// revalidating that it compiles, and invalidates the org's engine cache.
	CreatePolicy(ctx context.Context, orgID uuid.UUID, in sentinelDtos.AddPolicyRequestDto) (uuid.UUID, error)
	// ListPolicies returns an org's policies; an empty scope means all buckets,
	// otherwise only that bucket (organization|role|group|user).
	ListPolicies(ctx context.Context, orgID uuid.UUID, scope sentinelengine_models.PolicyScope) ([]sentinelDtos.PolicyView, error)
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
