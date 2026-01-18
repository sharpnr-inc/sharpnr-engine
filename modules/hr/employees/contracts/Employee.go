package employee_contracts

import "github.com/google/uuid"

type EmployeeReader struct {
	EmployeeID     uuid.UUID
	Name           string
	OrganizationID uuid.UUID
}

type DefaultSuperAdminEmployee struct {
	Name           *string
	Email          string
	OrgID          uuid.UUID
	HashedPassword string
}
