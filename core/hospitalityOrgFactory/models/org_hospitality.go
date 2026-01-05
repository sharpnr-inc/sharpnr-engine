package hospitality_org_models

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationHospitalityModules struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID

	Module              string
	HospitalityModuleID uuid.UUID
	IsEnabled           bool
	CreatedAt           time.Time
}

type OrganizationHospitalityConfiguration struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID `gorm:"uniqueIndex"`

	Timezone string
	Currency string

	CreatedAt time.Time
}

type OrganizationHospitality struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;uniqueIndex"`

	Status string // CREATED, CONFIGURING, ACTIVE

	CreatedAt time.Time
	UpdatedAt time.Time
}

type HospitalityModule struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	Code string `gorm:"uniqueIndex"` // BAR, RESTAURANT, HOTEL
	Name string

	CreatedAt time.Time
}
