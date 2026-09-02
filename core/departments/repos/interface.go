package orgDepartmentRepo

import orgDepartmentModels "sharpnr-engine/core/departments/models"

type OrgDepartmentRepositoryInterface interface {
	// Create a new department
	CreateDepartment(department *orgDepartmentModels.OrgDeparment) error

	// Get a department by ID
	GetDepartmentByID(id string) (*orgDepartmentModels.OrgDeparment, error)

	// Update a department
	UpdateDepartment(department *orgDepartmentModels.OrgDeparment) error

	// Delete a department by ID
	DeleteDepartment(id string) error

	// List all departments for an organization
	ListDepartments(organizationID string) ([]*orgDepartmentModels.OrgDeparment, error)
}
