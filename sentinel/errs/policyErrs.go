package sentinelServiceErrs

import "errors"

// Policy Related Errors
var (
	PolicyNameRequiredErr    = errors.New("policy name is required")
	PolicyTextRequiredErr    = errors.New("policy text is required")
	PolicyDoesNotCompileErr  = errors.New("policy does not compile")
	EmptyPolicyStatementsErr = errors.New("policy text contains no policy statements")
	PolicyAlreadyExistsErr   = errors.New("a policy with this name already exists")
	PolicyNotFoundInOrgErr   = errors.New("policy not found in this organization")
)
