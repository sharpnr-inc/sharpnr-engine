package sharpnr_identity_auth_services

import (
	"context"
	sharpnr_identity_auth_repositories "sharpnr-engine/core/sharpnr_identity/repositories"
	hr_employee_repo "sharpnr-engine/modules/hr/employees/repo"

	"github.com/google/uuid"
)

type sharpnrIdentityAuthService struct {
	empIdentity  hr_employee_repo.EmployeeIdentityRepositoryInterface
	employeeAuth hr_employee_repo.EmployeeAuthRepositoryInterface
	identityRepo sharpnr_identity_auth_repositories.SharpnrIdentityAuthRepositoryInterface
}

// GetProfile implements SharpnrIdentityAuthServiceInterface.
func (s *sharpnrIdentityAuthService) GetProfile(ctx context.Context, orgID uuid.UUID, empID uuid.UUID) error {
	panic("unimplemented")
}

// GetRenderer implements SharpnrIdentityAuthServiceInterface.
func (s *sharpnrIdentityAuthService) GetRenderer(ctx context.Context, orgID uuid.UUID, employeeID uuid.UUID) error {
	panic("unimplemented")
}

// VerifyLoggedInEmployee implements SharpnrIdentityAuthServiceInterface.
func (s *sharpnrIdentityAuthService) VerifyLoggedInEmployee(ctx context.Context, orgID uuid.UUID, employeeID uuid.UUID) error {
	panic("unimplemented")
}

func NewSharpnrIdentityAuthService(
	empIdentity hr_employee_repo.EmployeeIdentityRepositoryInterface,
	employeeAuth hr_employee_repo.EmployeeAuthRepositoryInterface,
	identityRepo sharpnr_identity_auth_repositories.SharpnrIdentityAuthRepositoryInterface,
) *sharpnrIdentityAuthService {
	return &sharpnrIdentityAuthService{
		empIdentity:  empIdentity,
		employeeAuth: employeeAuth,
		identityRepo: identityRepo,
	}
}
