// Package sentinelServiceErrs defines the error types used in Sentinel.
// It provides a centralized location for error definitions, making it easier to manage and maintain error handling throughout the Sentinel codebase.
package sentinelServiceErrs

import "errors"

// Employee Role Related Errors
var (
	EmployeeRoleNotFoundErr = errors.New("employee role not found")
	InvalidEmployeeErr      = errors.New("there is no employee with the given ID in this organization")
)
