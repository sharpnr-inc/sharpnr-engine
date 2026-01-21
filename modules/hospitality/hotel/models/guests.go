package hotel_models

import "github.com/google/uuid"

// --------------------------------------------------------------------------------------------//
type HotelGuest struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	// Ownership & scope
	OrganizationID uuid.UUID `gorm:"type:uuid;index;not null"`
	PropertyID     uuid.UUID `gorm:"type:uuid;index;not null"`
	HotelID        uuid.UUID `gorm:"type:uuid;index;not null"`
	LegalEntityID  uuid.UUID `gorm:"type:uuid;index;not null"`

	Name      string
	Mobile    string
	Email     string
	Gender    string
	DOB       string
	AadhaarId string
}

type HotelStayGuest struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	HotelID      uuid.UUID
	HotelGuestID uuid.UUID

	Name   string
	Mobile string
	Email  string
	DOB    string
}
