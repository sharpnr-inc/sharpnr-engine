package hotel_models

import (
	hotel_constants "sharpnr-engine/modules/hospitality/hotel/constants"
	"time"

	"github.com/google/uuid"
)

type HotelFloor struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	HotelID     uuid.UUID `gorm:"type:uuid;not null"`
	Name        string
	FloorNumber int `gorm:"not null"`
	CreatedAt   time.Time
}

func (HotelFloor) TableName() string {
	return "hospitality.hotel_floor"
}

type HotelRoom struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	HotelID      uuid.UUID `gorm:"type:uuid;not null;index"`
	HotelFloorID uuid.UUID `gorm:"type:uuid;not null;index"`
	// This can be null
	HotelRoomTypeID *uuid.UUID

	RoomCode string `gorm:"size:30;"`
	Name     string `gorm:"size:100"`

	Status hotel_constants.HotelRoomStatusType `gorm:"size:20;not null;default:'AVAILABLE'"`
	// AVAILABLE, OCCUPIED, RESERVED, MAINTENANCE, BLOCKED

	BedType   string
	MaxGuests int

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (HotelRoom) TableName() string {
	return "hospitality.hotel_room"
}

type HotelRoomDetail struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	HotelRoomID     uuid.UUID `gorm:"type:uuid;uniqueIndex"`
	HotelRoomTypeID *uuid.UUID

	Description string `gorm:"type:text"`
	BedType     string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (HotelRoomDetail) TableName() string {
	return "hospitality.hotel_room_detail"
}

// This is table is basically to override the Category price or type Price
type HotelRoomRate struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	HotelID     uuid.UUID `gorm:"type:uuid;not null;index"`
	HotelRoomID uuid.UUID `gorm:"type:uuid;not null;index"`

	BasePrice float64 `gorm:"not null"`

	EffectiveFrom time.Time
	EffectiveTo   *time.Time

	CreatedAt time.Time
}

func (HotelRoomRate) TableName() string {
	return "hospitality.hotel_room_rate"
}
