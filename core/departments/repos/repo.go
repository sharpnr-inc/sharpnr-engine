package orgDepartmentRepo

import orgDepartmentModels "sharpnr-engine/core/departments/models"

// CreateDepartment implements [OrgDepartmentRepositoryInterface].
func (o *orgDepartmentRepository) CreateDepartment(department *orgDepartmentModels.OrgDeparment) error {
	panic("unimplemented")
}

// DeleteDepartment implements [OrgDepartmentRepositoryInterface].
func (o *orgDepartmentRepository) DeleteDepartment(id string) error {
	panic("unimplemented")
}

// GetDepartmentByID implements [OrgDepartmentRepositoryInterface].
func (o *orgDepartmentRepository) GetDepartmentByID(id string) (*orgDepartmentModels.OrgDeparment, error) {
	panic("unimplemented")
}

// ListDepartments implements [OrgDepartmentRepositoryInterface].
func (o *orgDepartmentRepository) ListDepartments(organizationID string) ([]*orgDepartmentModels.OrgDeparment, error) {
	panic("unimplemented")
}

// UpdateDepartment implements [OrgDepartmentRepositoryInterface].
func (o *orgDepartmentRepository) UpdateDepartment(department *orgDepartmentModels.OrgDeparment) error {
	panic("unimplemented")
}
