package sharpnr_identity_auth_services

import (
	"context"
	employee_services "sharpnr-engine/modules/hr/employees/services"

	"github.com/google/uuid"
)

type sharpnrIdentityAuthService struct {
	employeeService employee_services.EmployeesIdentityServiceInterface
}

// VerifyLoggedInEmployee implements SharpnrIdentityAuthServiceInterface.
func (s *sharpnrIdentityAuthService) VerifyLoggedInEmployee(ctx context.Context, orgID uuid.UUID, employeeID uuid.UUID) error {
	panic("unimplemented")
}

func NewSharpnrIdentityAuthService(
	employeeService employee_services.EmployeesIdentityServiceInterface,
) *sharpnrIdentityAuthService {
	return &sharpnrIdentityAuthService{
		employeeService: employeeService,
	}
}
