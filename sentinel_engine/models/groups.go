package sentinelengine_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Group is a named entity users belong to; groups can nest via ParentGroupID.
type Group struct {
	ID             uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	OrganizationID uuid.UUID  `gorm:"type:uuid;not null;uniqueIndex:idx_group_org_slug"`
	Slug           string     `gorm:"type:text;not null;uniqueIndex:idx_group_org_slug"`
	Name           string     `gorm:"type:text;not null"`
	ParentGroupID  *uuid.UUID `gorm:"type:uuid;index"` // nullable: top-level groups have no parent
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (Group) TableName() string { return "sentinel_engine.groups" }
