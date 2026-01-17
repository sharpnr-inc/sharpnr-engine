package bar_models

import (
	"time"

	"sharpnr-engine/core/organization/models"
	kitchen_models "sharpnr-engine/modules/hospitality/kitchen/models"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"

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

	IsKitchen bool                    `gorm:"default:true" json:"is_kitchen"`
	KitchenID *uuid.UUID              `gorm:"type:uuid" json:"kitchen_id"`
	Kitchen   *kitchen_models.Kitchen `gorm:"foreignKey:KitchenID;constraint:OnDelete:SET NULL;" json:"kitchen"`

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

type BarTableSession struct {
	ID         uuid.UUID
	BarTableId uuid.UUID
	BarTable   *BarTable
	BarID      uuid.UUID
	Bar        *Bar

	HospitalityCustomerAndGuestID *uuid.UUID `gorm:"type:uuid"`

	FromTime time.Time  `gorm:"type:timestamp;not null"` // When session started
	ToTime   *time.Time `gorm:"type:timestamp"`          // When closed (null until closing)

	OrderItems []BarTablesSessionOrder

	SubTotal    float64 `gorm:"type:numeric(10,2);default:0"`
	GSTAndTaxes float64 `gorm:"type:numeric(10,2);default:0"`
	Total       float64 `gorm:"type:numeric(10,2);default:0"`
}

type BarTablesSessionOrder struct {
	ID                uuid.UUID
	BarTableSessionId uuid.UUID
	BarTableSession   *BarTableSession

	MenuItemID uuid.UUID                       `gorm:"type:uuid;not null"`
	MenuItem   hospitalityMenu_models.MenuItem `gorm:"foreignKey:MenuItemID"`

	// -----------------------------------------------------
	// 🔗 Variant (Optional - because not all menu items have variants)
	// -----------------------------------------------------
	VariantID *uuid.UUID                          `gorm:"type:uuid"`
	Variant   *hospitalityMenu_models.MenuVariant `gorm:"foreignKey:VariantID"`

	// -----------------------------------------------------
	// 🔗 Addons (Optional)
	// -----------------------------------------------------
	AddonID *uuid.UUID                    `gorm:"type:uuid"`
	Addon   *hospitalityMenu_models.Addon `gorm:"foreignKey:AddonID"`

	Quantity int    `gorm:"not null;default:1"`
	Notes    string `gorm:"type:text"`
}
