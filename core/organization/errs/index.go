package organizationErrs

import "errors"

var (
	OrganizationAlreadyExistsErr = errors.New("organization already exists")
	OrganizationNotFoundErr      = errors.New("organization not found")

	SignupSessionExpiredErr = errors.New("signup session not found or expired")
	InvalidOTPError         = errors.New("invalid OTP")
)
