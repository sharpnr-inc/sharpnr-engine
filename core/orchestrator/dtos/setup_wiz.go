package orchestrator_dtos

import (
	"time"

	"github.com/google/uuid"
)

type SetupOrganizationHQAddressReqDTO struct {
	StepID      uuid.UUID `json:"step_id" binding:"required"`
	AddressLine string    `json:"address_line" binding:"required"`
	City        string    `json:"city" binding:"required"`
	State       string    `json:"state" binding:"required"`
	Country     string    `json:"country" binding:"required"`
	Pincode     string    `json:"pincode" binding:"required"`
	Latitude    *float64  `json:"latitude,omitempty"`
	Longitude   *float64  `json:"longitude,omitempty"`
	AccuracyM   *int      `json:"accuracy_m,omitempty"`
}

type SetupOrgPropertyReqDTO struct {
	Name        string   `json:"name"`
	AddressLine string   `json:"address_line"`
	City        string   `json:"city"`
	State       string   `json:"state"`
	Country     string   `json:"country"`
	Pincode     string   `json:"pincode"`
	Longitude   *float64 `json:"longitude,omitempty"`
	Latitude    *float64 `json:"latitude,omitempty"`
	AccuracyM   *int     `json:"accuracy_m,omitempty"`
}

type SetupLegalEntityReqDTO struct {
	StepID      uuid.UUID `json:"step_id" binding:"required"`
	Type        string    `json:"type"`
	Name        string    `json:"name"`
	AddressLine string    `json:"address_line"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Country     string    `json:"country"`
	Pincode     string    `json:"pincode"`
}

type SetupProgressResponseDTO struct {
	OverallStatus string                         `json:"overall_status"`
	CurrentStep   *string                        `json:"current_step,omitempty"`
	ProgressPct   int                            `json:"progress_pct"`
	IsCompleted   bool                           `json:"is_completed"`
	Steps         []SetupStepProgressResponseDTO `json:"steps"`
}

type SetupStepProgressResponseDTO struct {
	ID          uuid.UUID  `json:"id"`
	Step        string     `json:"step"`
	Status      string     `json:"status"`
	Version     int        `json:"version"`
	StartedAt   time.Time  `json:"started_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}
