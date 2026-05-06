package org_setup_wizard_contracts

import (
	org_setup_wizards_constants "sharpnr-engine/core/org_setup_wizard/constants"
	"time"

	"github.com/google/uuid"
)

type SetupProgress struct {
	OverallStatus string              `json:"overall_status"`
	CurrentStep   *string             `json:"current_step,omitempty"`
	ProgressPct   int                 `json:"progress_pct"`
	IsCompleted   bool                `json:"is_completed"`
	Steps         []SetupStepProgress `json:"steps"`
}

type SetupStepProgress struct {
	ID          uuid.UUID                                 `json:"id"`
	Step        org_setup_wizards_constants.SetupStepType `json:"step"`
	Status      org_setup_wizards_constants.SetupStatus   `json:"status"`
	Version     int                                       `json:"version"`
	StartedAt   time.Time                                 `json:"started_at"`
	CompletedAt *time.Time                                `json:"completed_at,omitempty"`
}
