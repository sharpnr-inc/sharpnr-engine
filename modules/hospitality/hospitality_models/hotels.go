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
	GstinNumber        string
	FssaiLicenseNumber string

	IsKitchen bool       `gorm:"default:true" json:"is_kitchen"`
	KitchenID *uuid.UUID `gorm:"type:uuid" json:"kitchen_id"`
	Kitchen   *Kitchen   `gorm:"foreignKey:KitchenID;constraint:OnDelete:SET NULL;" json:"kitchen"`
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
	HotelID        string              `gorm:"type:uuid;not null"`
	OrganizationID string              `gorm:"type:uuid;not null"`
	Organization   models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;"`
	Price          int
	Description    string
	Status         hospitalityconstants.RoomStatus `gorm:"type:varchar(50);not null;default:'DISABLED'"`

	RoomName  string `gorm:"size:100;not null"`
	BedSize   string `gorm:"size:50"`
	Amenities string `gorm:"type:text"`
}
