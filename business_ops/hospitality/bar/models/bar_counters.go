package bar_models

import (
	"time"

	"github.com/google/uuid"
)

type BarCounter struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID
	PropertyID     uuid.UUID
	BarID          uuid.UUID

	ServicePointID uuid.UUID
	// references PropertyServicePoint (Type = BAR)

	Name string
	// "Main Bar Counter"

	IsActive bool

	CreatedAt time.Time
}

func (BarCounter) TableName() string {
	return "hospitality.bar_counter"
}

type BarTicket struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	PropertyID     uuid.UUID
	BarID          uuid.UUID
	ServicePointID uuid.UUID // Bar Counter

	Status string
	// PENDING, SERVED

	Items []BarTicketItem

	CreatedAt time.Time
}

func (BarTicket) TableName() string {
	return "hospitality.bar_ticket"
}

type BarTicketItem struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	BarTicketID uuid.UUID
	MenuItemID  uuid.UUID

	Quantity int

	CreatedAt time.Time
}

func (BarTicketItem) TableName() string {
	return "hospitality.bar_ticket_item"
}

type BarOrderTicketLink struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	BarOrderID uuid.UUID `gorm:"type:uuid;index"`
	TicketID   uuid.UUID `gorm:"type:uuid;index"`

	TicketType string
	// KOT, BOT

	CreatedAt time.Time
}

func (BarOrderTicketLink) TableName() string {
	return "hospitality.bar_order_ticket_link"
}
