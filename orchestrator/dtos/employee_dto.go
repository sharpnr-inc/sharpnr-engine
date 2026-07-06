package orchestrator_dtos

import (
	coreEmployeeConstants "sharpnr-engine/core/employees/constants"

	"github.com/google/uuid"
)

type EmployeeRole struct {
	ID          uuid.UUID `json:"id"`
	Role        string    `json:"role"`
	Description string    `json:"description"`
}

type EmployeeOfOrganization struct {
	ID       uuid.UUID                            `json:"id"`
	Name     string                               `json:"name"`
	Nickname string                               `json:"nickname"`
	Role     EmployeeRole                         `json:"role"`
	Status   coreEmployeeConstants.EmployeeStatus `json:"status"`
}
