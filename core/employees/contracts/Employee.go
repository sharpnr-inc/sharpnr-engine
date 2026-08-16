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
	EmployeeID     *uuid.UUID
	FirstName      *string
	LastName       *string
	SecondaryEmail *string
}

type AddNewSuperAdminInOrganizationRequestDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type AddEmployeeInBulkInOrganizationRequestDto struct {
	FirstName    string  `json:"firstName"`
	MiddleName   *string `json:"middleName,omitempty"`
	LastName     *string `json:"lastName,omitempty"`
	Email        string  `json:"email"`
	PhoneNumber  string  `json:"phoneNumber"`
	Nickname     *string `json:"nickname,omitempty"`
	EmployeeCode *string `json:"employeeCode,omitempty"`
}
