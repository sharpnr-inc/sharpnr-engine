package hotelmodels

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// RoomType represents different types of rooms in a hotel
type RoomType struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	Name         string  `gorm:"size:100;not null"`
	Code         string  `gorm:"size:20;not null"`
	Description  string  `gorm:"type:text"`
	BasePrice    float64 `gorm:"type:decimal(10,2);not null"`
	MaxOccupancy int     `gorm:"not null"`

	Amenities pq.StringArray `gorm:"type:text[]"`
	BedTypes  pq.StringArray `gorm:"type:text[]"`
	Size      int            `gorm:"default:0"` // in square meters

	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName specifies the table name
func (RoomType) TableName() string {
	return "hotel_room_types"
}

// Room represents a hotel room
type Room struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`
	RoomNumber     string    `gorm:"size:20;not null"`
	RoomTypeID     uuid.UUID `gorm:"type:uuid;not null"`
	Floor          int       `gorm:"not null"`

	Status         string `gorm:"size:50;default:'available'"` // available, occupied, maintenance, cleaning
	CleaningStatus string `gorm:"size:50;default:'clean'"`     // clean, dirty, in_progress

	Features     pq.StringArray `gorm:"type:text[]"`
	MaxOccupancy int            `gorm:"not null"`
	BasePrice    float64        `gorm:"type:decimal(10,2);not null"`

	Description string `gorm:"type:text"`
	IsActive    bool   `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	// Relationships
	RoomType RoomType `gorm:"foreignKey:RoomTypeID"`
}

// TableName specifies the table name
func (Room) TableName() string {
	return "hotel_rooms"
}

// Default room types for hotels
var DefaultRoomTypes = []struct {
	Name         string
	Code         string
	Description  string
	BasePrice    float64
	MaxOccupancy int
	Amenities    []string
	BedTypes     []string
	Size         int
}{
	{
		Name:         "Standard Room",
		Code:         "STD",
		Description:  "Comfortable standard room with essential amenities",
		BasePrice:    100.00,
		MaxOccupancy: 2,
		Amenities:    []string{"TV", "WiFi", "Air Conditioning", "Mini Bar", "Safe"},
		BedTypes:     []string{"1 Queen"},
		Size:         25,
	},
	{
		Name:         "Deluxe Room",
		Code:         "DLX",
		Description:  "Spacious deluxe room with premium amenities and city view",
		BasePrice:    150.00,
		MaxOccupancy: 3,
		Amenities:    []string{"TV", "WiFi", "Air Conditioning", "Mini Bar", "Safe", "Coffee Maker", "Work Desk", "Bathrobes"},
		BedTypes:     []string{"1 King"},
		Size:         35,
	},
	{
		Name:         "Suite",
		Code:         "STE",
		Description:  "Luxurious suite with separate living area and premium services",
		BasePrice:    300.00,
		MaxOccupancy: 4,
		Amenities:    []string{"TV", "WiFi", "Air Conditioning", "Mini Bar", "Safe", "Coffee Maker", "Work Desk", "Living Room", "Kitchenette", "Jacuzzi"},
		BedTypes:     []string{"1 King", "1 Sofa Bed"},
		Size:         60,
	},
	{
		Name:         "Family Room",
		Code:         "FAM",
		Description:  "Spacious family room with multiple beds",
		BasePrice:    200.00,
		MaxOccupancy: 5,
		Amenities:    []string{"TV", "WiFi", "Air Conditioning", "Mini Bar", "Safe", "Coffee Maker", "Game Console"},
		BedTypes:     []string{"1 Queen", "2 Single"},
		Size:         45,
	},
}
