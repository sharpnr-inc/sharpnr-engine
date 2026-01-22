package order_engine_models

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
