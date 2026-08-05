// Package authenticationEngineErrs defines the error types used in the
// authentication engine. It provides a centralized location for error
// definitions, making it easier to manage and maintain error handling
// throughout the authentication codebase.
package authenticationEngineErrs

import "errors"

var (
	AuthenticationErr = errors.New("authentication error")
)

// Request Context Related Errors
var (
	UserNotAuthenticatedErr    = errors.New("user not authenticated")
	OrgIDNotInContextErr       = errors.New("org_id not found in context")
	EmployeeIDNotInContextErr  = errors.New("employee_id not found in context")
	InvalidOrgIDFormatErr      = errors.New("invalid org_id format")
	InvalidEmployeeIDFormatErr = errors.New("invalid employee_id format")
)

// Authenticated User Profile Related Errors
var (
	OrgDetailsNotFoundErr = errors.New("organization details not found")
)
