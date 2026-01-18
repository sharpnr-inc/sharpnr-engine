package org_setup_wizard_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type OrganizationSetup struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;uniqueIndex;not null"`

	// High-level state
	Status string
	// NOT_STARTED, IN_PROGRESS, COMPLETED, BLOCKED

	// Current step pointer
	CurrentStep string
	// ORG_PROFILE, DOMAIN_SELECT, PROPERTY_SETUP,
	// LEGAL_ENTITY, ACCOUNTING, HR_INIT, DONE

	Version int `gorm:"default:1"`

	StartedAt   time.Time
	CompletedAt *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

// !! This is not migrated yet
type OrganizationSetupStep struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;index;not null"`

	StepCode string
	// ORG_PROFILE, DOMAIN_SELECT, PROPERTY_SETUP, etc.

	Status string
	// PENDING, IN_PROGRESS, COMPLETED, SKIPPED

	// Optional metadata
	StartedAt   *time.Time
	CompletedAt *time.Time

	// Optional: who completed it
	CompletedBy *uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// !! This is not migrated yet
type OrganizationSetupDraft struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;index;not null"`

	StepCode string
	// PROPERTY_SETUP, LEGAL_ENTITY, ACCOUNTING

	Data      datatypes.JSON `gorm:"type:jsonb;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
