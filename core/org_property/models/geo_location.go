package org_property_models

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationPropertyGeoLocation struct {
	ID                     uuid.UUID
	OrganizationPropertyID uuid.UUID

	Latitude  float64 `gorm:"type:decimal(9,6);not null"`
	Longitude float64 `gorm:"type:decimal(9,6);not null"`

	AccuracyM *int   `gorm:"type:int"`
	Source    string `gorm:"type:varchar(20)"`

	CapturedAt time.Time
	UpdatedAt  time.Time
}
