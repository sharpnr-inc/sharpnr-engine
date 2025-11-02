package hospitalitymodels

import "sharpnr-engine/core/models"

type Hotel struct {
	ID             string              `gorm:"type:uuid;primaryKey"`
	OrganizationID string              `gorm:"type:uuid;not null"`
	Organization   models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;"`

	Name           string `gorm:"size:100"`
	NumberOfFloors int    `gorm:"default:0"`
	NumberOfRooms  int    `gorm:"default:0"`

	Floors []HotelFloor `gorm:"foreignKey:HotelID;constraint:OnDelete:CASCADE;"`
}

type HotelFloor struct {
	ID          string `gorm:"type:uuid;primaryKey"`
	HotelID     string `gorm:"type:uuid;not null"`
	Hotel       Hotel  `gorm:"foreignKey:HotelID;constraint:OnDelete:CASCADE;"`
	FloorNumber int    `gorm:"not null"`
	Rooms       []Room `gorm:"foreignKey:FloorID;constraint:OnDelete:CASCADE;"`
}

type Room struct {
	ID             string              `gorm:"type:uuid;primaryKey"`
	FloorID        string              `gorm:"type:uuid;not null"`
	Floor          HotelFloor          `gorm:"foreignKey:FloorID;constraint:OnDelete:CASCADE;"`
	HotelID        string              `gorm:"type:uuid;not null"`
	OrganizationID string              `gorm:"type:uuid;not null"`
	Organization   models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;"`

	RoomName  string `gorm:"size:100;not null"`
	BedSize   string `gorm:"size:50"`
	Amenities string `gorm:"type:text"`
}
