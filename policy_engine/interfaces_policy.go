package policyengine

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
	CreateDefaultPolicy(org_id, orgType string) (OrgDefaultPolicy, error)

	// UpdatePolicy updates and returns the policy for the given organization.
	// orgId is the unique identifier for the organization.
	UpdatePolicy(orgId string) (OrgDefaultPolicy, error)
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
