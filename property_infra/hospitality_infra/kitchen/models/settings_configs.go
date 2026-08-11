package kitchen_models

import (
	"time"

	"github.com/google/uuid"
)

type KitchenAutoAcceptOrders struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null"`
	PropertyID     uuid.UUID `gorm:"type:uuid"`
	KitchenID      uuid.UUID `gorm:"type:uuid;not null"`

	IsAutoAcceptEnabled bool `gorm:"default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (KitchenAutoAcceptOrders) TableName() string {
	return "hospitality.kitchen_auto_accept_orders"
}

type KitchenPrinterConfig struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null"`
	PropertyID     uuid.UUID `gorm:"type:uuid"`
	KitchenID      uuid.UUID `gorm:"type:uuid;not null"`

	PrinterName string `gorm:"type:varchar(100)"`
	IPAddress   string `gorm:"type:varchar(45)"`
	PrinterType string `gorm:"type:varchar(50)"` // Thermal, Impact

	PrintDuplicate bool `gorm:"default:false"`
	IsEnabled      bool `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (KitchenPrinterConfig) TableName() string {
	return "hospitality.kitchen_printer_configs"
}

type KitchenThrottlingConfig struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null"`
	PropertyID     uuid.UUID `gorm:"type:uuid"`
	KitchenID      uuid.UUID `gorm:"type:uuid;not null"`

	MaxOrdersPerSlot    int  `gorm:"default:0"` // 0 = Unlimited
	SlotDurationMinutes int  `gorm:"default:15"`
	IsEnabled           bool `gorm:"default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (KitchenThrottlingConfig) TableName() string {
	return "hospitality.kitchen_throttling_configs"
}

type KitchenOperatingHours struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null"`
	PropertyID     uuid.UUID `gorm:"type:uuid"`
	KitchenID      uuid.UUID `gorm:"type:uuid;not null"`

	DayOfWeek   int    `gorm:"not null"`        // 0 = Sunday, 1 = Monday...
	OpeningTime string `gorm:"type:varchar(5)"` // e.g., "06:00"
	ClosingTime string `gorm:"type:varchar(5)"` // e.g., "22:00"
	IsClosed    bool   `gorm:"default:false"`   // For holidays or manual overrides

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (KitchenOperatingHours) TableName() string {
	return "hospitality.kitchen_operating_hours"
}
