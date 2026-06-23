package sentinelservice

import (
	"context"

	sentinel_actions "sharpnr-engine/sentinel_engine/actions"

	"github.com/google/uuid"
)

// SentinelService is the authorization service injected into other modules.
// For now it exposes provisioning (used at signup). Evaluation and grid
// management will be added here as those features land.
type SentinelService interface {
	// ProvisionOrganization sets up a new org's authorization defaults: it creates
	// the full-access "Owner" role and its policy, then assigns the founding
	// employee to that role. Runs in one transaction (joins an ambient tx if the
	// context carries one), so an org is never left half-provisioned.
	ProvisionOrganization(ctx context.Context, orgID, firstUserID uuid.UUID) error
	GetEmployeeRole(ctx context.Context, orgID, emplID uuid.UUID) error

	// RegenerateRolePolicies recompiles a role's attached managed policies into
	// the Cedar policies the engine evaluates, then invalidates the org's cache.
	// Call after any change to a role's attachments or an attached policy.
	RegenerateRolePolicies(ctx context.Context, orgID, roleID uuid.UUID) error

	// --- Managed-policy authoring + display (Layer 4) ---

	// Catalog returns the full permission catalog (the UI's picker source).
	Catalog() []sentinel_actions.Permission
	ListManagedPolicies(ctx context.Context, orgID uuid.UUID) ([]ManagedPolicyView, error)
	CreateManagedPolicy(ctx context.Context, orgID uuid.UUID, in CreateManagedPolicyInput) (uuid.UUID, error)
	ListRoles(ctx context.Context, orgID uuid.UUID) ([]RoleView, error)
	CreateRole(ctx context.Context, orgID, emplID uuid.UUID, in CreateRoleInput) (uuid.UUID, error)
	GetRolePrivileges(ctx context.Context, orgID, roleID uuid.UUID) (RolePrivilegesView, error)
	AttachPolicies(ctx context.Context, orgID, emplID, roleID uuid.UUID, policyIDs []uuid.UUID) error
	DetachPolicy(ctx context.Context, orgID, roleID, policyID uuid.UUID) error
}
