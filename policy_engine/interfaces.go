package policyengine

import "context"

// PolicySetup defines high-level operations for setting up and maintaining
// organizational policies.
//
// Typical flow:
//   - CreateDefaultPolicy is called when a new organization is registered,
//     providing it with a base policy according to org type.
//   - UpdatePolicy can be used to refresh or adjust the policy when
//     organizational details change.
type PolicySetup interface {
	// CreateDefaultPolicy generates a new default policy for the given organization.
	// The org_id uniquely identifies the organization, and orgType determines
	// the policy template to apply.
	CreateOrgDefaultPolicy(org_id, orgType string) (*OrgDefaultPolicy, error)

	// UpdatePolicy updates and returns the policy for the given organization.
	// orgId is the unique identifier for the organization.
	UpdatePolicy(orgId string) (*OrgDefaultPolicy, error)
}

// PolicyRepository defines low-level persistence methods for policies.
//
// This abstraction allows switching between different storage backends
// (e.g., PostgreSQL, Redis, file-based) without changing the business logic.
type PolicyRepository interface {
	// GetOrgPolicyByPolicyId fetches the organization’s policy from the datastore
	// based on its unique policy ID.
	// TODO: Define input parameters (e.g., policyId string) and return type.
	GetOrgPolicyByPolicyId()
}

type PolicyEngine interface {
	// Permission checks
	IsSuperAdmin(ctx context.Context, orgID, employeeID string) (bool, error)
	IsOrgOwner(ctx context.Context, orgID, employeeID string) (bool, error)
	HasRole(ctx context.Context, orgID, employeeID string, role Role) (bool, error)
	HasPermission(ctx context.Context, orgID, employeeID, resource string, permission Permission) (bool, error)

	// Policy modifications
	AddSuperAdmin(ctx context.Context, orgID, employeeID string) error
	RemoveSuperAdmin(ctx context.Context, orgID, employeeID string) error
	AddOrgOwner(ctx context.Context, orgID, employeeID string) error
	RemoveOrgOwner(ctx context.Context, orgID, employeeID string) error

	// Policy retrieval
	GetPolicy(ctx context.Context, orgID string) (*OrgDefaultPolicy, error)
	GetEmployeeRoles(ctx context.Context, orgID, employeeID string) ([]Role, error)
}
