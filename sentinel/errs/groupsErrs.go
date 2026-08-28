package sentinelServiceErrs

import "errors"

// Group Related Errors
var (
	GroupNameRequiredErr        = errors.New("group name is required")
	GroupAlreadyExistsErr       = errors.New("a group with this name already exists")
	ParentGroupNotFoundInOrgErr = errors.New("parent group not found in this organization")
	GroupNotFoundInOrgErr       = errors.New("group not found in this organization")
)
