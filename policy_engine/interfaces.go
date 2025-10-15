package policyengine

import (
	"context"
)

// PolicyRepository defines low-level persistence methods for policies.
//
// This abstraction allows switching between different storage backends
// (e.g., PostgreSQL, Redis, file-based) without changing the business logic.
type PolicyRepository interface {
	// GetOrgPolicyByPolicyId fetches the organization’s policy from the datastore
	// based on its unique policy ID.
	// TODO: Define input parameters (e.g., policyId string) and return type.
	GetOrgPolicyByPolicyId()
	CreatePolicyOfOrg(ctx context.Context, orgId, sharpnr_id, employeeId string) error
	GetEmployeePolicyByEmployeeId(employeeID string) (*EmployeePolicy, error)
	GetEmployeePolicy(orgID string, employeeID string) (*EmployeeAccess, error)
}
