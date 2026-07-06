package sharpnr_identity_auth_dtos

import "github.com/google/uuid"

type UpdateSuperAdminDetailsInput struct {
	OrgID uuid.UUID `json:"org_id"`

	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	EmployeeID     string `json:"employee_id"`
	SecondaryEmail string `json:"secondary_email"`
	PhoneNumber    string `json:"phone_number"`
}
