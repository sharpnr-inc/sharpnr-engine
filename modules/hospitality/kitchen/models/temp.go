package kitchen_models

import (
	"time"

	"github.com/google/uuid"
)

type KitchenTicketStaging struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null"`
	PropertyID     uuid.UUID `gorm:"type:uuid"`
	KitchenID      uuid.UUID `gorm:"type:uuid;not null"`

	IsAccepted *bool
	OrderID    uuid.UUID `gorm:"type:uuid;not null"`

	CreatedAt   time.Time
	UpdatedAt   time.Time
	CompletedAt *time.Time
}

func (KitchenTicketStaging) TableName() string {
	return "hospitality.kitchen_ticket_staging"
}

type KitchenTicketItemStaging struct {
	ID                     uuid.UUID `gorm:"type:uuid;primaryKey"`
	KitchenTicketStagingID uuid.UUID `gorm:"type:uuid;not null"`
	MenuItemID             uuid.UUID `gorm:"type:uuid;not null"`

	MenuItemName string `gorm:"size:120"`
	Quantity     int    `gorm:"not null"`
	Remarks      string `gorm:"size:255"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (KitchenTicketItemStaging) TableName() string {
	return "hospitality.kitchen_ticket_item_staging"
}

type KitchenTicketItemVariantStaging struct {
	ID                     uuid.UUID `gorm:"type:uuid;primaryKey"`
	KitchenTicketStagingID uuid.UUID `gorm:"type:uuid;not null"`

	VariantID   uuid.UUID `gorm:"type:uuid;not null"` // NOTE --> this is the "Varaint in the main table"
	VariantName string

	CreatedAt time.Time
}

func (KitchenTicketItemVariantStaging) TableName() string {
	return "hospitality.kitchen_ticket_item_variant_staging"
}

type KitchenTicketItemAddonStaging struct {
	ID                     uuid.UUID `gorm:"type:uuid;primaryKey"`
	KitchenTicketStagingID uuid.UUID `gorm:"type:uuid;not null"`

	AddOnID  uuid.UUID `gorm:"type:uuid;not null"`
	Quantity int       `gorm:"not null;default:0"`

	CreatedAt time.Time
}

func (KitchenTicketItemAddonStaging) TableName() string {
	return "hospitality.kitchen_ticket_item_addon_staging"
}
