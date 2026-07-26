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

type EmployeeDetailsResponseDto struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Department *string   `json:"department,omitempty"`
	CostCenter *string   `json:"costCenter,omitempty"`
	Role       *string   `json:"role,omitempty"`
	Status     string    `json:"status"`
	// further details will be added here this is just temporary
}

/* ---------------------------------- <***> --------------------------------- */

type AddEmployeeBulkRequestDto struct {
	FirstName    string     `json:"firstName"`
	MiddleName   *string    `json:"middleName,omitempty"`
	LastName     string     `json:"lastName"`
	NickName     *string    `json:"nickName,omitempty"`
	EmailAddress string     `json:"emailAddress"`
	PhoneNumber  string     `json:"phoneNumber"`
	EmployeeCode *string    `json:"employeeCode,omitempty"`
	DepartmentID *uuid.UUID `json:"departmentId,omitempty"`
}

/* ---------------------------------- <***> --------------------------------- */
