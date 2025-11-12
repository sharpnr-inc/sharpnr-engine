package hr_employee_repo

import (
	"context"
	sharphr_models "sharpnr-engine/modules/hr/hrModels"

	"github.com/google/uuid"
)

// This is Repository interface have Methods For Creating and Removing the Employee
type HREmployeeRepositoryInterface interface {
	// CreateEmployee persists a new employee in the database.
	// The employee must already have a valid OrganizationID (tenant).
	CreateEmployee(ctx context.Context, emp *sharphr_models.Employee) (*sharphr_models.Employee, error)
	// FindEmployeeById retrieves an employee by organization ID and employee ID.
	// Returns the employee (if found), a boolean indicating existence,
	// and an error if the query fails.
	FindEmployeeById(ctx context.Context, orgID, empID uuid.UUID) (*sharphr_models.Employee, bool, error)

	// FindByEmail retrieves an employee by organization ID and email address.
	// Returns the employee (if found), a boolean indicating existence,
	// and an error if the query fails.
	FindByEmail(ctx context.Context, email string) (*sharphr_models.Employee, bool, error)

	// ListEmployees retrieves all employees for a given organization.
	// Returns a slice of employees and an error if the query fails.
	ListEmployees(ctx context.Context, orgID uuid.UUID) ([]*sharphr_models.Employee, error)

	// UpdateEmployeeDetails updates the provided employee's details.
	// The returned value is the updated employee object.
	UpdateEmployeeDetails(ctx context.Context, emp *sharphr_models.Employee) (*sharphr_models.Employee, error)

	// RemoveEmployee deletes an employee from the given organization by employee ID.
	// Ensures that deletion is tenant-scoped by requiring orgID.
	RemoveEmployee(ctx context.Context, orgID, empID uuid.UUID) error
}
