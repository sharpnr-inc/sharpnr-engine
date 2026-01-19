package org_setup_wizard_contracts

import "github.com/google/uuid"

type SetupProgressOutput struct {
	StepID         uuid.UUID `json:"id"`
	Status         string    `json:"status"`
	CurrentStep    string    `json:"current_step,omitempty"`
	CompletedSteps []string  `json:"completed_steps"`
	TotalSteps     int       `json:"total_steps"`
	CompletedCount int       `json:"completed_count"`
}
