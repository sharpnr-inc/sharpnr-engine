package policyengine

import (
	"context"
)

type policyEngine struct {
	repo PolicyRepository
}

// AddOrgOwner implements PolicyEngine.
func (p *policyEngine) AddOrgOwner(ctx context.Context, orgID string, employeeID string) error {
	panic("unimplemented")
}

// AddSuperAdmin implements PolicyEngine.
func (p *policyEngine) AddSuperAdmin(ctx context.Context, orgID string, employeeID string) error {
	panic("unimplemented")
}

// GetEmployeeRoles implements PolicyEngine.
func (p *policyEngine) GetEmployeeRoles(ctx context.Context, orgID string, employeeID string) ([]Role, error) {
	panic("unimplemented")
}

// GetPolicy implements PolicyEngine.
func (p *policyEngine) GetPolicy(ctx context.Context, orgID string) (*OrgDefaultPolicy, error) {
	panic("unimplemented")
}

// HasPermission implements PolicyEngine.
func (p *policyEngine) HasPermission(ctx context.Context, orgID string, employeeID string, resource string, permission Permission) (bool, error) {
	panic("unimplemented")
}

// HasRole implements PolicyEngine.
func (p *policyEngine) HasRole(ctx context.Context, orgID string, employeeID string, role Role) (bool, error) {
	panic("unimplemented")
}

// IsOrgOwner implements PolicyEngine.
func (p *policyEngine) IsOrgOwner(ctx context.Context, orgID string, employeeID string) (bool, error) {
	panic("unimplemented")
}

// IsSuperAdmin implements PolicyEngine.
func (p *policyEngine) IsSuperAdmin(ctx context.Context, orgID string, employeeID string) (bool, error) {
	panic("unimplemented")
}

// RemoveOrgOwner implements PolicyEngine.
func (p *policyEngine) RemoveOrgOwner(ctx context.Context, orgID string, employeeID string) error {
	panic("unimplemented")
}

// RemoveSuperAdmin implements PolicyEngine.
func (p *policyEngine) RemoveSuperAdmin(ctx context.Context, orgID string, employeeID string) error {
	panic("unimplemented")
}

func NewPolicyEngine(repo PolicyRepository) PolicyEngine {
	return &policyEngine{repo: repo}
}
