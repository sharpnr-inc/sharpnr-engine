package dto

import (
	"sharpnr-engine/core/constants"

	"github.com/google/uuid"
)

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

// RegistrationDetails contains information about GST, CIN, and registration status.
type RegistrationDetails struct {
	IsRegister bool   `json:"is_register"`          // true if already registered
	GSTNumber  string `json:"gst_number,omitempty"` // GST number, optional
	CINNumber  string `json:"cin_number,omitempty"` // CIN number, optional
}

// AddressDetails contains the organization's address information.
type AddressDetails struct {
	Country     string `json:"country"`
	City        string `json:"city"`                   // City name
	State       string `json:"state"`                  // State name
	Pincode     string `json:"pincode"`                // Postal code
	AddressLine string `json:"address_line,omitempty"` // Street address, optional
	Landmark    string `json:"landmark,omitempty"`     // Optional landmark
}

// BankDetails represents one bank account of the organization.
type BankDetails struct {
	BankName      string `json:"bank_name"`              // Bank name
	AccountNumber string `json:"account_number"`         // Account number
	IFSCCode      string `json:"ifsc_code"`              // IFSC code
	BranchName    string `json:"branch_name,omitempty"`  // Optional branch name
	IsPrimary     bool   `json:"is_primary"`             // Primary account
	IsSecondary   bool   `json:"is_secondary,omitempty"` // Optional secondary flag
}

// OrgOwnerDetails represents an owner or stakeholder of the organization.
type OrgOwnerDetails struct {
	OwnerName  string `json:"owner_name"`            // Full name of owner
	OwnerEmail string `json:"owner_email,omitempty"` // Optional email
	OwnerPhone string `json:"owner_phone,omitempty"` // Optional phone number
}

type SuperAdminDetails struct {
	Name string
}

// Org_Complete_Regesistration_REQ_DTO is the request payload for completing organization registration.
type Org_Complete_Regesistration_REQ_DTO struct {
	Registration RegistrationDetails `json:"registration_details"`
	Address      AddressDetails      `json:"address_details"`
	Is_owner     bool                `json:"is_owner"`
	OwnerDetails []OrgOwnerDetails   `json:"owner_details,omitempty"` // Optional if no owners yet
}

type Org_Complete_Regesistration_RESPONSE_DTO struct {
	OrgID        string    `json:"org_id"`
	OrgName      string    `json:"org_name"`
	EmployeeID   uuid.UUID `json:"employee_Id"`
	EmployeeName string    `json:"employee_name"`
	EmployeeRole string    `json:"employee_role"`
}
