package org_setup_wizard_contracts

import (
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
	ID          uuid.UUID  `json:"id"`
	Step        string     `json:"step"`
	Status      string     `json:"status"`
	Version     int        `json:"version"`
	StartedAt   time.Time  `json:"started_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}
