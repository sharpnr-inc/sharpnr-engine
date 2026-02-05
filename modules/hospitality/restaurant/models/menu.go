package restaurant_models

import "github.com/google/uuid"

type RestaurantMenuMapping struct {
	ID           uuid.UUID
	RestaurantID uuid.UUID
	MenuID       uuid.UUID
}
