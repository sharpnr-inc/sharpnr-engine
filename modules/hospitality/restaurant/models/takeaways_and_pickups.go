package restaurant_models

import (
	"time"

	"github.com/google/uuid"
)

type RestaurantPickUpOrders struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`
	RestaurantID   uuid.UUID `gorm:"type:uuid;not null;index"`
	LegalEntityID  uuid.UUID `gorm:"type:uuid;not null;index"`

	HospitalityDomainCustomerID *uuid.UUID `gorm:"type:uuid;index"`

	OrderNumber string `gorm:"type:varchar(50);index"`

	Status string `gorm:"type:varchar(30);not null;index"`
	// PENDING, PREPARING, READY, COMPLETED, CANCELLED

	TotalAmount float64 `gorm:"type:numeric(12,2);not null"`

	IsPaid bool `gorm:"default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RestaurantPickUpOrders) TableName() string {
	return "hospitality.restaurant_pickup_orders"
}
