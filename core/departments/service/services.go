package orgDepartmentService

import (
	"context"
	orgDepartmentsDtos "sharpnr-engine/core/departments/dtos"

	"github.com/google/uuid"
)

// CreateDepartment implements [OrgDepartmentServiceInterface].
func (c *coreOrgDepartmentService) CreateDepartment(ctx context.Context, orgID, emplID uuid.UUID, request orgDepartmentsDtos.CreateDepartmentRequestDto) error {
	panic("unimplemented")
}

// DeleteDepartment implements [OrgDepartmentServiceInterface].
func (c *coreOrgDepartmentService) DeleteDepartment(ctx context.Context, orgID, emplID uuid.UUID, DepartmentID uuid.UUID) error {
	panic("unimplemented")
}

// GetDepartmentByID implements [OrgDepartmentServiceInterface].
func (c *coreOrgDepartmentService) GetDepartmentByID(ctx context.Context, orgID, emplID uuid.UUID, departmentID uuid.UUID) (orgDepartmentsDtos.DepartmentViewResponseDto, error) {
	panic("unimplemented")
}

// ListDepartments implements [OrgDepartmentServiceInterface].
func (c *coreOrgDepartmentService) ListDepartments(ctx context.Context, orgID, emplID uuid.UUID) ([]orgDepartmentsDtos.DepartmentViewResponseDto, error) {
	panic("unimplemented")
}

// UpdateDepartment implements [OrgDepartmentServiceInterface].
func (c *coreOrgDepartmentService) UpdateDepartment(ctx context.Context, orgID, emplID uuid.UUID, updateReq orgDepartmentsDtos.UpdateDepartmentRequestDto) error {
	panic("unimplemented")
}
