package coreEmployeeErrs

import "errors"

var (
	EmployeeNotFoundErr      = errors.New("employee not found")
	EmployeeAlreadyExistsErr = errors.New("employee already exists")
)

// Validation errors for the owner/super-admin creation flow.
var (
	EmployeeFirstNameRequiredErr = errors.New("first name is required")
	EmployeeEmailRequiredErr     = errors.New("email is required")
	EmployeePasswordRequiredErr  = errors.New("password is required")
	NotAnOwnerErr                = errors.New("only an owner can add another owner/super admin")
)
