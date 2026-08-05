// Package sentinelErrs defines the error types used in Sentinel.
// It provides a centralized location for error definitions, making it easier to manage and maintain error handling throughout the Sentinel codebase.
package sentinelServiceErrs

import "errors"

var (
	EmployeeRoleNotFoundErr = errors.New("employee role not found")
)
