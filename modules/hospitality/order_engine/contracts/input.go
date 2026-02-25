package hospitality_order_engine_contracts

import "github.com/google/uuid"

type CreateOrderInput struct {
	OwnersID   uuid.UUID `json:"ownersID" binding:"required"`
	MenuItemID uuid.UUID `json:"menuItemID" binding:"required"`
	Quantity   int       `json:"quantity" binding:"required"`
	Notes      string    `json:"notes"`
}
