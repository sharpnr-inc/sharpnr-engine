package restaurant_models

import (
	restaurant_constants "sharpnr-engine/modules/hospitality/restaurant/constants"
	"time"

	"github.com/google/uuid"
)

/* ------------------------ Restaurant Dining Session ----------------------- */

type RestaurantDiningSession struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	RestaurantID uuid.UUID `gorm:"type:uuid;not null;index"`
	TableID      uuid.UUID `gorm:"type:uuid;not null;index"`

	Status restaurant_constants.RestaurantDiningSessionStatusType `gorm:"size:20;not null"`
	// OPEN, CLOSED, CANCELLED

	StartedAt time.Time
	StartedBy uuid.UUID
	ClosedAt  *time.Time

	CreatedAt time.Time
}

func (RestaurantDiningSession) TableName() string {
	return "hospitality.restaurant_dining_session"
}

/* ------------------------- Dining Session Customer ------------------------ */

type RestaurantDiningSessionCustomer struct {
	ID                          uuid.UUID
	RestaurantDiningSessionID   uuid.UUID
	HospitalityDomainCustomerID uuid.UUID
	CreatedAt                   time.Time
}

func (RestaurantDiningSessionCustomer) TableName() string {
	return "hospitality.restaurant_dining_session_customer"
}

/* ----------------------- Restaurant Session payments ---------------------- */

type RestaurantSessionPayment struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID
	PropertyID     uuid.UUID
	RestaurantID   uuid.UUID
	LegalEntityID  uuid.UUID

	RestaurantDiningSessionID uuid.UUID `gorm:"type:uuid;not null;index"`

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

/* ----------------------------------- <> ----------------------------------- */
