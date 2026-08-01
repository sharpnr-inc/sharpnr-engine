// Package orgSignUpErrs provides domain-specific error definitions and constructs
// for the Organization Sign-Up lifecycle.
//
// All sentinel errors in this package can be checked using standard library
// errors.Is, while custom struct errors expose typed fields accessible
// via errors.As.
package orgSignUpErrs

import "errors"

var (
	OrganizationAlreadyExists = errors.New("Organization Already Exists")
)
