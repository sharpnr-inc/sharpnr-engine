package restaurant_contracts

import (
	restaurant_constants "sharpnr-engine/modules/hospitality/restaurant/constants"

	"github.com/google/uuid"
)

type RestaurantDetails struct {
	ID                  uuid.UUID
	PropertyID          uuid.UUID
	Name                string
	Status              bool
	IsConfigured        bool
	ConfigurationStatus restaurant_constants.RestaurantConfigurationStatusType
}
