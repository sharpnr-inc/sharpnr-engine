package restaurant_events

import (
	"sharpnr-engine/sero"
	"time"

	"github.com/google/uuid"
)

const (
	OrderCreatedInRestaurant sero.SeroEngineEventType = "ORDER_CREATED_IN_RESTAURANT"
)

type OrderCreatedPayload struct {
	OrderID        uuid.UUID `json:"order_id"`
	OrganizationID uuid.UUID
	DineSessionID  uuid.UUID `json:"dine_session_id"`
	RestaurantID   uuid.UUID `json:"restaurant_id"`
	OrderNumber    string    `json:"order_number"`
	ItemsCount     int       `json:"items_count"`
	CreatedAt      time.Time `json:"created_at"`
}
