package models

import "github.com/google/uuid"

// *This Struct is used by the Redis DB it used to hold data until verify the otp via email
// *This store hashed password and its currently supports 40 minutes TTL

type OrganizationSignUpStaging struct {
	ID       uuid.UUID
	Email    string
	Phone    string
	Otp      string
	Password string
}
