package orgDepartmentService

import (
	"context"
	orgDepartmentsDtos "sharpnr-engine/core/departments/dtos"

	"github.com/google/uuid"
)

type OrgDepartmentServiceInterface interface {
	// Create a new department
	CreateDepartment(ctx context.Context, orgID, emplID uuid.UUID, request orgDepartmentsDtos.CreateDepartmentRequestDto) error
	// Update a department
	UpdateDepartment(ctx context.Context, orgID, emplID uuid.UUID, updateReq orgDepartmentsDtos.UpdateDepartmentRequestDto) error
	// Delete a department by ID
	DeleteDepartment(ctx context.Context, orgID, emplID uuid.UUID, DepartmentID uuid.UUID) error
	ListDepartments(ctx context.Context, orgID, emplID uuid.UUID) ([]orgDepartmentsDtos.DepartmentViewResponseDto, error)
	GetDepartmentByID(ctx context.Context, orgID, emplID uuid.UUID, departmentID uuid.UUID) (orgDepartmentsDtos.DepartmentViewResponseDto, error)
}
