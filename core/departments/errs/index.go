// Package orgDepartmentErrs defines the sentinel errors returned by the
// departments module. Keeping them in one place lets the service signal a
// precise failure and the controllers map it to a status code with errors.Is,
// rather than matching on message text.
package orgDepartmentErrs

import "errors"

var (
	ErrDepartmentNotFound      = errors.New("Department not found")
	ErrDepartmentAlreadyExists = errors.New("Department already exists")
	ErrNoEmployeesInDepartment = errors.New("No employees found in the department")
	ErrEmployeeNotInDepartment = errors.New("Employee not found in the department")

	// ErrDepartmentNameRequired guards the one field a department cannot be
	// created or renamed without.
	ErrDepartmentNameRequired = errors.New("Department name is required")

	// ErrInvalidDepartmentHead is a department_head that is not a UUID.
	ErrInvalidDepartmentHead = errors.New("Department head must be a valid employee id")

	// ErrDepartmentHasEmployees blocks deleting a department people are still
	// assigned to — move them out first, so nobody silently loses their
	// department.
	ErrDepartmentHasEmployees = errors.New("Department still has employees assigned to it")
)
