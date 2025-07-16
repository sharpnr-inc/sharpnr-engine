package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/sharpnr-inc/sharpnr-engine/core/constants"
)

type Operations struct {
	ID             uuid.UUID            `gorm:"type:uuid;primaryKey"`
	Name           string               `gorm:"type:varchar(255); not null"`
	Slug           string               `gorm:"type:varchar(255); uniqueIndex:idx_org_slug;not null"`
	Description    string               `gorm:"type:text"`
	OrganizationID uuid.UUID            `gorm:"type:uuid;uniqueIndex:idx_org_slug;not null"`
	DepartmentID   uuid.UUID            `gorm:"type:uuid;noy null"`
	ModuleType     constants.ModuleType `gorm:"type:varchar(50);not null"`
	IsDefault      bool                 `gorm:"default:false"`
	CreatedAt      time.Time            `gorm:"autoCreateTime"`
	UpdatedAt      time.Time            `gorm:"autoCreateTime"`
}

func (Operations) TableName() string {
	return "operations"
}
