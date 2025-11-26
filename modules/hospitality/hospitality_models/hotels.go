package hospitalitymodels

import (
	"time"

	"sharpnr-engine/core/models"
	"sharpnr-engine/modules/hospitality/hospitalityconstants"

	"github.com/google/uuid"
)

// -----------------------------
// Hotel
// -----------------------------
type Hotel struct {
	ID             uuid.UUID            `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID            `gorm:"type:uuid;not null"`
	Organization   *models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;"`

	Name           string `gorm:"size:100"`
	NumberOfFloors int    `gorm:"default:0"`
	NumberOfRooms  int    `gorm:"default:0"`

	Floors []HotelFloor `gorm:"foreignKey:HotelID;constraint:OnDelete:CASCADE;"`

	IsConfigured       bool
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DisplayName        string
	LegalName          string
	GstinNumber        string
	FssaiLicenseNumber string

	IsKitchen   bool       `gorm:"default:true"`
	KitchenID   *uuid.UUID `gorm:"type:uuid"`
	Kitchen     *Kitchen   `gorm:"foreignKey:KitchenID;constraint:OnDelete:SET NULL;"`
	Address     string
	City        string
	State       string
	Pincode     string
	IsKitchenOn bool
}

// -----------------------------
// Floors
// -----------------------------
type HotelFloor struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	FloorName      string
	HotelID        uuid.UUID `gorm:"type:uuid;not null"`
	Hotel          *Hotel    `gorm:"foreignKey:HotelID;constraint:OnDelete:CASCADE;"`
	FloorNumber    int       `gorm:"not null"`
	Rooms          []Room    `gorm:"foreignKey:FloorID;constraint:OnDelete:CASCADE;"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null"`
}

// -----------------------------
// Rooms
// -----------------------------

// TODO: Create a Separate table for Room Price, gor price history tracking and Dynamic pricing

type Room struct {
	ID             uuid.UUID           `gorm:"type:uuid;primaryKey"`
	FloorID        uuid.UUID           `gorm:"type:uuid;not null"`
	Floor          *HotelFloor         `gorm:"foreignKey:FloorID;constraint:OnDelete:CASCADE;"`
	HotelID        uuid.UUID           `gorm:"type:uuid;not null"`
	OrganizationID uuid.UUID           `gorm:"type:uuid;not null"`
	Organization   models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;"`

	Price       int
	Description string
	Status      hospitalityconstants.RoomStatus `gorm:"type:varchar(50);not null;default:'DISABLED'"`

	RoomName  string `gorm:"size:100;not null"`
	BedSize   string `gorm:"size:50"`
	Amenities string `gorm:"type:text"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// -----------------------------
// Booking
// -----------------------------
type HotelRoomBooking struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	HotelID uuid.UUID `gorm:"type:uuid;not null"`
	RoomID  uuid.UUID `gorm:"type:uuid;not null"`
	Room    *Room     `gorm:"foreignKey:RoomID"`
	Status  string    `gorm:"size:20"` // reserved, checked_in, checked_out, cancelled

	// Guest Details
	PrimaryGuestID uuid.UUID      `gorm:"type:uuid;not null"`
	PrimaryGuest   *HotelGuest    `gorm:"foreignKey:PrimaryGuestID"`
	Guests         []BookingGuest `gorm:"foreignKey:BookingID"`

	Charges  []BookingCharge  `gorm:"foreignKey:BookingID"`
	Payments []BookingPayment `gorm:"foreignKey:BookingID"`

	// Orders & Services (One booking → many)
	Orders   []HotelBookingOrder `gorm:"foreignKey:BookingID"`
	Services []HotelService      `gorm:"foreignKey:BookingID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type BookingCharge struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	// TODO Name change :->  HotelRoomBookingID (Table+Column)
	BookingID        uuid.UUID `gorm:"type:uuid;not null"` // Foreign Key Will
	ChargeType       string    `gorm:"size:50"`            // room_rent, room_service, bar, laundry, spa, minibar
	Description      string    `gorm:"size:200"`
	Amount           float64
	CreatedAt        time.Time
	BookingPaymentID uuid.UUID
	PaymentStatus    string
}

type BookingPayment struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	BookingID uuid.UUID `gorm:"type:uuid;not null"`
	Amount    float64   `gorm:"not null"`
	Mode      string    `gorm:"size:50"`  // cash, upi, card, room_charge
	Status    string    `gorm:"size:20"`  // success, pending, failed
	Reference string    `gorm:"size:100"` // UPI txn ID, Ref no.
	CreatedAt time.Time
}

// -----------------------------
// Guest (Primary)
// -----------------------------
type HotelGuest struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string
	Mobile    string
	Email     string
	Gender    string
	DOB       string
	AadhaarId string
}

// -----------------------------
// Additional Guests
// -----------------------------
type BookingGuest struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	BookingID uuid.UUID `gorm:"type:uuid"`
	Name      string
	Mobile    string
}

// -----------------------------
// Order (One booking -> many orders)
// -----------------------------
type HotelBookingOrder struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	BookingID uuid.UUID `gorm:"type:uuid;not null"`
	OrderType string    `gorm:"size:50"` // restaurant, room_service, bar, laundry, spa
	Status    string    `gorm:"size:20"` // pending, processing, completed, cancelled
	Total     float64

	Items []HotelBookingOrderItem `gorm:"foreignKey:BookingOrderID"`
}

// Items in an order
type HotelBookingOrderItem struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	BookingOrderID uuid.UUID `gorm:"type:uuid;not null"`
	ItemName       string    `gorm:"size:200"`
	ItemID         uuid.UUID
	Quantity       int
	Price          float64
}

// -----------------------------
// Services requested during stay
// -----------------------------
type HotelService struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	BookingID uuid.UUID `gorm:"type:uuid;not null"`
	Time      time.Time

	Items []HotelServiceItem `gorm:"foreignKey:ServiceID"`
}

type HotelServiceItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	ServiceID uuid.UUID `gorm:"type:uuid;not null"`
	Name      string    `gorm:"size:100;not null"`
	Type      string    `gorm:"size:50"` // spa, laundry, housekeeping
	Price     float64   `gorm:"not null"`
}
