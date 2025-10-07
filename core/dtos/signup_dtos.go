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
	// Email of the organization.
	Email string `json:"email" example:"hotel@example.com"`
}

// Org_Verify_Otp_REQUEST_DTO represents the data required to verify an OTP for organization authentication or registration.
type Org_Verify_Otp_REQUEST_DTO struct {
	// OTP is the one-time password sent to the organization's email or mobile.
	OTP string `json:"otp" example:"123456"`
	// Email string
}

// Org_Verify_Otp_RESPONSE_DTO represents the response returned after successfully verifying an OTP for an organization.
type Org_Verify_Otp_RESPONSE_DTO struct {
	// Name of the organization.
	Name string `json:"org_name" example:"Sharpnr Hotel"`

	// Email of the organization.
	Email string `json:"email" example:"hotel@example.com"`

	// Sharpnr_Id is the unique identifier assigned to the organization.
	Sharpnr_Id string `json:"sharpnr_id" example:"SHRabc123def456"`
	// ORG_TYPE specifies the type of the organization.
	ORG_TYPE constants.OrgType `json:"org_type" example:"HOTEL"`
}

type Org_Complete_Regesistration_DTO struct {
	IsRegistered           string `json:"is_registered"`
	Address                string
	BankDetails            string
	OwnerDetails           string
	OrgRegistrationDetails string
}
