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

// AddNewSuperAdminInOrganizationRequestDto is the payload for adding a second
// owner/super-admin to a live organization.
//
// TEMP: Password is here only because the account is created directly for now.
// The intended flow mails the invitee a provisioning link and lets them set
// their own password — drop this field when that lands. See
// coreEmployeeServices.AddNewOwnerOrSuperAdmin.
type AddNewSuperAdminInOrganizationRequestDto struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
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
