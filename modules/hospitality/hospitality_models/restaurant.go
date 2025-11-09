package hospitalitymodels

import (
	"time"

	"sharpnr-engine/core/models"

	"github.com/google/uuid"
)

// ---------------------------
// 1️⃣ RESTAURANT
// ---------------------------

type Restaurant struct {
	ID             uuid.UUID           `gorm:"type:uuid;primaryKey" json:"id"`
	Name           string              `gorm:"size:100;not null" json:"name"`
	OrganizationID uuid.UUID           `gorm:"type:uuid;not null" json:"organization_id"`
	Organization   models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;" json:"organization"`

	IsPureVeg   bool      `gorm:"default:false" json:"is_pure_veg"`
	OpeningTime time.Time `json:"opening_time"`
	ClosingTime time.Time `json:"closing_time"`

	Floors []RestaurantFloor `gorm:"foreignKey:RestaurantID;constraint:OnDelete:CASCADE;" json:"floors"`
	Tables []RestaurantTable `gorm:"foreignKey:RestaurantID;constraint:OnDelete:CASCADE;" json:"tables"` // Optional: direct access without floors

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	IsKitchen bool       `gorm:"default:true" json:"is_kitchen"`
	KitchenID *uuid.UUID `gorm:"type:uuid" json:"kitchen_id"`
	Kitchen   *Kitchen   `gorm:"foreignKey:KitchenID;constraint:OnDelete:SET NULL;" json:"kitchen"`
}

// ---------------------------
// 2️⃣ RESTAURANT FLOOR
// ---------------------------

type RestaurantFloor struct {
	ID             uuid.UUID           `gorm:"type:uuid;primaryKey" json:"id"`
	RestaurantID   uuid.UUID           `gorm:"type:uuid;not null" json:"restaurant_id"`
	Restaurant     Restaurant          `gorm:"foreignKey:RestaurantID;constraint:OnDelete:CASCADE;" json:"restaurant"`
	OrganizationID uuid.UUID           `gorm:"type:uuid;not null" json:"organization_id"` // For quick filtering
	Organization   models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;" json:"organization"`

	Name        string `gorm:"size:100;not null" json:"name"` // e.g., "Ground Floor", "Rooftop Lounge"
	Description string `gorm:"size:255" json:"description"`
	FloorNumber int    `gorm:"default:0" json:"floor_number"` // Optional, for sorting

	Tables []RestaurantTable `gorm:"foreignKey:FloorID;constraint:OnDelete:CASCADE;" json:"tables"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ---------------------------
// 3️⃣ RESTAURANT TABLE
// ---------------------------

type RestaurantTable struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	RestaurantID uuid.UUID  `gorm:"type:uuid;not null" json:"restaurant_id"`
	Restaurant   Restaurant `gorm:"foreignKey:RestaurantID;constraint:OnDelete:CASCADE;" json:"restaurant"`

	FloorID *uuid.UUID       `gorm:"type:uuid" json:"floor_id"`
	Floor   *RestaurantFloor `gorm:"foreignKey:FloorID;constraint:OnDelete:SET NULL;" json:"floor"`

	Name     string `gorm:"size:50;not null" json:"name"`
	Capacity int    `gorm:"not null;default:4" json:"capacity"`
	Status   string `gorm:"size:20;default:'AVAILABLE'" json:"status"` // AVAILABLE, OCCUPIED, RESERVED, BLOCKED

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
