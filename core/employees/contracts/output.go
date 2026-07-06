package coreEmployeeContracts

import (
	coreEmployeeConstants "sharpnr-engine/core/employees/constants"

	"github.com/google/uuid"
)

type EmployeeProfileOutput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	DOB       string `json:"dob"`
	Gender    string `json:"gender"`
}

type OrgEmployeeList struct {
	ID           uuid.UUID                            `json:"id"`
	EmployeeName string                               `json:"employeeName"`
	EmployeeID   string                               `json:"employeeID"`
	Status       coreEmployeeConstants.EmployeeStatus `json:"status"`
}

type EmployeeDetailsOutput struct {
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	ID        uuid.UUID `json:"id"`
}
