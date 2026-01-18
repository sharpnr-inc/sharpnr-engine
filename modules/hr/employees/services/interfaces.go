package employee_services

import (
	"context"

	"github.com/google/uuid"
)

type EmployeeServiceInterface interface {
	CreateDefaultSuperAdmin(
		ctx context.Context,
		orgId uuid.UUID,
		rootEmail string,
		hasPass string,
	) (uuid.UUID, error)
}
