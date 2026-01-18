package employee_services

import (
	"context"

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

	IsUserExists(
		ctx context.Context,
		employee string,
	) (bool, error)

	AddEmployeeProvision()

	CreateSharpnrIdentityProfile(ctx context.Context, empID uuid.UUID) error
}
