package employee_services

import (
	"context"
	employee_contracts "sharpnr-engine/modules/hr/employees/contracts"
	sharphr_dtos "sharpnr-engine/modules/hr/employees/dtos"

	"github.com/google/uuid"
)

type EmployeesIdentityServiceInterface interface {
	EmployeesGeneralServiceInterface
	// Creates a default super admin for the organization. This is used when an organization is created. The super admin has all permissions and can manage the organization.
	CreateDefaultSuperAdmin(ctx context.Context, orgId uuid.UUID, rootEmail string, hasPass string) (uuid.UUID, error)
	UpdateSuperAdminEmployee(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, updates *employee_contracts.UpdateSuperAdminEmployeeInput) error

	// Verify user
	AuthenticateEmployee(ctx context.Context, email string, pass string) (*employee_contracts.EmployeeReader, error)
	GetEmployeeProfile(ctx context.Context, orgID uuid.UUID, empID uuid.UUID) (*employee_contracts.EmployeeProfileOutput, error)
	GetEmployeeDetails(ctx context.Context, orgID, emplID uuid.UUID) (*employee_contracts.EmployeeDetailsOutput, error)
}

/* ---------------------------------- <***> --------------------------------- */

// this is interface which will used by the HR module to interact with the employee module, this interface the have employee Ops used by the Org Hr
type EmployeesHRServiceInterface interface {
	EmployeesGeneralServiceInterface
	ListOrgAllEmployees(ctx context.Context, orgID, emplID uuid.UUID) ([]*sharphr_dtos.OrgEmployeeList, error)
}

/* ---------------------------------- <***> --------------------------------- */

type EmployeesGeneralServiceInterface interface {
	IsUserExists(ctx context.Context, employeeEmail string) (bool, error)
}

/* ---------------------------------- <***> --------------------------------- */
