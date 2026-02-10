package employee_services

import (
	"context"
	employee_contracts "sharpnr-engine/modules/hr/employees/contracts"

	"github.com/google/uuid"
)

type EmployeeServiceInterface interface {

	// Creates the Default Super Admin and
	// This Takes the
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
}
