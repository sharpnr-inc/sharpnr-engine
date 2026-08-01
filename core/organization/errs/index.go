// This package conatains all the errors related Erros and Types related to Organization Signup and Verification Process.
package organizationErrs

var (
	OrganizationAlreadyExistsErr = "organization already exists"
	OrganizationNotFoundErr      = "organization not found"

	SignupSessionExpiredErr = "signup session expired"
	InvalidOTPError         = "invalid OTP"
)
