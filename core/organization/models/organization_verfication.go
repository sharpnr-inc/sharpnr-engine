package models

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationVerification struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;uniqueIndex"`

	EmailVerified    bool `gorm:"default:false"`
	MobileVerified   bool `gorm:"default:false"`
	PhysicalVerified bool `gorm:"default:false"`
	GSTVerified      bool `gorm:"default:false"`
	CINVerified      bool `gorm:"default:false"`

	VerifiedAt *time.Time
	VerifiedBy *uuid.UUID // admin / system

	CreatedAt time.Time
	UpdatedAt time.Time
}
