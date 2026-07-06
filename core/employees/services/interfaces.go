package coreEmployeeServices

import (
	"context"

	coreEmployeeContracts "sharpnr-engine/core/employees/contracts"

	"github.com/google/uuid"
)

type EmployeesGeneralServiceInterface interface {
	IsUserExists(ctx context.Context, employeeEmail string) (bool, error)
}

// This interface is used by authentication modules Employee will directly interact with the employee module through this interface.
// This interface will have employee Ops which are used by the other module interact with the employee module.
type EmployeeAuthenticationServiceInterface interface {
	EmployeesGeneralServiceInterface
	// Creates a default super admin for the organization. This is used when an organization is created. The super admin has all permissions and can manage the organization.
	CreateDefaultSuperAdmin(ctx context.Context, orgId uuid.UUID, rootEmail string, hasPass string) (uuid.UUID, error)
	UpdateSuperAdminEmployee(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, updates *coreEmployeeContracts.UpdateSuperAdminEmployeeInput) error

	// Verify user
	AuthenticateEmployee(ctx context.Context, email string, pass string) (*coreEmployeeContracts.EmployeeReader, error)
	GetEmployeeProfile(ctx context.Context, orgID uuid.UUID, empID uuid.UUID) (*coreEmployeeContracts.EmployeeProfileOutput, error)
	GetEmployeeDetails(ctx context.Context, orgID, emplID uuid.UUID) (*coreEmployeeContracts.EmployeeDetailsOutput, error)
}

/* ---------------------------------- <***> --------------------------------- */

// this is interface which will used by the HR module to interact with the employee module, this interface the have employee Ops used by the Org Hr
type EmployeesHRServiceInterface interface {

	// General service which will used by the other module to interact with the employee module,
	EmployeesGeneralServiceInterface
	ListOrgAllEmployees(ctx context.Context, orgID, emplID uuid.UUID) ([]*coreEmployeeContracts.OrgEmployeeList, error)
}

/* ---------------------------------- <***> --------------------------------- */

type EmployeeIamEngineServiceInterface interface {
	// this is General service which will used by the other module to interact with the employee module,
	EmployeesGeneralServiceInterface
	GetEmployeeProfile(ctx context.Context, orgID uuid.UUID, empID uuid.UUID) (*coreEmployeeContracts.EmployeeProfileOutput, error)
}
