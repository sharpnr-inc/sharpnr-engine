package restaurant_models

import (
	restaurant_constants "sharpnr-engine/modules/hospitality/restaurant/constants"

	"github.com/google/uuid"
)

type RestaurantOperationStatus struct {
	ID           uuid.UUID                                        `gorm:"type:uuid;primaryKey"`
	RestaurantID uuid.UUID                                        `gorm:"type:uuid;not null;index"`
	Status       restaurant_constants.RestaurantRunningStatusType `gorm:"type:varchar(50);not null"`
}

func (RestaurantOperationStatus) TableName() string {
	return "hospitality.restaurant_operation_statuses"
}

type RestaurantConfigurationStatus struct {
	ID           uuid.UUID                                        `gorm:"type:uuid;primaryKey"`
	RestaurantID uuid.UUID                                        `gorm:"type:uuid;not null;index"`
	Status       restaurant_constants.RestaurantConfigurationType `gorm:"type:varchar(50);not null"`
}

func (RestaurantConfigurationStatus) TableName() string {
	return "hospitality.restaurant_configuration_statuses"
}
