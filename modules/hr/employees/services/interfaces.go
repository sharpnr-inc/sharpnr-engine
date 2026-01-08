package hr_employee_services

import (
	"context"
	sharphr_models "sharpnr-engine/modules/hr/hrModels"
	sharphrdtos "sharpnr-engine/modules/hr/hr_dtos"

	"github.com/google/uuid"
)

type HrEmployeeServiceInterface interface {
	// AddEmployeeInOrg creates a new employee for the given organization.
	// Returns the created employee and an error, if any.
	AddEmployeeInOrg(ctx context.Context, sharpnr_id string, req sharphrdtos.CreateEmployeeInOrganization_RequestDTO) (*sharphrdtos.CreateEmployeeInOrganizationResponseDTO, error)

	// RemoveEmployeeInOrg removes an employee from the organization by employee ID.
	// Ensures tenant isolation by requiring the organization ID.
	RemoveEmployeeInOrg(ctx context.Context, orgID, empID uuid.UUID) error

	// UpdateEmployeeDetails updates an existing employee's details within the organization.
	// Returns the updated employee object and an error, if any.
	UpdateEmployeeDetails(ctx context.Context, orgID uuid.UUID, emp *sharphr_models.Employee) (*sharphr_models.Employee, error)

	CreateNewEmployee(
		ctx context.Context,
		orgID uuid.UUID,
		req sharphrdtos.CreateEmployeeInOrganization_RequestDTO,
	) error

	CreateDefaultSuperAdmin(ctx context.Context, orgId uuid.UUID, rootEmail string) (uuid.UUID, error)
}
