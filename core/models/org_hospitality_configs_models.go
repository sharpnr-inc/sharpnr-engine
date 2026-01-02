package models

import (
	"time"

	"github.com/google/uuid"
)

// Hospitality is Considered Sector and Each Sector have Modules
// OrganizationHospitalityModules will record the Modules we Create like PUB, CLUB, BAR, RESTAURANT, HOTEL

type OrganizationHospitalityConfiguration struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID `gorm:"uniqueIndex"`
}

type OrganizationHospitalityModules struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID

	Module              string
	HospitalityModuleID uuid.UUID
	IsEnabled           bool
	CreatedAt           time.Time
}
