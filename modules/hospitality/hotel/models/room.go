package hotel_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type HotelFloor struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	HotelID uuid.UUID `gorm:"type:uuid;not null"`

	Name        string
	FloorNumber int `gorm:"not null"`
	CreatedAt   time.Time
}

type HotelRoom struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	HotelID      uuid.UUID `gorm:"type:uuid;not null;index"`
	HotelFloorID uuid.UUID `gorm:"type:uuid;not null;index"`

	RoomCode string `gorm:"size:30;not null;uniqueIndex"`
	Name     string `gorm:"size:100"`

	Status string `gorm:"size:20;not null;default:'AVAILABLE'"`
	// AVAILABLE, OCCUPIED, RESERVED, MAINTENANCE, BLOCKED

	BedType   string
	MaxGuests int

	CreatedAt time.Time
	UpdatedAt time.Time
}

type HotelRoomDetail struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	HotelRoomID uuid.UUID `gorm:"type:uuid;uniqueIndex"`

	Description string `gorm:"type:text"`

	HasBalcony bool
	HasPool    bool
	HasLounge  bool
	HasSpa     bool

	AmenitiesJSON datatypes.JSON
	// AC, WIFI, TV, SAFE, MINI_BAR

	CreatedAt time.Time
	UpdatedAt time.Time
}

type HotelRoomRate struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	HotelID     uuid.UUID `gorm:"type:uuid;not null;index"`
	HotelRoomID uuid.UUID `gorm:"type:uuid;not null;index"`

	BasePrice float64 `gorm:"not null"`

	EffectiveFrom time.Time
	EffectiveTo   *time.Time

	CreatedAt time.Time
}

// !! ----  These tables are not Migrated yet For Hotel Version 2

type HotelRoomFeature struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	Code string `gorm:"size:50;uniqueIndex"`
	// AC, WIFI, MINI_BAR, LOUNGE_ACCESS, POOL_ACCESS, SPA_ACCESS

	Name        string
	Description string

	IsBillable bool
	// false → complimentary
	// true  → extra charge

	CreatedAt time.Time
}

type HotelRoomType struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	HotelID uuid.UUID `gorm:"type:uuid;index;not null"`

	Code string `gorm:"size:50;not null"` // DELUXE, SUITE
	Name string `gorm:"size:100"`

	Description string `gorm:"type:text"`

	BaseOccupancy int
	MaxOccupancy  int

	IsActive bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type HotelRoomTypeFeature struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	RoomTypeID uuid.UUID `gorm:"type:uuid;index;not null"`
	FeatureID  uuid.UUID `gorm:"type:uuid;index;not null"`

	IsIncluded bool
	// true  → included in base price
	// false → optional

	ExtraCharge *float64
	// only if IsIncluded = false

	CreatedAt time.Time
}
