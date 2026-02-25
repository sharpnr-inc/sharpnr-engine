package models

import "github.com/google/uuid"

// *This Struct is used by the Redis DB it used to hold data until verify the otp via email
// *This store hashed password and its currently supports 40 minutes TTL
// ? Reminder :--> To configure the application so in the development mode it supports the 40 minutes
// ? and in the Production it needs to setup the 10 minutes TTL

type OrganizationSignUpStaging struct {
	ID       uuid.UUID
	Email    string
	Phone    string
	Otp      string
	Password string
}

// ** This table is not decided yet normally or by convention its the job of Setup Wizard to answer
// ** if the setup of the organization is Complete, For temporary period we are saving the data
// ** if the organization is set up or not

type OrganizationSetupStatus struct {
	ID              uuid.UUID
	OrganizationID  uuid.UUID
	IsSetupComplete bool
}

func (OrganizationSetupStatus) TableName() string {
	return "core.organization_setup_status"
}
