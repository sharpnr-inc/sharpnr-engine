package authenticationEngineDtos

import "github.com/google/uuid"

type UserDataResponseDto struct {
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	ID        uuid.UUID `json:"id"`
	Role      string    `json:"role"`
	Org       OrgData   `json:"org_data"`
}
