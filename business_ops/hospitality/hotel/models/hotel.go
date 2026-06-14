package hotel_models

import (
	"time"

	"github.com/google/uuid"
)

type Hotel struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;not null"`
	// this is OrganizationProperty
	PropertyID    uuid.UUID `gorm:"type:uuid;index;not null"`
	LegalEntityID uuid.UUID `gorm:"type:uuid;index;not null"`

	Name string `gorm:"size:100"`

	IsActive     bool `gorm:"default:false"`
	IsConfigured bool `gorm:"default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Hotel) TableName() string {
	return "hospitality.hotel"
}

type HotelAmenities struct {
	ID      uuid.UUID
	HotelID uuid.UUID
	// this is OrganizationProperty
	PropertyID  uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (HotelAmenities) TableName() string {
	return "hospitality.hotel_amenities"
}

type HotelRoomAddons struct {
	ID                     uuid.UUID `gorm:"type:uuid;primaryKey"`
	HotelID                uuid.UUID
	OrganizationPropertyID uuid.UUID
	Name                   string
	Description            string
	Code                   string
	IsComplimentary        bool
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

func (HotelRoomAddons) TableName() string {
	return "hospitality.hotel_room_addons"
}

// HotelRoom Feature are Room Feature they are not billable they are feature
// on the other Hotel Room Addons Are billable
//✔ scoped per hotel
//✔ reusable

type HotelRoomFeature struct {
	ID                     uuid.UUID `gorm:"type:uuid;primaryKey"`
	HotelID                uuid.UUID
	OrganizationPropertyID uuid.UUID

	Code string `gorm:"size:50;uniqueIndex"`
	// AC, WIFI, MINI_BAR, LOUNGE_ACCESS, POOL_ACCESS, SPA_ACCESS

	Name        string
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (HotelRoomFeature) TableName() string {
	return "hospitality.hotel_room_feature"
}

type HotelSetting struct {
	ID             uuid.UUID
	HotelID        uuid.UUID
	KitchenEnabled bool
}
