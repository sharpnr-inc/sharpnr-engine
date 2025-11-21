package hospitalitymodels

import (
	"sharpnr-engine/core/models"
	"sharpnr-engine/modules/hospitality/hospitalityconstants"
	"time"

	"github.com/google/uuid"
)

type Hotel struct {
	ID             uuid.UUID           `gorm:"primaryKey"`
	OrganizationID uuid.UUID           `gorm:"not null"`
	Organization   models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;"`

	Name           string `gorm:"size:100"`
	NumberOfFloors int    `gorm:"default:0"`
	NumberOfRooms  int    `gorm:"default:0"`

	Floors []HotelFloor `gorm:"foreignKey:HotelID;constraint:OnDelete:CASCADE;"`

	// Flags
	IsConfigured       bool `json:"is_configured"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DisplayName        string
	LegalName          string
	GstinNumber        string
	FssaiLicenseNumber string

	IsKitchen bool       `gorm:"default:true" json:"is_kitchen"`
	KitchenID *uuid.UUID `gorm:"type:uuid" json:"kitchen_id"`
	Kitchen   *Kitchen   `gorm:"foreignKey:KitchenID;constraint:OnDelete:SET NULL;" json:"kitchen"`

	Address      string
	City         string
	State        string
	Pincode      string
	IsKitchenOnn bool
}

type HotelFloor struct {
	ID             string `gorm:"type:uuid;primaryKey"`
	FloorName      string
	HotelID        string    `gorm:"type:uuid;not null"`
	Hotel          Hotel     `gorm:"foreignKey:HotelID;constraint:OnDelete:CASCADE;"`
	FloorNumber    int       `gorm:"not null"`
	Rooms          []Room    `gorm:"foreignKey:FloorID;constraint:OnDelete:CASCADE;"`
	OrganizationID uuid.UUID `gorm:"not null"`
}

type Room struct {
	ID             string              `gorm:"type:uuid;primaryKey"`
	FloorID        string              `gorm:"type:uuid;not null"`
	Floor          HotelFloor          `gorm:"foreignKey:FloorID;constraint:OnDelete:CASCADE;"`
	HotelID        uuid.UUID           `gorm:"type:uuid;not null"`
	OrganizationID uuid.UUID           `gorm:"type:uuid;not null"`
	Organization   models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;"`
	Price          int
	Description    string
	Status         hospitalityconstants.RoomStatus `gorm:"type:varchar(50);not null;default:'DISABLED'"`

	RoomName  string `gorm:"size:100;not null"`
	BedSize   string `gorm:"size:50"`
	Amenities string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type HotelBooking struct {
	Id      uuid.UUID
	HotelId uuid.UUID
	RoomId  string
	Room    Room
	Guest   HotelGuests
	Status  string `gorm:"size:20"` // reserved, checked_in, checked_out, cancelled
	Order   HotelBookingOrder
}

type HotelBookingOrder struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	BookingID uuid.UUID
	OrderType string `gorm:"size:50"` // restaurant, room_service, bar, laundry, spa
	Status    string `gorm:"size:20"` // pending, processing, completed, cancelled
	Total     float64
}

type HotelBookingOrderItem struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	BookingOrderID uuid.UUID
	ItemName       string `gorm:"size:200"`
	ItemId         uuid.UUID
	Quantity       int     `gorm:"not null"`
	Price          float64 `gorm:"not null"`
	BookingOrder   HotelBookingOrder
}

type HotelService struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name  string    `gorm:"size:100;not null"`
	Type  string    `gorm:"size:50"` // spa, laundry, housekeeping
	Price float64   `gorm:"not null"`
}

type HotelGuests struct {
	ID     uuid.UUID
	Name   string
	Mobile string
	Email  string
}
