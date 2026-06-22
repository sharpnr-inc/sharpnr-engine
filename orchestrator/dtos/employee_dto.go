package orchestrator_dtos

import "github.com/google/uuid"

type EmployeeRole struct {
	ID          uuid.UUID
	Role        string
	Description string
}

type EmployeeOfOrganization struct {
	ID       uuid.UUID
	Name     string
	Nickname string
	Role     EmployeeRole
	Status   string
}
