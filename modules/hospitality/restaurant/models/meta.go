package restaurant_models

import "github.com/google/uuid"

type RestaurantSetup struct {
	ID           uuid.UUID
	RestaurantID uuid.UUID
	IsSetup      bool
	IsConfigured bool
}
