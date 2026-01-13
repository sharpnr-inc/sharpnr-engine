package hotel_models

import (
	"time"

	"github.com/google/uuid"
)

type Hotel struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID

	Name      string
	IsActive  bool
	CreatedAt time.Time
}

type HotelAddress struct {
	ID          uuid.UUID
	HotelID     uuid.UUID
	AddressLine string
	City        string
	State       string
	Country     string
	Pincode     string
}

type HotelCompliances struct {
	ID      uuid.UUID
	HotelID uuid.UUID
	GST     string
	CIN     string
	// Add more later
}

type HotelFeatures struct {
	ID        uuid.UUID
	HotelID   uuid.UUID
	Feature   string
	IsEnabled bool
}

type HotelConfiguration struct {
	ID      uuid.UUID
	HotelID uuid.UUID
}

// Floors

type HotelFloors struct {
	ID      uuid.UUID
	HotelID uuid.UUID
}
