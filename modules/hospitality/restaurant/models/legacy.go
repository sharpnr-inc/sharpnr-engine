package restaurant_models

import (
	restaurant_constants "sharpnr-engine/modules/hospitality/restaurant/constants"
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
	IsActive       bool      `gorm:"default:false"`
	IsConfigured   bool      `gorm:"default:false"`

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

type RestaurantDiningSession struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	RestaurantID uuid.UUID `gorm:"type:uuid;not null;index"`
	TableID      uuid.UUID `gorm:"type:uuid;not null;index"`

	Status restaurant_constants.RestaurantDiningSessionStatusType `gorm:"size:20;not null"`
	// OPEN, CLOSED, CANCELLED

	StartedAt time.Time
	ClosedAt  *time.Time

	CreatedAt time.Time
}

func (RestaurantDiningSession) TableName() string {
	return "hospitality.restaurant_dining_session"
}

type RestaurantTakawaysSales struct {
	ID           uuid.UUID
	RestaurantID uuid.UUID
}
