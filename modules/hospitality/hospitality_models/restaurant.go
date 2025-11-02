package hospitalitymodels

import (
	"time"

	"sharpnr-engine/core/models"
)

type Restaurant struct {
	ID             string              `gorm:"type:uuid;primaryKey" json:"id"`
	Name           string              `gorm:"size:100;not null" json:"name"`
	OrganizationID string              `gorm:"type:uuid;not null" json:"organization_id"`
	Organization   models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;" json:"organization"`

	IsPureVeg   bool      `gorm:"default:false" json:"is_pure_veg"`
	OpeningTime time.Time `json:"opening_time"`
	ClosingTime time.Time `json:"closing_time"`

	Tables []RestaurantTable `gorm:"foreignKey:RestaurantID;constraint:OnDelete:CASCADE;" json:"tables"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// RestaurantTable represents an individual table in a restaurant
type RestaurantTable struct {
	ID           string     `gorm:"type:uuid;primaryKey" json:"id"`
	RestaurantID string     `gorm:"type:uuid;not null" json:"restaurant_id"`
	Restaurant   Restaurant `gorm:"foreignKey:RestaurantID;constraint:OnDelete:CASCADE;" json:"restaurant"`

	Name     string `gorm:"size:50;not null" json:"name"`
	Capacity int    `gorm:"not null;default:4" json:"capacity"`
	Status   string `gorm:"size:20;default:'AVAILABLE'" json:"status"` // e.g. AVAILABLE, OCCUPIED, RESERVED

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
