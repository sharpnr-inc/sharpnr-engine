package org_setup_wizards_constants

type SetupStatus string

const (
	SetupInProgress SetupStatus = "IN_PROGRESS"
	SetupCompleted  SetupStatus = "COMPLETED"
)

type SetupStepType string

const (
	SetupStepSuperAdminProfile SetupStepType = "SUPER_ADMIN_PROFILE"
	SetupStepOperations                      = "OPERATIONS_SETUP"
	SetupStepDone              SetupStepType = "DONE"
)
