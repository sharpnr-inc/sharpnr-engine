package interfaces

import (
	"context"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/core/models"

	"github.com/google/uuid"
)

// This is Interface for Service
// EmployeeService defines the business logic for managing employees
// within an organization (tenant).
type EmployeeService interface {
	// AddEmployeeInOrg creates a new employee for the given organization.
	// Returns the created employee and an error, if any.
	AddEmployeeInOrg(ctx context.Context, sharpnr_id string, req dto.CreateEmployeeInOrganisationRequestDTO) (*dto.CreateEmployeeInOrganisationResponseDTO, error)

	// RemoveEmployeeInOrg removes an employee from the organization by employee ID.
	// Ensures tenant isolation by requiring the organization ID.
	RemoveEmployeeInOrg(ctx context.Context, orgID, empID uuid.UUID) error

	// UpdateEmployeeDetails updates an existing employee's details within the organization.
	// Returns the updated employee object and an error, if any.
	UpdateEmployeeDetails(ctx context.Context, orgID uuid.UUID, emp *models.Employee) (*models.Employee, error)
}

// EmployeeRepository defines the contract for managing employees
// within a multi-tenant organization context.
type EmployeeRepository interface {
	// CreateEmployee persists a new employee in the database.
	// The employee must already have a valid OrganizationID (tenant).
	CreateEmployee(ctx context.Context, emp *models.Employee) (*models.Employee, error)
	// FindEmployeeById retrieves an employee by organization ID and employee ID.
	// Returns the employee (if found), a boolean indicating existence,
	// and an error if the query fails.
	FindEmployeeById(ctx context.Context, orgID, empID uuid.UUID) (*models.Employee, bool, error)

	// FindByEmail retrieves an employee by organization ID and email address.
	// Returns the employee (if found), a boolean indicating existence,
	// and an error if the query fails.
	FindByEmail(ctx context.Context, orgID uuid.UUID, email string) (*models.Employee, bool, error)

	// ListEmployees retrieves all employees for a given organization.
	// Returns a slice of employees and an error if the query fails.
	ListEmployees(ctx context.Context, orgID uuid.UUID) ([]*models.Employee, error)

	// UpdateEmployeeDetails updates the provided employee's details.
	// The returned value is the updated employee object.
	UpdateEmployeeDetails(ctx context.Context, emp *models.Employee) (*models.Employee, error)

	// RemoveEmployee deletes an employee from the given organization by employee ID.
	// Ensures that deletion is tenant-scoped by requiring orgID.
	RemoveEmployee(ctx context.Context, orgID, empID uuid.UUID) error
}
