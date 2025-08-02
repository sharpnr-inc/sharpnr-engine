package models

import (
	"time"

	"sharpnr-engine/core/constants"

	"github.com/google/uuid"
)

type Operations struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`

	Name string `gorm:"type:varchar(255);not null"`
	Slug string `gorm:"type:varchar(255);not null;uniqueIndex:idx_org_slug"`

	Description string `gorm:"type:text"`

	OrganizationID uuid.UUID    `gorm:"type:uuid;not null;uniqueIndex:idx_org_slug"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	DepartmentID uuid.UUID  `gorm:"type:uuid;not null"`
	Department   Department `gorm:"foreignKey:DepartmentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	ModuleType constants.ModuleType `gorm:"type:varchar(50);not null"`

	IsDefault bool `gorm:"default:false;not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (Operations) TableName() string {
	return "operations"
}
