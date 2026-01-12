package hospitality_inventory_models

import (
	"time"

	"github.com/google/uuid"
)

type HospitalityInventory struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrgID      uuid.UUID `gorm:"type:uuid;index"`
	Module     string    `gorm:"size:20;index"` // BAR | RESTAURANT | HOTEL
	Name       string
	SKU        string
	CategoryID uuid.UUID `gorm:"type:uuid"`
	Unit       string    // ml, ltr, pcs, kg
	IsActive   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type HospitalityInventoryCategory struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrgID     uuid.UUID `gorm:"type:uuid;index"`
	Module    string    `gorm:"size:20;index"`
	Name      string
	CreatedAt time.Time
}

type HospitalityInventoryLocation struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrgID     uuid.UUID `gorm:"type:uuid;index"`
	Module    string    `gorm:"size:20;index"`
	Name      string
	CreatedAt time.Time
}

type HospitalityInventoryStock struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrgID       uuid.UUID `gorm:"type:uuid;index"`
	InventoryID uuid.UUID `gorm:"type:uuid;index"`
	LocationID  uuid.UUID `gorm:"type:uuid;index"`
	Quantity    float64
	UpdatedAt   time.Time
}

// Phase 2 tables

type HospitalityInventorySupplier struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrgID uuid.UUID `gorm:"type:uuid;index"`
	Name  string
	Phone string
	GSTIN string
}
