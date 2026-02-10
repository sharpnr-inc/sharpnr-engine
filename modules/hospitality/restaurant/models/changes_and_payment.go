package restaurant_models

import (
	"time"

	"github.com/google/uuid"
)

type RestaurantCharge struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID
	PropertyID     uuid.UUID
	RestaurantID   uuid.UUID
	LegalEntityID  uuid.UUID

	SessionID uuid.UUID `gorm:"type:uuid;not null;index"`

	ChargeType string
	// FOOD, SERVICE_CHARGE, PACKAGING

	SourceRefID *uuid.UUID
	// KOT_ID

	Description string

	Quantity float64
	UnitRate float64
	Amount   float64

	ChargedAt time.Time
	CreatedAt time.Time
}

func (RestaurantCharge) TableName() string {
	return "hospitality.restaurant_charge"
}

type RestaurantSessionPayment struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID
	PropertyID     uuid.UUID
	RestaurantID   uuid.UUID
	LegalEntityID  uuid.UUID

	SessionID uuid.UUID `gorm:"type:uuid;not null;index"`

	Amount float64

	Method string
	// CASH, UPI, CARD

	Status string
	// RECEIVED, FAILED

	Reference string

	ReceivedAt time.Time
}

func (RestaurantSessionPayment) TableName() string {
	return "hospitality.restaurant_session_payment"
}
