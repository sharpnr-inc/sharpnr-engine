package sharpIam_repo

import (
	"context"
	sharp_iam_models "sharpnr-engine/core/sharp_iam/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SharpIamRepositoryInterface interface {
	// Policy Management
	CreatePolicy(ctx context.Context, policy *sharp_iam_models.SharpIamPolicies) error
	GetPolicyByID(ctx context.Context, id uuid.UUID) (*sharp_iam_models.SharpIamPolicies, error)
	UpdatePolicy(ctx context.Context, policy *sharp_iam_models.SharpIamPolicies) error
	DeletePolicy(ctx context.Context, id uuid.UUID) error
	ListPolicies(ctx context.Context, orgID uuid.UUID, domain string) ([]*sharp_iam_models.SharpIamPolicies, error)

	GetActivePolicies(ctx context.Context, employeeID uuid.UUID, roleIDs []uuid.UUID) ([]*sharp_iam_models.SharpIamPolicies, error)

	// 2. Hierarchy Resolution (The "Ancestry" Proof)
	// Given a Hotel ID, return [HospitalityDomainID, PropertyID, LegalEntityID, OrgID]
	GetResourceParents(resType string, resID uuid.UUID) ([]uuid.UUID, error)

	// 3. Batch Operations
	// Useful for the "Signup" logic where you create an Org and a SuperAdmin Role at once
	CreateInitialOrgPermissions(orgID uuid.UUID, adminID uuid.UUID) error

	// Role Management
	CreateRole(ctx context.Context, role *sharp_iam_models.SharpIamRoles) error
	GetRoleByID(ctx context.Context, id uuid.UUID) (*sharp_iam_models.SharpIamRoles, error)
	UpdateRole(ctx context.Context, role *sharp_iam_models.SharpIamRoles) error
	DeleteRole(ctx context.Context, id uuid.UUID) error
	ListRoles(ctx context.Context, orgID uuid.UUID, domain string) ([]*sharp_iam_models.SharpIamRoles, error)

	// Employee-Role Assignment
	AssignRoleToEmployee(ctx context.Context, employeeID uuid.UUID, roleID uuid.UUID) error
	RemoveRoleFromEmployee(ctx context.Context, employeeID uuid.UUID, roleID uuid.UUID) error
	ListRolesForEmployee(ctx context.Context, employeeID uuid.UUID) ([]*sharp_iam_models.SharpIamRoles, error)
}

type sharpIamRepository struct {
	db *gorm.DB
}

// CreateInitialOrgPermissions implements SharpIamRepositoryInterface.
func (r *sharpIamRepository) CreateInitialOrgPermissions(orgID uuid.UUID, adminID uuid.UUID) error {
	panic("unimplemented")
}

// GetActivePolicies implements SharpIamRepositoryInterface.
func (r *sharpIamRepository) GetActivePolicies(ctx context.Context, employeeID uuid.UUID, roleIDs []uuid.UUID) ([]*sharp_iam_models.SharpIamPolicies, error) {
	panic("unimplemented")
}

// GetResourceParents implements SharpIamRepositoryInterface.
func (r *sharpIamRepository) GetResourceParents(resType string, resID uuid.UUID) ([]uuid.UUID, error) {
	panic("unimplemented")
}

func NewSharpIamRepository(db *gorm.DB) *sharpIamRepository {
	return &sharpIamRepository{db: db}
}
