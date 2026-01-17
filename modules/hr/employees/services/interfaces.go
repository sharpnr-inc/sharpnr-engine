package employee_services

import (
	"context"
	employee_models "sharpnr-engine/modules/hr/employees/models"

	"github.com/google/uuid"
)

type EmployeeServiceInterface interface {
	// AddEmployeeInOrg creates a new employee for the given organization.
	// Returns the created employee and an error, if any.
	AddEmployeeInOrg(ctx context.Context, sharpnr_id string) error

	// RemoveEmployeeInOrg removes an employee from the organization by employee ID.
	// Ensures tenant isolation by requiring the organization ID.
	RemoveEmployeeInOrg(ctx context.Context, orgID, empID uuid.UUID) error

	// UpdateEmployeeDetails updates an existing employee's details within the organization.
	// Returns the updated employee object and an error, if any.
	UpdateEmployeeDetails(ctx context.Context, orgID uuid.UUID, emp *employee_models.Employee) (*employee_models.Employee, error)

	CreateNewEmployee(
		ctx context.Context,
		orgID uuid.UUID,
	) error

	CreateDefaultSuperAdmin(
		ctx context.Context,
		orgId uuid.UUID,
		rootEmail string,
		hasPass string,
	) (uuid.UUID, error)
}

// Employee Public Service Interface
// This interface is expose internal method for employee
/*

 */
type EmployeePublicServiceInterface interface {
	CheckEmpExistByEmail(ctx context.Context, orgId uuid.UUID, email string) error
	GetEmpByEmail(ctx context.Context, email string) (interface{}, error)
}
