package hospitalitymodels

import (
	"time"

	"sharpnr-engine/core/models"

	"github.com/google/uuid"
)

// -------------------------------------------
// 1️⃣ Kitchen
// -------------------------------------------
type Kitchen struct {
	ID             uuid.UUID           `gorm:"type:uuid;primaryKey"`
	Name           string              `gorm:"size:100;not null"`
	Description    string              `gorm:"size:255"`
	OrganizationID uuid.UUID           `gorm:"type:uuid;not null"`
	Organization   models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;"`

	// Optional: you can later group kitchens by type
	Type string `gorm:"size:50"` // e.g., MAIN, BAR, BAKERY, ROOM_SERVICE

	MenuItems []MenuItem `gorm:"foreignKey:KitchenID;constraint:OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// -------------------------------------------
// 2️⃣ KitchenTicket (KOT)
// -------------------------------------------
// When an order is placed, one or more KOTs are generated.
// Each KOT is assigned to a specific Kitchen.
type KitchenTicket struct {
	ID             uuid.UUID           `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID           `gorm:"type:uuid;not null"`
	Organization   models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;"`

	KitchenID uuid.UUID `gorm:"type:uuid;not null"`
	Kitchen   Kitchen   `gorm:"foreignKey:KitchenID;constraint:OnDelete:CASCADE;"`

	OrderID uuid.UUID `gorm:"type:uuid;not null"`

	Status   string `gorm:"size:20;default:'PENDING'"` // PENDING, IN_PROGRESS, COMPLETED, CANCELLED
	TicketNo string `gorm:"size:50;uniqueIndex"`       // KOT number (auto-generated)

	Items []KitchenTicketItem `gorm:"foreignKey:TicketID;constraint:OnDelete:CASCADE;"`

	CreatedAt   time.Time
	UpdatedAt   time.Time
	CompletedAt *time.Time
}

// -------------------------------------------
// 3️⃣ KitchenTicketItem (KOT Line Item)
// -------------------------------------------
// Each line in a KOT corresponds to an OrderItem
type KitchenTicketItem struct {
	ID       uuid.UUID     `gorm:"type:uuid;primaryKey"`
	TicketID uuid.UUID     `gorm:"type:uuid;not null"`
	Ticket   KitchenTicket `gorm:"foreignKey:TicketID;constraint:OnDelete:CASCADE;"`

	OrderItemID uuid.UUID            `gorm:"type:uuid;not null"`
	OrderItem   HospitalityOrderItem `gorm:"foreignKey:OrderItemID;constraint:OnDelete:CASCADE;"`

	MenuItemID uuid.UUID `gorm:"type:uuid;not null"`
	MenuItem   MenuItem  `gorm:"foreignKey:MenuItemID;constraint:OnDelete:CASCADE;"`

	Quantity  int    `gorm:"not null"`
	Status    string `gorm:"size:20;default:'PENDING'"` // PENDING, COOKING, READY
	Remarks   string `gorm:"size:255"`
	IsPrinted bool   `gorm:"default:false"` // for physical KOT print tracking

	CreatedAt time.Time
	UpdatedAt time.Time
}
