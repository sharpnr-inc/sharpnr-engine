package org_setup_wizard_resolvers

import (
	org_setup_wizards_constants "sharpnr-engine/core/org_setup_wizard/constants"
	org_setup_wizard_contracts "sharpnr-engine/core/org_setup_wizard/contracts"
	org_setup_wizard_models "sharpnr-engine/core/org_setup_wizard/models"
	"sort"
)

// Resolves Steps and return the How much user complete the remain on setup wiz
// it takes the slice organization setup and resolves it
func ResolveSetupProgress(steps []*org_setup_wizard_models.OrganizationSetup) org_setup_wizard_contracts.SetupProgress {

	if len(steps) == 0 {
		return org_setup_wizard_contracts.SetupProgress{
			OverallStatus: org_setup_wizards_constants.OverallStatusNotStarted,
			IsCompleted:   false,
			Steps:         []org_setup_wizard_contracts.SetupStepProgress{},
		}
	}

	// Deterministic ordering
	sort.Slice(steps, func(i, j int) bool {
		if steps[i].Version != steps[j].Version {
			return steps[i].Version < steps[j].Version
		}
		return steps[i].StartedAt.Before(steps[j].StartedAt)
	})

	resultSteps := make([]org_setup_wizard_contracts.SetupStepProgress, 0, len(steps))

	var (
		completedCount int
		currentStep    *org_setup_wizards_constants.SetupStepType
	)

	for _, s := range steps {

		step := org_setup_wizard_contracts.SetupStepProgress{
			ID:          s.ID,
			Step:        s.CurrentStep,
			Status:      s.Status,
			Version:     s.Version,
			StartedAt:   s.StartedAt,
			CompletedAt: s.CompletedAt,
		}

		resultSteps = append(resultSteps, step)

		switch s.Status {
		// A skipped step counts towards completion: the user opted out of it
		// rather than leaving it pending.
		case org_setup_wizards_constants.SetupStatusCompleted,
			org_setup_wizards_constants.SetupStatusSkipped:
			completedCount++
		case org_setup_wizards_constants.SetupStatusInProgress:
			if currentStep == nil {
				currentStep = &s.CurrentStep
			}
		}
	}

	total := len(org_setup_wizards_constants.SetupSteps)

	var overallStatus org_setup_wizards_constants.OverallSetupStatus
	var isCompleted bool

	switch {
	case completedCount == 0:
		overallStatus = org_setup_wizards_constants.OverallStatusNotStarted
	case completedCount == total:
		overallStatus = org_setup_wizards_constants.OverallStatusCompleted
		isCompleted = true
	default:
		overallStatus = org_setup_wizards_constants.OverallStatusInProgress
	}

	return org_setup_wizard_contracts.SetupProgress{
		OverallStatus: overallStatus,
		CurrentStep:   currentStep,
		IsCompleted:   isCompleted,
		Steps:         resultSteps,
	}
}
