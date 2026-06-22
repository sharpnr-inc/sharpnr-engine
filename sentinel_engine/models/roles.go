package sentinelengine_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Role is a named entity that policies reference (principal in Role::"<slug>").
// Slug is the Cedar entity id; it must be stable because policies embed it as text.
type Role struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_role_org_slug"`
	Slug           string    `gorm:"type:text;not null;uniqueIndex:idx_role_org_slug"` // e.g. "editor"
	Name           string    `gorm:"type:text;not null"`                               // human label
	Description    string    `gorm:"type:text"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (Role) TableName() string { return "sentinel_engine.roles" }
