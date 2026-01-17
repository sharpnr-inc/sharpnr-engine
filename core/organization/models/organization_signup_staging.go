package models

import "github.com/google/uuid"

type OrganizationSignUpStaging struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Phone    string
	Otp      string
	Password string
}
