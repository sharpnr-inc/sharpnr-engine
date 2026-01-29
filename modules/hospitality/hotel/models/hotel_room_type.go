package hotel_models

import (
	"time"

	"github.com/google/uuid"
)

// Hotel Room Type Can called as Category
// this belongs to all together to To hotel

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

func (HotelRoomType) TableName() string {
	return "hospitality.hotel_room_type"
}

type HotelRoomTypeFeature struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey"`
	HotelID            uuid.UUID
	HotelRoomFeatureID uuid.UUID
	HotelRoomTypeID    uuid.UUID
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (HotelRoomTypeFeature) TableName() string {
	return "hospitality.hotel_room_type_feature"
}

type HotelRoomTypeAddon struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey"`
	HotelID           uuid.UUID
	HotelRoomAddonsID uuid.UUID
	HotelRoomTypeID   uuid.UUID
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (HotelRoomTypeAddon) TableName() string {
	return "hospitality.hotel_room_type_addon"
}

type HotelRoomTypeRate struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	HotelID    uuid.UUID `gorm:"type:uuid;not null;index"`
	RoomTypeID uuid.UUID `gorm:"type:uuid;not null;index"`

	BasePrice float64 `gorm:"not null"`

	EffectiveFrom time.Time `gorm:"not null"`
	EffectiveTo   *time.Time

	CreatedAt time.Time
}

func (HotelRoomTypeRate) TableName() string {
	return "hospitality.hotel_room_type_rate"
}
