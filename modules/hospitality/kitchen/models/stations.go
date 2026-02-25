package kitchen_models

import (
	"time"

	"github.com/google/uuid"
)

type KitchenStation struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	KitchenID uuid.UUID `gorm:"type:uuid;index;not null"`

	Name string
	// GRILL, FRY, BAR_COUNTER, PIZZA

	IsActive bool

	CreatedAt time.Time
}

type KitchenTicketStation struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	KitchenTicketItemID uuid.UUID `gorm:"type:uuid;index;not null"`
	KitchenStationID    uuid.UUID `gorm:"type:uuid;index;not null"`

	AssignedAt time.Time
}
