package bar_models

import (
	"time"

	"github.com/google/uuid"
)

type BarCharge struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID
	PropertyID     uuid.UUID
	BarID          uuid.UUID
	LegalEntityID  uuid.UUID

	SessionID *uuid.UUID

	ChargeType string
	// DRINK, FOOD, COVER, SERVICE

	SourceRefID *uuid.UUID
	// KOT_ID

	Description string

	Quantity float64
	UnitRate float64
	Amount   float64

	ChargedAt time.Time
}

func (BarCharge) TableName() string {
	return "hospitality.bar_charges"
}

type BarPayment struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;index;not null"`
	PropertyID     uuid.UUID `gorm:"type:uuid;index;not null"`
	BarID          uuid.UUID `gorm:"type:uuid;index;not null"`
	LegalEntityID  uuid.UUID `gorm:"type:uuid;index;not null"`

	SessionID uuid.UUID `gorm:"type:uuid;index;not null"`
	// BarTableSessionID

	PaymentMode string `gorm:"size:30;not null"`
	// CASH, CARD, UPI, WALLET, ROOM_CHARGE

	Amount float64 `gorm:"not null"`

	Reference string `gorm:"size:100"`
	// UPI ref / Card ref / Receipt no

	Status string `gorm:"size:20;default:'SUCCESS'"`
	// SUCCESS, FAILED, REFUNDED

	PaidAt    time.Time
	CreatedAt time.Time
}

func (BarPayment) TableName() string {
	return "hospitality.bar_payment"
}
