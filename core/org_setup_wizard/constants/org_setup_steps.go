package org_setup_wizards_constants

// org_setup_wizard_constants

type SetupStatus string

const (
	SetupNotStarted SetupStatus = "NOT_STARTED"
	SetupInProgress SetupStatus = "IN_PROGRESS"
	SetupCompleted  SetupStatus = "COMPLETED"
	SetupBlocked    SetupStatus = "BLOCKED"
)

type SetupStep string

const (
	StepOrgProfile    SetupStep = "ORG_PROFILE"
	StepDomainSelect  SetupStep = "DOMAIN_SELECT"
	StepPropertySetup SetupStep = "PROPERTY_SETUP"
	StepLegalEntity   SetupStep = "LEGAL_ENTITY"
	StepAccounting    SetupStep = "ACCOUNTING"
	StepHRInit        SetupStep = "HR_INIT"
	StepDone          SetupStep = "DONE"
)
