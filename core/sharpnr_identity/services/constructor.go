package sharpnr_identity_auth_services

import (
	"context"
	employee_services "sharpnr-engine/modules/hr/employees/services"

	"github.com/google/uuid"
)

type sharpnrIdentityAuthService struct {
	employeeService employee_services.EmployeeServiceInterface
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
	employeeService employee_services.EmployeeServiceInterface,
) *sharpnrIdentityAuthService {
	return &sharpnrIdentityAuthService{
		employeeService: employeeService,
	}
}
