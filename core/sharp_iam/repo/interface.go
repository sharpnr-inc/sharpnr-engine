package sharpIam_repo

import sharp_iam_models "sharpnr-engine/core/sharp_iam/models"

type SharpIamRepositoryInterface interface {
	// Policy Management
	CreatePolicy(policy *sharp_iam_models.SharpIamPolicies) error
	GetPolicyByID(id string) (*sharp_iam_models.SharpIamPolicies, error)
	UpdatePolicy(policy *sharp_iam_models.SharpIamPolicies) error
	DeletePolicy(id string) error
	ListPolicies(orgID string, domain string) ([]*sharp_iam_models.SharpIamPolicies, error)

	// Role Management
	CreateRole(role *sharp_iam_models.SharpIamRoles) error
	GetRoleByID(id string) (*sharp_iam_models.SharpIamRoles, error)
	UpdateRole(role *sharp_iam_models.SharpIamRoles) error
	DeleteRole(id string) error
	ListRoles(orgID string, domain string) ([]*sharp_iam_models.SharpIamRoles, error)

	// Employee-Role Assignment
	AssignRoleToEmployee(employeeID string, roleID string) error
	RemoveRoleFromEmployee(employeeID string, roleID string) error
	ListRolesForEmployee(employeeID string) ([]*sharp_iam_models.SharpIamRoles, error)
}
