package hospitalitymodels

import (
	"time"

	"sharpnr-engine/core/models"

	"github.com/google/uuid"
)

// ---------------------------
// 1️⃣ RESTAURANT
// ---------------------------

type Restaurant struct {
	ID             uuid.UUID            `gorm:"type:uuid;primaryKey" json:"id"`
	Name           string               `gorm:"size:100;not null" json:"name"`
	OrganizationID uuid.UUID            `gorm:"type:uuid;not null;index" json:"organization_id"`
	Organization   *models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;" json:"organization,omitempty"`

	IsPureVeg   bool              `gorm:"default:false" json:"is_pure_veg"`
	OpeningTime string            `gorm:"size:5" json:"opening_time"` // e.g., "09:00"
	ClosingTime string            `gorm:"size:5" json:"closing_time"` // e.g., "22:00"
	Floors      []RestaurantFloor `gorm:"foreignKey:RestaurantID;constraint:OnDelete:CASCADE;" json:"floors,omitempty"`

	IsKitchen bool       `gorm:"default:true" json:"is_kitchen"`
	KitchenID *uuid.UUID `gorm:"type:uuid;index" json:"kitchen_id,omitempty"`
	Kitchen   *Kitchen   `gorm:"foreignKey:KitchenID;constraint:OnDelete:SET NULL;" json:"kitchen,omitempty"`

	// Legal & Business Info
	LegalName   string `gorm:"size:255" json:"legal_name"`
	DisplayName string `gorm:"size:255" json:"display_name"`
	GSTNumber   string `gorm:"size:50" json:"gst_number"`
	FSSAINumber string `gorm:"size:50" json:"fssai_number"`
	CINNumber   string `gorm:"size:50" json:"cin_number"`

	// Contact Info
	ContactEmail  string `gorm:"size:100" json:"contact_email"`
	ContactNumber string `gorm:"size:20" json:"contact_number"`
	ManagerName   string `gorm:"size:100" json:"manager_name"`

	// Address
	Address string `gorm:"size:500" json:"address"`
	City    string `gorm:"size:100" json:"city"`
	State   string `gorm:"size:100" json:"state"`
	Pincode string `gorm:"size:10" json:"pincode"`
	Country string `gorm:"size:100;default:'India'" json:"country"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ---------------------------
// 2️⃣ RESTAURANT FLOOR
// ---------------------------

type RestaurantFloor struct {
	ID             uuid.UUID            `gorm:"type:uuid;primaryKey" json:"id"`
	RestaurantID   uuid.UUID            `gorm:"type:uuid;not null;index" json:"restaurant_id"`
	Restaurant     *Restaurant          `gorm:"foreignKey:RestaurantID;constraint:OnDelete:CASCADE;" json:"-"` // Exclude from JSON to avoid circular refs
	OrganizationID uuid.UUID            `gorm:"type:uuid;not null;index" json:"organization_id"`
	Organization   *models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;" json:"-"`

	Name        string `gorm:"size:100;not null" json:"name"`
	Description string `gorm:"size:255" json:"description"`
	FloorNumber int    `gorm:"default:0" json:"floor_number"`

	Tables []RestaurantTable `gorm:"foreignKey:FloorID;constraint:OnDelete:CASCADE;" json:"tables,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ---------------------------
// 3️⃣ RESTAURANT TABLE
// ---------------------------

type RestaurantTable struct {
	ID           uuid.UUID        `gorm:"type:uuid;primaryKey" json:"id"`
	RestaurantID uuid.UUID        `gorm:"type:uuid;not null;index" json:"restaurant_id"`
	Restaurant   *Restaurant      `gorm:"foreignKey:RestaurantID;constraint:OnDelete:CASCADE;" json:"-"`
	FloorID      uuid.UUID        `gorm:"type:uuid;index" json:"floor_id"`
	Floor        *RestaurantFloor `gorm:"foreignKey:FloorID;constraint:OnDelete:SET NULL;" json:"-"`

	Name     string `gorm:"size:50;not null" json:"name"`
	Capacity int    `gorm:"not null" json:"capacity"`
	Status   string `gorm:"size:20;default:'AVAILABLE'" json:"status"` // AVAILABLE, OCCUPIED, RESERVED, BLOCKED

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RestaurantTableBookingSession struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// -----------------------------------------------------
	// 🔗 Table Reference (Required)
	// -----------------------------------------------------
	RestaurantTableID uuid.UUID       `gorm:"type:uuid;not null"`
	RestaurantTable   RestaurantTable `gorm:"foreignKey:RestaurantTableID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	// -----------------------------------------------------
	// 🔗 Customer (Optional until closing the session)
	// -----------------------------------------------------
	HospitalityCustomerAndGuestID *uuid.UUID                   `gorm:"type:uuid"`
	CustomerAndGuest              *HospitalityCustomerAndGuest `gorm:"foreignKey:HospitalityCustomerAndGuestID"`

	// -----------------------------------------------------
	// 🔗 Orders for the session
	// -----------------------------------------------------
	Orders []RestaurantTableOrder `gorm:"foreignKey:RestaurantTableBookingSessionID"`

	// -----------------------------------------------------
	// 🕒 Session Timing
	// -----------------------------------------------------
	FromTime time.Time  `gorm:"type:timestamp;not null"` // When session started
	ToTime   *time.Time `gorm:"type:timestamp"`          // When closed (null until closing)

	CreatedAt time.Time
	UpdatedAt time.Time

	// -----------------------------------------------------
	// 💰 Billing Fields
	// -----------------------------------------------------
	SubTotal    float64 `gorm:"type:numeric(10,2);default:0"`
	GSTAndTaxes float64 `gorm:"type:numeric(10,2);default:0"`
	Total       float64 `gorm:"type:numeric(10,2);default:0"`

	// -----------------------------------------------------
	// 🧾 Invoice (Optional, Added After Billing)
	// -----------------------------------------------------
	RestaurantInvoiceID *uuid.UUID `gorm:"type:uuid"`
}

type RestaurantTableOrder struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// -----------------------------------------------------
	// 🔗 Menu Item (Required)
	// -----------------------------------------------------
	MenuItemID uuid.UUID `gorm:"type:uuid;not null"`
	MenuItem   MenuItem  `gorm:"foreignKey:MenuItemID"`

	// -----------------------------------------------------
	// 🔗 Variant (Optional - because not all menu items have variants)
	// -----------------------------------------------------
	VariantID *uuid.UUID   `gorm:"type:uuid"`
	Variant   *MenuVariant `gorm:"foreignKey:VariantID"`

	// -----------------------------------------------------
	// 🔗 Addons (Optional)
	// -----------------------------------------------------
	AddonID *uuid.UUID `gorm:"type:uuid"`
	Addon   *Addon     `gorm:"foreignKey:AddonID"`

	// -----------------------------------------------------
	// 🔢 Quantity & Notes
	// -----------------------------------------------------
	Quantity int    `gorm:"not null;default:1"`
	Notes    string `gorm:"type:text"`

	// -----------------------------------------------------
	// 🔗 Link to Table Booking Session
	// -----------------------------------------------------
	RestaurantTableBookingSessionID uuid.UUID                     `gorm:"type:uuid;not null"`
	BookingSession                  RestaurantTableBookingSession `gorm:"foreignKey:RestaurantTableBookingSessionID"`

	// -----------------------------------------------------
	// 🔗 Table (for history, not mandatory if session links)
	// -----------------------------------------------------
	RestaurantTableID uuid.UUID       `gorm:"type:uuid;not null"`
	RestaurantTable   RestaurantTable `gorm:"foreignKey:RestaurantTableID"`
}
