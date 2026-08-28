package sentinelServiceErrs

import "errors"

// Role Related Errors
var (
	RoleNameRequiredErr  = errors.New("role name is required")
	RoleAlreadyExistsErr = errors.New("a role with this name already exists")
	RoleNotFoundInOrgErr = errors.New("role not found in this organization")
)
