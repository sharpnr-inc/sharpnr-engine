package restaurant_models

import (
	"time"

	"github.com/google/uuid"
)

// ---------------------------
// 1️⃣ RESTAURANT
// ---------------------------

type Restaurant struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name string    `gorm:"size:100;not null" json:"name"`

	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`
	PropertyID     uuid.UUID `gorm:"type:uuid;not null;index"`
	LegalEntityID  uuid.UUID `gorm:"type:uuid;not null;index"`

	IsPureVeg   bool   `gorm:"default:false" json:"is_pure_veg"`
	OpeningTime string `gorm:"size:5" json:"opening_time"` // e.g., "09:00"
	ClosingTime string `gorm:"size:5" json:"closing_time"` // e.g., "22:00"

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RestaurantFloor struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	RestaurantID uuid.UUID `gorm:"type:uuid;not null;index"`

	Name        string
	FloorNumber int

	CreatedAt time.Time
}

// ---------------------------
// 3️⃣ RESTAURANT TABLE
// ---------------------------

type RestaurantTable struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	RestaurantID uuid.UUID `gorm:"type:uuid;not null;index"`
	FloorID      uuid.UUID `gorm:"type:uuid;not null;index"`

	Code string `gorm:"size:30;not null"`
	Name string `gorm:"size:100"`

	Capacity int

	Status string `gorm:"size:20;default:'AVAILABLE'"`
	// AVAILABLE, OCCUPIED, RESERVED, BLOCKED

	CreatedAt time.Time
	UpdatedAt time.Time
}

type RestaurantDiningSession struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	RestaurantID uuid.UUID `gorm:"type:uuid;not null;index"`
	TableID      uuid.UUID `gorm:"type:uuid;not null;index"`

	Status string `gorm:"size:20;not null"`
	// OPEN, CLOSED, CANCELLED

	StartedAt time.Time
	ClosedAt  *time.Time

	CreatedAt time.Time
}
