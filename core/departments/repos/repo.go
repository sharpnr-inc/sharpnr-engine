package orgDepartmentRepo

import (
	"context"
	orgDepartmentsDtos "sharpnr-engine/core/departments/dtos"
	orgDepartmentModels "sharpnr-engine/core/departments/models"

	"github.com/google/uuid"
)

// CreateDepartment implements [OrgDepartmentRepositoryInterface].
func (o *orgDepartmentRepository) CreateDepartment(ctx context.Context, department *orgDepartmentModels.OrgDeparment) error {
	panic("unimplemented")
}

// DeleteDepartment implements [OrgDepartmentRepositoryInterface].
func (o *orgDepartmentRepository) DeleteDepartment(ctx context.Context, departmentID uuid.UUID) error {
	panic("unimplemented")
}

// GetDepartmentByID implements [OrgDepartmentRepositoryInterface].
func (o *orgDepartmentRepository) GetDepartmentByID(ctx context.Context, departmentID uuid.UUID) (*orgDepartmentModels.OrgDeparment, error) {
	panic("unimplemented")
}

// ListDepartments implements [OrgDepartmentRepositoryInterface].
func (o *orgDepartmentRepository) ListDepartments(ctx context.Context, organizationID uuid.UUID) ([]orgDepartmentModels.OrgDeparment, error) {
	panic("unimplemented")
}

// UpdateDepartment implements [OrgDepartmentRepositoryInterface].
func (o *orgDepartmentRepository) UpdateDepartment(ctx context.Context, department orgDepartmentsDtos.UpdateDepartmentRequestDto) error {
	panic("unimplemented")
}
