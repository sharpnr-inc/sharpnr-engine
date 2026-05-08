package org_setup_wizard_contracts

import "github.com/google/uuid"

type SuperAdminDetailsInput struct {
	StepID         uuid.UUID `json:"step_id"`
	OrgID          uuid.UUID `json:"org_id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	SecondaryEmail string    `json:"secondary_email"`
	SecondaryPhone string    `json:"secondary_phone"`
	EmployeeID     string    `json:"employee_id"`
}

type OrganizationDetailsSetupInput struct {
	StepID uuid.UUID `json:"step_id"`
	OrgID  uuid.UUID `json:"org_id"`
	Name   string    `json:"org_name"`
}
