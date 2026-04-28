package employee_services

import (
	"context"
	employee_contracts "sharpnr-engine/modules/hr/employees/contracts"
	sharphr_dtos "sharpnr-engine/modules/hr/employees/dtos"

	"github.com/google/uuid"
)

type EmployeeServiceInterface interface {
	CreateDefaultSuperAdmin(
		ctx context.Context,
		orgId uuid.UUID,
		rootEmail string,
		hasPass string,
	) (uuid.UUID, error)

	// This simply Tells if user exist or not
	IsUserExists(
		ctx context.Context,
		employeeEmail string,
	) (bool, error)

	// Verify user

	AuthenticateEmployee(ctx context.Context, email string, pass string) (*employee_contracts.EmployeeReader, error)

	GetEmployeeProfile(ctx context.Context, orgID uuid.UUID, empID uuid.UUID) (*employee_contracts.EmployeeProfileOutput, error)

	GetEmployeeDetails(ctx context.Context, orgID, emplID uuid.UUID) (*employee_contracts.EmployeeDetailsOutput, error)

	ListOrgAllEmployees(
		ctx context.Context,
		orgID, emplID uuid.UUID,
	) ([]*sharphr_dtos.OrgEmployeeList, error)
}

// This is the interface for the service that will used by HR module to interact with the employee service. This is used to decouple the HR module from the employee service implementation. This is also used to avoid circular dependencies between the HR module and the employee service.
type EmployeeSharpHRServiceInterface interface {
}
