package bar_models

import (
	"time"

	"github.com/google/uuid"
)

type Bar struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name string    `gorm:"size:100;not null" json:"name"`

	OrganizationID uuid.UUID `gorm:"type:uuid;not null" json:"organization_id"`
	PropertyID     uuid.UUID `gorm:"type:uuid;index;not null"`
	LegalEntityID  uuid.UUID `gorm:"type:uuid;index;not null"`

	// OpeningTime  time.Time `json:"opening_time"`
	// ClosingTime  time.Time `json:"closing_time"`
	// IsMultiFloor bool      `gorm:"default:false"`

	IsActive bool `gorm:"default:false"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Bar) TableName() string {
	return "hospitality.bar"
}

type BarLayout struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	BarID        uuid.UUID
	PropertyID   uuid.UUID `gorm:"type:uuid;index;not null"`
	IsMultiFloor bool
}

func (BarLayout) TableName() string {
	return "hospitality.bar_layout"
}

type BarOperationTiming struct {
	ID          uuid.UUID
	BarID       uuid.UUID
	PropertyID  uuid.UUID `gorm:"type:uuid;index;not null"`
	OpeningTime time.Time `json:"opening_time"`
	ClosingTime time.Time `json:"closing_time"`
}

func (BarOperationTiming) TableName() string {
	return "hospitality.bar_operation_timing"
}
