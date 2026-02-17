package restaurant_models

import (
	"time"

	"github.com/google/uuid"
)

type RestaurantCharge struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`
	PropertyID     uuid.UUID `gorm:"type:uuid;index"`
	RestaurantID   uuid.UUID `gorm:"type:uuid;not null;index"`
	LegalEntityID  uuid.UUID `gorm:"type:uuid;not null;index"`

	RestaurantDiningSessionID *uuid.UUID `gorm:"type:uuid;index"`
	RestaurantPickUpOrdersID  *uuid.UUID `gorm:"type:uuid;index"`

	RestaurantOrderID     *uuid.UUID `gorm:"type:uuid;index"`
	RestaurantOrderItemID *uuid.UUID `gorm:"type:uuid;index"`

	ChargeType string `gorm:"type:varchar(50);not null;index"`
	// FOOD, SERVICE_CHARGE, PACKAGING, TAX, DISCOUNT, ROUND_OFF

	SourceRefID *uuid.UUID
	// KOT_ID, TAX_RULE_ID, DISCOUNT_RULE_ID

	Description string `gorm:"type:text"`

	Quantity float64 `gorm:"type:numeric(10,3);not null"`
	UnitRate float64 `gorm:"type:numeric(12,2);not null"`
	Amount   float64 `gorm:"type:numeric(12,2);not null"`

	ChargedAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RestaurantCharge) TableName() string {
	return "hospitality.restaurant_charge"
}
