package hotel_models

import (
	"time"

	"github.com/google/uuid"
)

type HotelCharge struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID
	PropertyID     uuid.UUID
	HotelID        uuid.UUID
	LegalEntityID  uuid.UUID

	StayID uuid.UUID

	ChargeType string
	// ROOM_RENT, FOOD, LAUNDRY, SPA, MINIBAR

	SourceModule string
	// HOTEL, RESTAURANT, KITCHEN

	SourceRefID *uuid.UUID
	// KOT_ID, SERVICE_ID, etc.

	Description string

	Quantity float64
	UnitRate float64
	Amount   float64

	ChargeDate time.Time
	CreatedAt  time.Time
}

type HotelStayPayment struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID
	PropertyID     uuid.UUID
	HotelID        uuid.UUID
	LegalEntityID  uuid.UUID

	StayID uuid.UUID

	Amount float64

	Method string
	// CASH, UPI, CARD

	Status string
	// RECEIVED, FAILED

	Reference string

	ReceivedAt time.Time
}
