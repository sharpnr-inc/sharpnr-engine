package hotel_models

import (
	"time"

	"sharpnr-engine/core/models"
	"sharpnr-engine/modules/hospitality/hospitalityconstants"
	kitchen_models "sharpnr-engine/modules/hospitality/kitchen/models"

	"github.com/google/uuid"
)

// -----------------------------
// Hotel master
// -----------------------------
type Hotel struct {
	ID             uuid.UUID            `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID            `gorm:"type:uuid;not null"`
	Organization   *models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;"`

	Name           string `gorm:"size:100"`
	NumberOfFloors int    `gorm:"default:0"`
	NumberOfRooms  int    `gorm:"default:0"`

	Floors []HotelFloor `gorm:"foreignKey:HotelID;constraint:OnDelete:CASCADE;"`

	IsConfigured bool
	CreatedAt    time.Time
	UpdatedAt    time.Time

	DisplayName        string
	LegalName          string
	GstinNumber        string
	CinNumber          string
	FssaiLicenseNumber string

	CheckInTime  string `gorm:"default:10:00"`
	CheckOutTime string `gorm:"default:13:00"`

	// -------------------------
	// Front Desk Overview (1:1)
	// -------------------------
	FrontDeskOverview *HotelFrontDeskOverview `gorm:"foreignKey:HotelID;constraint:OnDelete:CASCADE"`

	IsKitchen   bool                    `gorm:"default:true"`
	KitchenID   *uuid.UUID              `gorm:"type:uuid"`
	Kitchen     *kitchen_models.Kitchen `gorm:"foreignKey:KitchenID;constraint:OnDelete:SET NULL;"`
	IsKitchenOn bool

	Address string
	City    string
	State   string
	Country string
	Pincode string

	TaxRules []HotelTaxRule `gorm:"foreignKey:HotelID;constraint:OnDelete:CASCADE;"`
}

//--------------------------------------------------------------------------------------------//

type HotelTaxRule struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	HotelID uuid.UUID `gorm:"type:uuid;not null;index"`

	RuleName string

	TaxType hospitalityconstants.TaxType `gorm:"type:varchar(50);not null"`

	ConditionType hospitalityconstants.TaxConditionType `gorm:"type:varchar(50);not null"`

	MinValue *float64
	MaxValue *float64

	TaxRate float64

	EffectiveFrom time.Time
	EffectiveTo   *time.Time

	IsActive bool

	HotelTaxComponent []HotelTaxComponent `gorm:"foreignKey:HotelTaxRuleID;constraint:OnDelete:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// --------------------------------------------------------------------------------------------//

type HotelTaxComponent struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	HotelTaxRuleID uuid.UUID `gorm:"type:uuid;not null;index"`

	ComponentType hospitalityconstants.TaxComponentType `gorm:"type:varchar(10);not null"`
	Rate          float64

	CreatedAt time.Time
}

//--------------------------------------------------------------------------------------------//

type HotelFrontDeskOverview struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	HotelID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	Hotel   Hotel     `gorm:"foreignKey:HotelID;constraint:OnDelete:CASCADE"`

	TotalRoom           int `gorm:"not null;default:0"`
	UpcomingReservation int `gorm:"not null;default:0"`
	RoomOccupied        int `gorm:"not null;default:0"`
	AvailableRooms      int `gorm:"not null;default:0"`
	PendingRequest      int `gorm:"not null;default:0"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

//--------------------------------------------------------------------------------------------//

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

//--------------------------------------------------------------------------------------------//

// -----------------------------
// Rooms
// -----------------------------

type Room struct {
	ID             uuid.UUID           `gorm:"type:uuid;primaryKey"`
	FloorID        uuid.UUID           `gorm:"type:uuid;not null"`
	Floor          *HotelFloor         `gorm:"foreignKey:FloorID;constraint:OnDelete:CASCADE;"`
	HotelID        uuid.UUID           `gorm:"type:uuid;not null"`
	OrganizationID uuid.UUID           `gorm:"type:uuid;not null"`
	Organization   models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;"`

	Prices []HotelRoomPrice `gorm:"foreignKey:RoomID"`

	Description      string
	Status           hospitalityconstants.HotelRoomStatusType `gorm:"type:varchar(50);not null;default:'DISABLED'"`
	QRCodes          []HotelRoomQRCode                        `gorm:"foreignKey:RoomID"`
	RoomName         string                                   `gorm:"size:100;not null"`
	BedSize          string                                   `gorm:"size:50"`
	Amenities        string                                   `gorm:"type:text"`
	CurrentBookingID *uuid.UUID                               `gorm:"type:uuid"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

//--------------------------------------------------------------------------------------------//

// -----------------------------
// Hotel Room Price
// -----------------------------
type HotelRoomPrice struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	RoomID uuid.UUID `gorm:"type:uuid;not null;index"`
	Room   Room      `gorm:"foreignKey:RoomID;constraint:OnDelete:CASCADE"`

	HotelID uuid.UUID `gorm:"type:uuid;not null;index"`
	Hotel   Hotel     `gorm:"foreignKey:HotelID;constraint:OnDelete:CASCADE"`

	Price int `gorm:"not null"`

	// Optional but HIGHLY recommended
	EffectiveFrom time.Time
	EffectiveTo   *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

//--------------------------------------------------------------------------------------------//

// -----------------------------
// Hotel Room QR code
// -----------------------------
type HotelRoomQRCode struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	RoomID uuid.UUID `gorm:"type:uuid;not null"`
	Room   Room      `gorm:"foreignKey:RoomID;constraint:OnDelete:CASCADE;"`

	// what this QR is for
	Purpose string `gorm:"size:50;not null"`
	// ex: ROOM_SERVICE, FEEDBACK, CHECKOUT

	Token string `gorm:"size:255;uniqueIndex;not null"`

	IsActive bool `gorm:"default:true"`

	ExpiresAt *time.Time

	ScanCount     int `gorm:"default:0"`
	LastScannedAt *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

//--------------------------------------------------------------------------------------------//

// -----------------------------
// Booking
// -----------------------------
type HotelRoomBooking struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	HotelID uuid.UUID `gorm:"type:uuid;not null"`

	RoomID uuid.UUID `gorm:"type:uuid;not null"`

	Room *Room `gorm:"foreignKey:RoomID"`

	Status string `gorm:"size:20"` // reserved, checked_in, checked_out, cancelled

	// Guest Details
	PrimaryGuestID uuid.UUID      `gorm:"type:uuid;not null"`
	PrimaryGuest   *HotelGuest    `gorm:"foreignKey:PrimaryGuestID;references:ID"`
	Guests         []BookingGuest `gorm:"foreignKey:HotelRoomBookingID"`

	CheckInDate  time.Time
	CheckOutDate time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

//--------------------------------------------------------------------------------------------//

type HotelBookingCharge struct {
	ID                 uuid.UUID                            `gorm:"type:uuid;primaryKey"`
	HotelRoomBookingID uuid.UUID                            `gorm:"type:uuid;not null"` // Foreign Key Will
	ChargeType         hospitalityconstants.HotelChargeType `gorm:"size:50"`            // room_rent, room_service, bar, laundry, spa, minibar

	ReferenceID *uuid.UUID

	Description string  `gorm:"size:200"`
	Quantity    float64 `gorm:"default:1"`
	UnitRate    float64 `gorm:"default:0"`
	Amount      float64

	CreatedAt time.Time
}

//--------------------------------------------------------------------------------------------//

type HotelBill struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	HotelRoomBookingID uuid.UUID `gorm:"uniqueIndex"`
	HotelID            uuid.UUID

	BaseAmount  float64 `gorm:"not null"`
	TaxAmount   float64 `gorm:"not null"`
	TotalAmount float64 `gorm:"not null"`

	Status string // DRAFT, FINALIZED, CANCELLED

	CreatedAt   time.Time
	FinalizedAt *time.Time
}

//--------------------------------------------------------------------------------------------//

type HotelBillTax struct {
	ID uuid.UUID

	HotelBillID uuid.UUID

	ComponentType string
	Rate          float64
	Amount        float64
}

//--------------------------------------------------------------------------------------------//

type HotelRoomBookingPayment struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey"`
	HotelID            uuid.UUID
	HotelRoomBookingID uuid.UUID                          `gorm:"type:uuid;not null"`
	Amount             float64                            `gorm:"not null"`
	Mode               hospitalityconstants.PaymentMethod `gorm:"size:50"`  // cash, upi, card, room_charge
	Status             hospitalityconstants.PaymentStatus `gorm:"size:20"`  // success, pending, failed
	Reference          string                             `gorm:"size:100"` // UPI txn ID, Ref no.
	CreatedAt          time.Time
}

//--------------------------------------------------------------------------------------------//

// -----------------------------
// Guest (Primary)
// -----------------------------
type HotelGuest struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	HotelId uuid.UUID `gorm:"type:uuid"`
	OrgId   uuid.UUID `gorm:"type:uuid"`

	Name      string
	Mobile    string
	Email     string
	Gender    string
	DOB       string
	AadhaarId string
}

//--------------------------------------------------------------------------------------------//

// -----------------------------
// Additional Guests
// -----------------------------
type BookingGuest struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey"`
	HotelID            uuid.UUID
	HotelGuestID       uuid.UUID
	HotelRoomBookingID uuid.UUID `gorm:"type:uuid"`
	Name               string
	Mobile             string
	Email              string
	DOB                string
}

//--------------------------------------------------------------------------------------------//

// -----------------------------
// Services requested during stay
// -----------------------------
type HotelService struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	BookingID uuid.UUID `gorm:"type:uuid;not null"`
	Time      time.Time

	Items []HotelServiceItem `gorm:"foreignKey:ServiceID"`
}

//--------------------------------------------------------------------------------------------//

type HotelServiceItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	ServiceID uuid.UUID `gorm:"type:uuid;not null"`
	Name      string    `gorm:"size:100;not null"`
	Type      string    `gorm:"size:50"` // spa, laundry, housekeeping
	Price     float64   `gorm:"not null"`
}

//--------------------------------------------------------------------------------------------//
