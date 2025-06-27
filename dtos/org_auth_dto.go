package dto

// File: dtos/org_auth_dto.go
// This DTO represents the organization authentication details.

type OrgAuthDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}
