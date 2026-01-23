package hotel_models

import (
	"time"

	"github.com/google/uuid"
)

// --------------------------------------------------------------------------------------------//
type HotelGuest struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// Ownership & Scope
	OrganizationID uuid.UUID `gorm:"type:uuid;index;not null"`
	PropertyID     uuid.UUID `gorm:"type:uuid;index;not null"`
	HotelID        uuid.UUID `gorm:"type:uuid;index;not null"`
	LegalEntityID  uuid.UUID `gorm:"type:uuid;index;not null"`

	Name      string `gorm:"type:varchar(150);not null"`
	Mobile    string `gorm:"type:varchar(20);index"`
	Email     string `gorm:"type:varchar(150);index"`
	Gender    string `gorm:"type:varchar(10)"`
	DOB       string `gorm:"type:varchar(20)"`       // Stored as string or time.Time depending on preference
	AadhaarId string `gorm:"type:varchar(20);index"` // Sensitive data - usually indexed for search

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (HotelGuest) TableName() string {
	return "hospitality.hotel_guest"
}

type HotelStayGuest struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// Links this specific guest record to a Stay
	HotelStayID  uuid.UUID `gorm:"type:uuid;index;not null"`
	HotelID      uuid.UUID `gorm:"type:uuid;index;not null"`
	HotelGuestID uuid.UUID `gorm:"type:uuid;index;not null"`

	// Snapshot data (as it appeared at the time of check-in)
	Name   string `gorm:"type:varchar(150)"`
	Mobile string `gorm:"type:varchar(20)"`
	Email  string `gorm:"type:varchar(150)"`
	DOB    string `gorm:"type:varchar(20)"`

	CreatedAt time.Time
}

func (HotelStayGuest) TableName() string {
	return "hospitality.hotel_stay_guest"
}
