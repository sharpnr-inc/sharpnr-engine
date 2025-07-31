package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Terminal struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`

	OrganizationID uuid.UUID    `gorm:"type:uuid;not null;index"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Name        string `gorm:"size:100;not null"`
	Description string `gorm:"type:text"`

	// Many-to-many with employees
	Employees []Employee `gorm:"many2many:terminal_employees;"`

	// Terminal ↔ Operations (many-to-many)
	Operations []Operations `gorm:"many2many:terminal_operations;"`

	IsActive  bool `gorm:"default:true;not null"`
	IsDefault bool `gorm:"default:false;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
