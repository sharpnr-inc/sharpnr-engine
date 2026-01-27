package kitchen_models

import (
	"time"

	"github.com/google/uuid"
)

type Kitchen struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID         uuid.UUID `gorm:"type:uuid;not null"`
	OrganizationPropertyID uuid.UUID `gorm:"type:uuid;not null"`

	Name        string `gorm:"size:100;not null"`
	Description string
	Type        string `gorm:"size:50"` // e.g., MAIN, BAR, BAKERY, ROOM_SERVICE
	IsActive    bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Kitchen) TableName() string {
	return "hospitality.kitchen"
}

type KitchenMenuItem struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	KitchenID  uuid.UUID `gorm:"type:uuid;index;not null"`
	MenuItemID uuid.UUID `gorm:"type:uuid;index;not null"`

	IsActive bool `gorm:"default:true"`

	CreatedAt time.Time
}

func (KitchenMenuItem) TableName() string {
	return "hospitality.kitchen_menu_item"
}

type KitchenTicket struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null"`
	PropertyID     uuid.UUID

	KitchenID      uuid.UUID `gorm:"type:uuid;not null"`
	ServicePointID uuid.UUID `gorm:"type:uuid;index;not null"`

	Status   string `gorm:"size:20;default:'PENDING'"`
	TicketNo string `gorm:"size:50;uniqueIndex"`

	CreatedAt   time.Time
	UpdatedAt   time.Time
	CompletedAt *time.Time
}

func (KitchenTicket) TableName() string {
	return "hospitality.kitchen_ticket"
}

// -------------------------------------------
// 3️⃣ KitchenTicketItem (KOT Line Item)
// -------------------------------------------
// Each line in a KOT corresponds to an OrderItem
type KitchenTicketItem struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	KitchenTicketID uuid.UUID `gorm:"type:uuid;not null"`

	MenuItemID uuid.UUID `gorm:"type:uuid;not null"`
	VariantID  *uuid.UUID

	Quantity  int    `gorm:"not null"`
	Remarks   string `gorm:"size:255"`
	IsPrinted bool   `gorm:"default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (KitchenTicketItem) TableName() string {
	return "hospitality.kitchen_ticket_item"
}

type KitchenTicketItemAddon struct {
	ID                  uuid.UUID `gorm:"type:uuid;primaryKey"`
	KitchenTicketItemID uuid.UUID `gorm:"type:uuid;not null"`
	AddOnID             uuid.UUID `gorm:"type:uuid;not null"`
	Quantity            int       `gorm:"not null;default:0"`
	CreatedAt           time.Time
}

func (KitchenTicketItemAddon) TableName() string {
	return "hospitality.kitchen_ticket_item_addon"
}
