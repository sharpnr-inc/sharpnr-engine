package employee_services

import (
	"context"
	employee_contracts "sharpnr-engine/modules/hr/employees/contracts"

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

type EmployeePublicServiceInterface interface {
	CheckEmpExistByEmail(ctx context.Context, orgId uuid.UUID, email string) (bool, error)

	GetEmpByEmail(ctx context.Context, email string) (interface{}, error)
}

type EmployeeInternalServiceInterface interface {
	CreateSuperDefaultAdmin(ctx context.Context, req employee_contracts.DefaultSuperAdminEmployee) error
}
