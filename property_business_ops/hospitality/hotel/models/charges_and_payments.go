package hotel_models

import (
	"time"

	"github.com/google/uuid"
)

type HotelCharge struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`
	PropertyID     uuid.UUID `gorm:"type:uuid;not null;index"`
	HotelID        uuid.UUID `gorm:"type:uuid;not null;index"`
	LegalEntityID  uuid.UUID `gorm:"type:uuid;not null;index"`

	// The specific stay this charge belongs to
	StayID uuid.UUID `gorm:"type:uuid;not null;index"`

	// ROOM_RENT, FOOD, LAUNDRY, SPA, MINIBAR
	ChargeType string `gorm:"type:varchar(50);index"`

	// Source: HOTEL, RESTAURANT, KITCHEN
	SourceModule string `gorm:"type:varchar(50)"`

	// KOT_ID, SERVICE_ID, etc. (Links to the original order)
	SourceRefID *uuid.UUID `gorm:"type:uuid;index"`

	Description string `gorm:"type:text"`

	Quantity float64 `gorm:"type:decimal(10,2);default:1"`
	UnitRate float64 `gorm:"type:decimal(12,2)"`
	Amount   float64 `gorm:"type:decimal(12,2)"` // Usually Qty * UnitRate

	ChargeDate time.Time `gorm:"index"`
	CreatedAt  time.Time
}

func (HotelCharge) TableName() string {
	return "hospitality.hotel_charge"
}

type HotelStayPayment struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`
	PropertyID     uuid.UUID `gorm:"type:uuid;not null;index"`
	HotelID        uuid.UUID `gorm:"type:uuid;not null;index"`
	LegalEntityID  uuid.UUID `gorm:"type:uuid;not null;index"`

	// Links to the specific guest stay
	StayID uuid.UUID `gorm:"type:uuid;not null;index"`

	Amount float64 `gorm:"type:decimal(12,2);not null"`

	// Method: CASH, UPI, CARD, BANK_TRANSFER
	Method string `gorm:"type:varchar(20);not null"`

	// Status: RECEIVED, FAILED, REFUNDED
	Status string `gorm:"type:varchar(20);not null;index;default:'RECEIVED'"`

	// Transaction reference (UPI ID, Card Auth Code, Receipt No)
	Reference string `gorm:"type:varchar(100)"`

	ReceivedAt time.Time `gorm:"index"`
	CreatedAt  time.Time
}

func (HotelStayPayment) TableName() string {
	return "hospitality.hotel_stay_payment"
}
