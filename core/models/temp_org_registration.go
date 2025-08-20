package models

// This is the Model or Struct for The Organisation.
// We save the Org Data in the Redis Temporary until user Verify it OTP
// Only Then we will push Data in the Main DB

type TempOrgRegistrationData struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	OTP         string `json:"otp"`
	OrgType     string
	Password    string
}
