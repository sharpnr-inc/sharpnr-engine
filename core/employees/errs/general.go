package coreEmployeeErrs

import "errors"

var (
	EmployeeNotFoundErr = errors.New("employee not found")
	EmployeeAlreadyExistsErr = errors.New("employee already exists")
)
