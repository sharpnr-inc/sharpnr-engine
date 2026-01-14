package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type OrganizationSetting struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	// Optional scoping
	Domain OrgDomainType `gorm:"type:domain_code;index"`
	Module string        `gorm:"size:50;index"`

	Scope SettingScope `gorm:"size:20;not null"`

	// Only used when Scope = UNIT
	UnitID *uuid.UUID `gorm:"type:uuid;index"`

	Settings datatypes.JSON `gorm:"type:jsonb;not null;default:'{}'"`

	Version   int
	UpdatedAt time.Time
}
