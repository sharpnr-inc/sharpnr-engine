package restaurant_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

/* ----------------------------------- <*> ---------------------------------- */

type RestaurantSettings struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey"`
	RestaurantID uuid.UUID      `gorm:"type:uuid;not null;index"`
	Data         datatypes.JSON `gorm:"type:jsonb;not null;default:'{}'"`
}

func (RestaurantSettings) TableName() string {
	return "hospitality.restaurant_settings"
}

/* ----------------------------------- <*> ---------------------------------- */

type RestaurantOrderNumberingSettings struct {
	ID                   uuid.UUID `gorm:"type:uuid;primaryKey"`
	RestaurantSettingsID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`

	Prefix        string `gorm:"type:varchar(20);default:'ORD'"`
	StartFrom     int64  `gorm:"default:1"`
	CurrentNumber int64  `gorm:"default:1"`

	ResetType string `gorm:"type:varchar(20);default:'daily'"`
	// none | daily | monthly | yearly

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RestaurantOrderNumberingSettings) TableName() string {
	return "hospitality.restaurant_order_numbering_settings"
}

/* ----------------------------------- <*> ---------------------------------- */

type RestaurantKitchenSettings struct {
	ID                   uuid.UUID `gorm:"type:uuid;primaryKey"`
	RestaurantSettingsID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`

	AutoAcceptOrders bool `gorm:"default:true"`
	AutoPrint        bool `gorm:"default:false"`
	GroupByCourse    bool `gorm:"default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RestaurantKitchenSettings) TableName() string {
	return "hospitality.restaurant_kitchen_settings"
}

/* ----------------------------------- <*> ---------------------------------- */

type RestaurantPOSSettings struct {
	ID                   uuid.UUID `gorm:"type:uuid;primaryKey"`
	RestaurantSettingsID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`

	EnableTableSelection bool `gorm:"default:true"`
	EnableSplitBill      bool `gorm:"default:true"`
	EnableDiscount       bool `gorm:"default:true"`
	RequireCustomerInfo  bool `gorm:"default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RestaurantPOSSettings) TableName() string {
	return "hospitality.restaurant_pos_settings"
}

/* ----------------------------------- <*> ---------------------------------- */

type RestaurantTaxSettings struct {
	ID                   uuid.UUID `gorm:"type:uuid;primaryKey"`
	RestaurantSettingsID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`

	DefaultTaxRate float64 `gorm:"type:numeric(5,2);default:0"`
	IsTaxInclusive bool    `gorm:"default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RestaurantTaxSettings) TableName() string {
	return "hospitality.restaurant_tax_settings"
}

/* ----------------------------------- <*> ---------------------------------- */

type RestaurantInvoiceSettings struct {
	ID                   uuid.UUID `gorm:"type:uuid;primaryKey"`
	RestaurantSettingsID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`

	InvoicePrefix  string `gorm:"type:varchar(20);default:'INV'"`
	ShowLogo       bool   `gorm:"default:true"`
	ShowTaxBreakup bool   `gorm:"default:true"`

	FooterNote string `gorm:"type:text"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RestaurantInvoiceSettings) TableName() string {
	return "hospitality.restaurant_invoice_settings"
}

/* ----------------------------------- <*> ---------------------------------- */

type RestaurantPreferedKitchen struct {
	ID                   uuid.UUID `gorm:"type:uuid;primaryKey"`
	RestaurantSettingsID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`

	KitchenID uuid.UUID `gorm:"type:uuid;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RestaurantPreferedKitchen) TableName() string {
	return "hospitality.restaurant_prefered_kitchen"
}
