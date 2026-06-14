package hospitality_order_engine_models

import (
	"time"

	"github.com/google/uuid"
)

type OrderKitchenRouting struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrderItemID uuid.UUID
	KitchenID   uuid.UUID

	KitchenTicketID *uuid.UUID

	Status string
	// PENDING, SENT, COMPLETED

	CreatedAt time.Time
}

func (OrderKitchenRouting) TableName() string {
	return "hospitality.order_kitchen_routing"
}
