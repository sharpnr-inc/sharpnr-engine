package employee_contracts

import "github.com/google/uuid"

type EmployeeReader struct {
	EmployeeID     uuid.UUID
	Name           string
	OrganizationID uuid.UUID
}
