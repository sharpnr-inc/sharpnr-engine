package restaurant_contracts

import "github.com/google/uuid"

type RestaurantDetails struct {
	ID           uuid.UUID
	Name         string
	Status       bool
	IsConfigured bool
}
