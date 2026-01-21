package org_property_models

import (
	"time"

	"github.com/google/uuid"
)

type PropertyServicePoint struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;index;not null"`
	PropertyID     uuid.UUID `gorm:"type:uuid;index;not null"`

	Code string `gorm:"size:50;not null"`
	// T1, BAR_COUNTER, ROOM_402, TAKEAWAY, ONLINE

	Name        string `gorm:"size:100"`
	Description string

	Domain string
	// HOSPITALITY (future: MEDICAL, RETAIL)

	Module string
	// HOTEL, RESTAURANT, BAR, CLOUD_KITCHEN

	ModuleRefID *uuid.UUID
	// hotel_id, restaurant_id, bar_id (optional but useful)

	ServicePointType string
	// TABLE, ROOM, COUNTER, ONLINE, TAKEAWAY

	IsActive bool `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServicePointKitchenMapping struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	PropertyServicePointID uuid.UUID `gorm:"type:uuid;index;not null"`
	KitchenID              uuid.UUID `gorm:"type:uuid;index;not null"`

	// Optional routing rules
	Priority int `gorm:"default:1"`
	// 1 = primary kitchen, 2 = fallback

	IsActive bool `gorm:"default:true"`

	CreatedAt time.Time
}
