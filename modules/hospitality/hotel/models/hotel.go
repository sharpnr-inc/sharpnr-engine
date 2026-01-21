package hotel_models

import (
	"time"

	"github.com/google/uuid"
)

type Hotel struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;not null"`
	PropertyID     uuid.UUID `gorm:"type:uuid;index;not null"`
	LegalEntityID  uuid.UUID `gorm:"type:uuid;index;not null"`

	Name         string `gorm:"size:100"`
	CheckInTime  string `gorm:"default:10:00"`
	CheckOutTime string `gorm:"default:13:00"`

	IsActive bool `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
