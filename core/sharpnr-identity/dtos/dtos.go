package sharpnr_identity_auth_dtos

import "github.com/google/uuid"

type LoginRequest_dto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type WebLoginResponse_dto struct {
	OrgId      uuid.UUID `json:"org_id"`
	Email      string    `json:"email"`
	EmployeeID uuid.UUID `json:"employee_id"`
	Name       string    `json:"name"`
}

type MobileLoginRequest_dto struct {
}
