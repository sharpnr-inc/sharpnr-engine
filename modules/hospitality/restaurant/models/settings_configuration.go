package restaurant_models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type RestaurantSettings struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey"`
	RestaurantID uuid.UUID      `gorm:"type:uuid;not null;index"`
	Data         datatypes.JSON `gorm:"type:jsonb;not null;default:'{}'"`
}

func (RestaurantSettings) TableName() string {
	return "hospitality.restaurant_settings"
}
