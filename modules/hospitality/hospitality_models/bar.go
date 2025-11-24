package hospitalitymodels

import (
	"time"

	"sharpnr-engine/core/models"

	"github.com/google/uuid"
)

// ---------------------------------------------
// 1️⃣ BAR
// ---------------------------------------------
type Bar struct {
	ID             uuid.UUID            `gorm:"type:uuid;primaryKey" json:"id"`
	Name           string               `gorm:"size:100;not null" json:"name"`
	OrganizationID uuid.UUID            `gorm:"type:uuid;not null" json:"organization_id"`
	Organization   *models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;" json:"organization"`

	OpeningTime time.Time `json:"opening_time"`
	ClosingTime time.Time `json:"closing_time"`

	// Optional: If bar has multiple floors (e.g., rooftop bar)
	Floors []BarFloor `gorm:"foreignKey:BarID;constraint:OnDelete:CASCADE;" json:"floors"`

	IsKitchen bool       `gorm:"default:true" json:"is_kitchen"`
	KitchenID *uuid.UUID `gorm:"type:uuid" json:"kitchen_id"`
	Kitchen   *Kitchen   `gorm:"foreignKey:KitchenID;constraint:OnDelete:SET NULL;" json:"kitchen"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// ========== General + Legal + Display Details (Sync with ORG if enabled) ==========

	IsShareLegalDetailsWithMainOrg bool `gorm:"default:false" json:"is_share_legal_details_with_main_org"`

	DisplayName       string `gorm:"size:150" json:"display_name"`
	LegalName         string `gorm:"size:150" json:"legal_name"`
	RegisteredAddress string `gorm:"size:255" json:"registered_address"`
	GSTNumber         string `gorm:"size:20" json:"gst_number"`
	CINNumber         string `gorm:"size:20" json:"cin_number"`
	FSSAINumber       string `gorm:"size:20" json:"fssai_number"`

	IsMultifloor bool   `gorm:"default:false" json:"is_multifloor"`
	AddressLine  string `gorm:"size:255" json:"address_line"`
	City         string `gorm:"size:100" json:"city"`
	State        string `gorm:"size:100" json:"state"`
	Country      string `gorm:"size:100" json:"country"`
	Pincode      string `gorm:"size:20" json:"pincode"`
}

// ---------------------------------------------
// 2️⃣ BAR FLOOR (Optional Layer)
// ---------------------------------------------
type BarFloor struct {
	ID             uuid.UUID            `gorm:"type:uuid;primaryKey" json:"id"`
	BarID          uuid.UUID            `gorm:"type:uuid;not null" json:"bar_id"` // ✅ match type
	Bar            *Bar                 `gorm:"foreignKey:BarID;constraint:OnDelete:CASCADE;" json:"bar"`
	OrganizationID uuid.UUID            `gorm:"type:uuid;not null" json:"organization_id"`
	Organization   *models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;" json:"organization"`

	Name        string `gorm:"size:100;not null" json:"name"` // e.g., "Ground Floor", "Rooftop"
	Description string `gorm:"size:255" json:"description"`
	FloorNumber int    `gorm:"default:0" json:"floor_number"` // For sorting and UI display order

	Tables []BarTable `gorm:"foreignKey:FloorID;constraint:OnDelete:CASCADE;" json:"tables"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ---------------------------------------------
// 3️⃣ BAR TABLE
// ---------------------------------------------
type BarTable struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	FloorID uuid.UUID `gorm:"type:uuid" json:"floor_id"`
	Floor   *BarFloor `gorm:"foreignKey:FloorID;constraint:OnDelete:SET NULL;" json:"floor"`

	Name     string `gorm:"size:50;not null" json:"name"` // e.g., "B1", "VIP Table 1"
	Capacity int    `gorm:"not null;default:4" json:"capacity"`
	Status   string `gorm:"size:20;default:'AVAILABLE'" json:"status"` // AVAILABLE, OCCUPIED, RESERVED

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// -------------------------------------------
// 1️⃣ BarTicket (BOT)
// -------------------------------------------
type BarTicket struct {
	ID             uuid.UUID            `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID            `gorm:"type:uuid;not null"`
	Organization   *models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;"`

	BarID uuid.UUID `gorm:"type:uuid;not null"`
	Bar   *Bar      `gorm:"foreignKey:BarID;constraint:OnDelete:CASCADE;"`

	OrderID uuid.UUID        `gorm:"type:uuid;not null"`
	Order   HospitalityOrder `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`

	Status   string `gorm:"size:20;default:'PENDING'"` // PENDING, MIXING, SERVED, CANCELLED
	TicketNo string `gorm:"size:50;uniqueIndex"`

	Items []BarTicketItem `gorm:"foreignKey:TicketID;constraint:OnDelete:CASCADE;"`

	CreatedAt   time.Time
	UpdatedAt   time.Time
	CompletedAt *time.Time
}

// -------------------------------------------
// 2️⃣ BarTicketItem (BOT Line Item)
// -------------------------------------------
type BarTicketItem struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	TicketID uuid.UUID `gorm:"type:uuid;not null"`
	Ticket   BarTicket `gorm:"foreignKey:TicketID;constraint:OnDelete:CASCADE;"`

	OrderItemID uuid.UUID            `gorm:"type:uuid;not null"`
	OrderItem   HospitalityOrderItem `gorm:"foreignKey:OrderItemID;constraint:OnDelete:CASCADE;"`

	MenuItemID uuid.UUID `gorm:"type:uuid;not null"`
	MenuItem   MenuItem  `gorm:"foreignKey:MenuItemID;constraint:OnDelete:CASCADE;"`

	Quantity  int    `gorm:"not null"`
	Status    string `gorm:"size:20;default:'PENDING'"` // PENDING, MIXING, READY
	IsPrinted bool   `gorm:"default:false"`
	Remarks   string `gorm:"size:255"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
