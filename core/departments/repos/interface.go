package orgDepartmentRepo

import (
	"context"
	orgDepartmentsDtos "sharpnr-engine/core/departments/dtos"
	orgDepartmentModels "sharpnr-engine/core/departments/models"

	"github.com/google/uuid"
)

type OrgDepartmentRepositoryInterface interface {
	// Create a new department
	CreateDepartment(ctx context.Context, department *orgDepartmentModels.OrgDeparment) error

	// Get a department by ID
	GetDepartmentByID(ctx context.Context, departmentID uuid.UUID) (*orgDepartmentModels.OrgDeparment, error)

	// Update a department
	UpdateDepartment(ctx context.Context, department orgDepartmentsDtos.UpdateDepartmentRequestDto) error

	// Delete a department by ID
	DeleteDepartment(ctx context.Context, departmentID uuid.UUID) error

	// List all departments for an organization
	ListDepartments(ctx context.Context, organizationID uuid.UUID) ([]orgDepartmentModels.OrgDeparment, error)
}
