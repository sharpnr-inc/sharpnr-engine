package dto

import "sharpnr-engine/core/constants"

// Org_Signup_REQUEST_DTO represents the data required to create or register an organization.
type Org_Signup_REQUEST_DTO struct {
	// Name of the organization.
	Name string `json:"name" example:"Sharpnr Hotel"`

	// Email address of the organization for contact and login purposes.
	Email string `json:"email" example:"hotel@example.com"`

	// Mobile number of the organization or primary contact person.
	Mobile string `json:"mobile" example:"+911234567890"`

	// ORG_TYPE specifies the type of organization.
	// Uses the OrgType constant defined in the constants package.
	ORG_TYPE constants.OrgType `json:"org_type" example:"HOTEL"`

	// Password for the organization login.
	Password string `json:"password" example:"supersecure"`
}

// Org_Signup_RESPONSE_DTO represents the response returned after successfully creating or registering an organization.
type Org_Signup_RESPONSE_DTO struct {
	// Token is the authentication token issued for the newly created organization.
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`

	// Message provides additional information about the signup operation.
	Message string `json:"message" example:"OTP sent successfully"`

	// Email of the organization.
	Email string `json:"email" example:"hotel@example.com"`
}

// Org_Verify_Otp_REQUEST_DTO represents the data required to verify an OTP for organization authentication or registration.
type Org_Verify_Otp_REQUEST_DTO struct {
	// OTP is the one-time password sent to the organization's email or mobile.
	OTP string `json:"otp" example:"123456"`

	// Token is the temporary token associated with the OTP verification process.
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// Org_Verify_Otp_RESPONSE_DTO represents the response returned after successfully verifying an OTP for an organization.
type Org_Verify_Otp_RESPONSE_DTO struct {
	// Name of the organization.
	Name string `json:"name" example:"Sharpnr Hotel"`

	// Email of the organization.
	Email string `json:"email" example:"hotel@example.com"`

	// Sharpnr_Id is the unique identifier assigned to the organization.
	Sharpnr_Id string `json:"sharpnr_id" example:"SHRabc123def456"`

	// Token is the authentication token issued after successful OTP verification.
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`

	// ORG_TYPE specifies the type of the organization.
	ORG_TYPE constants.OrgType `json:"org_type" example:"HOTEL"`
}
