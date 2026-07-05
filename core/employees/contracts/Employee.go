package coreEmployeeContracts

import "github.com/google/uuid"

type EmployeeReader struct {
	EmployeeID     uuid.UUID
	Name           string
	OrganizationID uuid.UUID
}

type DefaultSuperAdminEmployeeInput struct {
	Name           *string
	Email          string
	OrgID          uuid.UUID
	HashedPassword string
}

type UpdateSuperAdminEmployeeInput struct {
	OrganizationID uuid.UUID
	EmployeeID     *string
	FirstName      *string
	LastName       *string
	SecondaryEmail *string
}
