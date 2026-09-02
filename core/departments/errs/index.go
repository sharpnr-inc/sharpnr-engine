package orgDepartmentErrs

import "errors"

var (
	ErrDepartmentNotFound      = errors.New("Department not found")
	ErrDepartmentAlreadyExists = errors.New("Department already exists")
	ErrNoEmployeesInDepartment = errors.New("No employees found in the department")
	ErrEmployeeNotInDepartment = errors.New("Employee not found in the department")
)
