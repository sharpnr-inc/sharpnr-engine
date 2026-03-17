package restaurant_models

import (
	restaurant_constants "sharpnr-engine/modules/hospitality/restaurant/constants"
	"time"

	"github.com/google/uuid"
)

/* ------------------------ Restaurant Dining Session ----------------------- */

type RestaurantDiningSession struct {
	ID           uuid.UUID                                              `gorm:"type:uuid;primaryKey"`
	RestaurantID uuid.UUID                                              `gorm:"type:uuid;not null;index"`
	TableID      uuid.UUID                                              `gorm:"type:uuid;not null;index"`
	Status       restaurant_constants.RestaurantDiningSessionStatusType `gorm:"size:20;not null"`

	StartedAt time.Time
	StartedBy uuid.UUID
	ClosedAt  *time.Time
	CreatedAt time.Time
}

func (RestaurantDiningSession) TableName() string {
	return "hospitality.restaurant_dining_session"
}

/* ---------------------------------- <***> --------------------------------- */
