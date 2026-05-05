package org_setup_wizard_models

import (
	org_setup_wizards_constants "sharpnr-engine/core/org_setup_wizard/constants"
	"time"

	"github.com/google/uuid"
)

// A organization can have multiple setup steps, each step will be recorded in the database with its status and timestamp. This //
// allows us to track the progress of the setup wizard for each organization and provide insights into which steps are completed,
// in progress, or pending.

// Relation :--> OrganizationSetup Table 1:N Organization Table

type OrganizationSetup struct {
	ID             uuid.UUID                                 `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID                                 `gorm:"type:uuid;index;not null"`
	Status         org_setup_wizards_constants.SetupStatus   `gorm:"type:varchar(20);not null"`
	CurrentStep    org_setup_wizards_constants.SetupStepType `gorm:"type:varchar(50);not null"`
	Version        int                                       `gorm:"default:1"`
	StartedAt      time.Time
	CompletedAt    *time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (OrganizationSetup) TableName() string {
	return "setup_wizard.organization_setup"
}
