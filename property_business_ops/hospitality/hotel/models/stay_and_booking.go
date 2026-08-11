package hotel_models

import (
	"time"

	"github.com/google/uuid"
)

type HotelStay struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	HotelID uuid.UUID `gorm:"type:uuid;not null;index"`
	RoomID  uuid.UUID `gorm:"type:uuid;not null;index"`

	PrimaryGuestID uuid.UUID `gorm:"type:uuid;not null"`

	Status string `gorm:"size:20;not null"`
	// RESERVED, CHECKED_IN, CHECKED_OUT, CANCELLED

	CheckInDate  time.Time
	CheckOutDate time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (HotelStay) TableName() string {
	return "hospitality.hotel_stay"
}
