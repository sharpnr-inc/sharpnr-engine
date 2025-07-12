package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Terminal struct {
	ID             uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	OrganizationID uuid.UUID  `gorm:"type:uuid;not null;index"`
	DepartmentID   *uuid.UUID `gorm:"type:uuid;index"` // Nullable for master terminal

	Name        string `gorm:"size:100;not null"`
	Type        string `gorm:"size:50;not null;index"` // master, accounting, hr, reception, housekeeping, etc.
	Description string `gorm:"type:text"`

	Features    pq.StringArray `gorm:"type:text[]"`
	Permissions pq.StringArray `gorm:"type:text[]"`

	IsActive  bool `gorm:"default:true"`
	IsDefault bool `gorm:"default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relationships
	Organization Organization `gorm:"foreignKey:OrganizationID"`
	Department   *Department  `gorm:"foreignKey:DepartmentID"`
}
