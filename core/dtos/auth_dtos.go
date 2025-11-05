package dto

import "github.com/google/uuid"

type LoginRequestDTO struct {
	Email    string
	Password string
}

type LoginResponseDTO struct {
	OrgId         string
	OrgType       string
	EmployeeID    uuid.UUID
	EmployeeEmail string
}
