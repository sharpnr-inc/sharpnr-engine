package hospitalitymodels

import (
	"time"

	"sharpnr-engine/core/models"

	"github.com/google/uuid"
)

// ---------------------------------------------
// 1️⃣ BAR
// ---------------------------------------------
type Bar struct {
	ID             uuid.UUID           `gorm:"type:uuid;primaryKey" json:"id"`
	Name           string              `gorm:"size:100;not null" json:"name"`
	OrganizationID uuid.UUID           `gorm:"type:uuid;not null" json:"organization_id"`
	Organization   models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;" json:"organization"`

	OpeningTime time.Time `json:"opening_time"`
	ClosingTime time.Time `json:"closing_time"`

	// Optional: If bar has multiple floors (e.g., rooftop bar)
	Floors []BarFloor `gorm:"foreignKey:BarID;constraint:OnDelete:CASCADE;" json:"floors"`

	// Tables directly under bar (for single-floor bars)
	Tables []BarTable `gorm:"foreignKey:BarID;constraint:OnDelete:CASCADE;" json:"tables"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ---------------------------------------------
// 2️⃣ BAR FLOOR (Optional Layer)
// ---------------------------------------------
type BarFloor struct {
	ID             uuid.UUID           `gorm:"type:uuid;primaryKey" json:"id"`
	BarID          uuid.UUID           `gorm:"type:uuid;not null" json:"bar_id"` // ✅ match type
	Bar            Bar                 `gorm:"foreignKey:BarID;constraint:OnDelete:CASCADE;" json:"bar"`
	OrganizationID uuid.UUID           `gorm:"type:uuid;not null" json:"organization_id"`
	Organization   models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;" json:"organization"`

	Name        string `gorm:"size:100;not null" json:"name"` // e.g., "Ground Floor", "Rooftop"
	Description string `gorm:"size:255" json:"description"`
	FloorNumber int    `gorm:"default:0" json:"floor_number"` // For sorting and UI display order

	Tables []BarTable `gorm:"foreignKey:FloorID;constraint:OnDelete:CASCADE;" json:"tables"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ---------------------------------------------
// 3️⃣ BAR TABLE
// ---------------------------------------------
type BarTable struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	BarID uuid.UUID `gorm:"type:uuid;not null" json:"bar_id"`
	Bar   Bar       `gorm:"foreignKey:BarID;constraint:OnDelete:CASCADE;" json:"bar"`

	// Optional Floor Relationship
	FloorID *uuid.UUID `gorm:"type:uuid" json:"floor_id"`
	Floor   *BarFloor  `gorm:"foreignKey:FloorID;constraint:OnDelete:SET NULL;" json:"floor"`

	Name     string `gorm:"size:50;not null" json:"name"` // e.g., "B1", "VIP Table 1"
	Capacity int    `gorm:"not null;default:4" json:"capacity"`
	Status   string `gorm:"size:20;default:'AVAILABLE'" json:"status"` // AVAILABLE, OCCUPIED, RESERVED

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
