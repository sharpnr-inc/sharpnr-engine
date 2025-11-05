package models

import "sharpnr-engine/core/constants"

// This is the Model or Struct for The Organisation.
// We save the Org Data in the Redis Temporary until user Verify it OTP
// Only Then we will push Data in the Main DB

type TempOrgRegistrationData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	OTP      string `json:"otp"`
	OrgType  constants.OrgType
	Password string
}
