package orgDepartmentRepo

import (
	"context"
	orgDepartmentsDtos "sharpnr-engine/core/departments/dtos"
	orgDepartmentModels "sharpnr-engine/core/departments/models"

	"github.com/google/uuid"
)

// OrgDepartmentRepositoryInterface is the departments module's persistence
// port. Every method carries organizationID because the tenant predicate
// belongs on the query, not on the caller's memory: a department id from
// another org must read as "not found" rather than resolve.
type OrgDepartmentRepositoryInterface interface {
	// Create a new department
	CreateDepartment(ctx context.Context, department *orgDepartmentModels.OrgDeparment) error

	// Get a department by ID
	GetDepartmentByID(ctx context.Context, organizationID, departmentID uuid.UUID) (*orgDepartmentModels.OrgDeparment, error)

	// Update a department
	UpdateDepartment(ctx context.Context, organizationID uuid.UUID, department orgDepartmentsDtos.UpdateDepartmentRequestDto) error

	// Delete a department by ID
	DeleteDepartment(ctx context.Context, organizationID, departmentID uuid.UUID) error

	// List all departments for an organization
	ListDepartments(ctx context.Context, organizationID uuid.UUID) ([]orgDepartmentModels.OrgDeparment, error)

	// DepartmentExists reports whether a live department in the organization
	// already uses this name or code. excludeID is skipped, so an update can
	// ask the question without matching the row it is updating; pass uuid.Nil
	// on create.
	DepartmentExists(ctx context.Context, organizationID uuid.UUID, name, code string, excludeID uuid.UUID) (bool, error)

	// CountEmployeesInDepartment is the roster size, and the guard that blocks
	// deleting a department people are still assigned to.
	CountEmployeesInDepartment(ctx context.Context, organizationID, departmentID uuid.UUID) (int64, error)
}
