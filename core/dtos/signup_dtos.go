package dto

import "sharpnr-engine/core/constants"

// Org_Signup_REQUEST_DTO represents the data required to create or register an organization.
// This DTO is used when a new organization is signing up in the system.
type Org_Signup_REQUEST_DTO struct {
	// Name of the organization.
	Name string

	// Email address of the organization for contact and login purposes.
	Email string

	// Mobile number of the organization or primary contact person.
	Mobile string

	// ORG_TYPE specifies the type of organization.
	// Uses the OrgType constant defined in the constants package.
	ORG_TYPE constants.OrgType

	Password string
}

// Org_Signup_RESPONSE_DTO represents the response returned
// after successfully creating or registering an organization.
type Org_Signup_RESPONSE_DTO struct {
	// Token is the authentication token issued for the newly created organization.
	Token string

	// Message provides additional information about the signup operation,
	// such as success confirmation or any relevant notes.
	Message string

	Email string
}

// Org_Verify_Otp_REQUEST_DTO represents the data required to verify
// an OTP for organization authentication or registration.
//
// This DTO is typically used when an organization submits the OTP
// they received to complete verification.
type Org_Verify_Otp_REQUEST_DTO struct {
	// OTP is the one-time password sent to the organization's email or mobile.
	OTP string

	// Token is the temporary token associated with the OTP verification process.
	Token string
}

// Org_Verify_Otp_RESPONSE_DTO represents the response returned
// after successfully verifying an OTP for an organization.
//
// This DTO typically includes the organization's details, type, and
// an authentication token for subsequent requests.
type Org_Verify_Otp_RESPONSE_DTO struct {
	// Name of the organization.
	Name string

	// Email of the organization.
	Email string

	// Sharpnr_Id is the unique identifier assigned to the organization
	// within the Sharpnr system.
	Sharpnr_Id string

	// Token is the authentication token issued after successful OTP verification.
	Token string

	// ORG_TYPE specifies the type of the organization.
	// Uses the OrgType constant defined in the constants package.
	ORG_TYPE constants.OrgType
}
