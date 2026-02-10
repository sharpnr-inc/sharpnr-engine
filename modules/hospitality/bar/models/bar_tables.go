package bar_models

import (
	"time"

	"github.com/google/uuid"
)

type BarFloor struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	BarID      uuid.UUID `gorm:"type:uuid;index;not null"`
	PropertyID uuid.UUID `gorm:"type:uuid;index;not null"`

	Name        string
	FloorNumber int

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (BarFloor) TableName() string {
	return "hospitality.bar_floor"
}

// ---------------------------------------------
// 3️⃣ BAR TABLE
// ---------------------------------------------
type BarTable struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	BarID      uuid.UUID  `gorm:"type:uuid;index;not null"`
	FloorID    *uuid.UUID `gorm:"type:uuid"`
	PropertyID uuid.UUID  `gorm:"type:uuid;index;not null"`

	Code     string `gorm:"size:50;not null"` // B1, VIP-01
	Capacity int    `gorm:"not null"`

	Status string
	// AVAILABLE, OCCUPIED, RESERVED, BLOCKED

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (BarTable) TableName() string {
	return "hospitality.bar_table"
}

type BarTableSession struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	BarID      uuid.UUID `gorm:"type:uuid;index;not null"`
	TableID    uuid.UUID `gorm:"type:uuid;index;not null"`
	PropertyID uuid.UUID `gorm:"type:uuid;index;not null"`

	GuestID *uuid.UUID

	Status string
	// OPEN, CLOSED, CANCELLED

	OpenedAt time.Time
	ClosedAt *time.Time
}

func (BarTableSession) TableName() string {
	return "hospitality.bar_table_session"
}
