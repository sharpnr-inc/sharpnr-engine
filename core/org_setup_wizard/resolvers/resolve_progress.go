package org_setup_wizard_resolvers

import (
	org_setup_wizard_contracts "sharpnr-engine/core/org_setup_wizard/contracts"
	org_setup_wizard_models "sharpnr-engine/core/org_setup_wizard/models"
	"sort"
)

func ResolveSetupProgress(steps []*org_setup_wizard_models.OrganizationSetup) org_setup_wizard_contracts.SetupProgress {
	if len(steps) == 0 {
		return org_setup_wizard_contracts.SetupProgress{
			OverallStatus: "NOT_STARTED",
			ProgressPct:   0,
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
		currentStep    *string
		hasBlocked     bool
	)

	for _, s := range steps {

		step := org_setup_wizard_contracts.SetupStepProgress{
			ID:          s.ID,
			Status:      s.Status,
			Version:     s.Version,
			StartedAt:   s.StartedAt,
			CompletedAt: s.CompletedAt,
		}

		resultSteps = append(resultSteps, step)

		switch s.Status {
		case "COMPLETED":
			completedCount++
		case "IN_PROGRESS":
			if currentStep == nil {
			}
		}
	}

	total := len(steps)

	var overallStatus string
	var isCompleted bool

	switch {
	case hasBlocked:
		overallStatus = "BLOCKED"
	case completedCount == 0:
		overallStatus = "NOT_STARTED"
	case completedCount == total:
		overallStatus = "COMPLETED"
		isCompleted = true
	default:
		overallStatus = "IN_PROGRESS"
	}

	progressPct := int(float64(completedCount) / float64(total) * 100)

	return org_setup_wizard_contracts.SetupProgress{
		OverallStatus: overallStatus,
		CurrentStep:   currentStep,
		ProgressPct:   progressPct,
		IsCompleted:   isCompleted,
		Steps:         resultSteps,
	}
}
