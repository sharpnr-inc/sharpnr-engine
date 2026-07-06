package sharpnr_identity_auth_dtos

import "github.com/google/uuid"

type ProfileResponse_dto struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DOB       string `json:"dob"`
	Address   string `json:"address"`
	Gender    string `json:"gender"`
}

type UserDataResponseDto struct {
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	ID        uuid.UUID `json:"id"`
	Role      string    `json:"role"`
	Org       OrgData   `json:"org_data"`
}
