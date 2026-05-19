package org_setup_wizards_constants

type SetupStatus string

const (
	SetupStatusInProgress SetupStatus = "IN_PROGRESS"
	SetupStatusCompleted  SetupStatus = "COMPLETED"
)

type SetupStepType string

const (
	SetupStepSuperAdminProfile SetupStepType = "SUPER_ADMIN_PROFILE"
	SetupStepOrgDetailsUpdate  SetupStepType = "ORGANIZATION_DETAILS_UPDATES"
	SetupStepOperations        SetupStepType = "OPERATIONS_SETUP"
)

type OverallSetupStatus string

const (
	OverallStatusNotStarted OverallSetupStatus = "NOT_STARTED"
	OverallStatusInProgress OverallSetupStatus = "IN_PROGRESS"
	OverallStatusCompleted  OverallSetupStatus = "COMPLETED"
)

type StepDefinition struct {
	Step  SetupStepType
	Order int
}

var SetupSteps = []StepDefinition{
	{Step: SetupStepSuperAdminProfile, Order: 1},
	{Step: SetupStepOrgDetailsUpdate, Order: 2},
	{Step: SetupStepOperations, Order: 3},
}

func GetNextStep(current SetupStepType) (SetupStepType, bool) {
	for i, step := range SetupSteps {
		if step.Step == current {
			if i+1 < len(SetupSteps) {
				return SetupSteps[i+1].Step, true
			}
			return "", false // no next step
		}
	}
	return "", false
}

func IsLastStep(step SetupStepType) bool {
	if len(SetupSteps) == 0 {
		return false
	}
	return SetupSteps[len(SetupSteps)-1].Step == step
}

func IsValidStep(step SetupStepType) bool {
	for _, s := range SetupSteps {
		if s.Step == step {
			return true
		}
	}
	return false
}
