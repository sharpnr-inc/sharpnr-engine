package restaurant_models

import (
	restaurant_constants "sharpnr-engine/property_business_ops/hospitality/restaurant/constants"
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

	PropertyID uuid.UUID `gorm:"type:uuid;not null;index"`

	// Legal Entity Can be Null for the Restaurant
	// It is used for the Getting the Tax Details and Tax Codes for the item Againts the Price
	LegalEntityID *uuid.UUID `gorm:"type:uuid;index"`

	IsPureVeg   bool   `gorm:"default:false" json:"is_pure_veg"`
	OpeningTime string `gorm:"type:varchar(5);not null"` // "06:33"
	ClosingTime string `gorm:"type:varchar(5);not null"` // "15:01"
	IsOvernight bool   `gorm:"default:false;not null"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Restaurant) TableName() string {
	return "hospitality.restaurant"
}

/* ----------------------------- ################### ----------------------------- */

type RestaurantFloor struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	RestaurantID uuid.UUID `gorm:"type:uuid;not null;index"`
	Name         string
	FloorNumber  int
	CreatedAt    time.Time
}

func (RestaurantFloor) TableName() string {
	return "hospitality.restaurant_floor"
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

	Status restaurant_constants.RestaurantTableStatusType `gorm:"size:20;default:'AVAILABLE'"`
	// AVAILABLE, OCCUPIED, RESERVED, BLOCKED

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RestaurantTable) TableName() string {
	return "hospitality.restaurant_table"
}
