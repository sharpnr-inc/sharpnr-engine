package accounting_engine_contracts

import "github.com/google/uuid"

type LegalEntitySuppliersResponseDto struct {
	SuppliersID   uuid.UUID
	SuppliersName string
}

type AddNewSuppliersForLegalEntityRequestDto struct {
	Name        string `gorm:"not null"`
	GSTNumber   string
	PANNumber   string
	Phone       string
	Email       string
	AddressLine string
	City        string
	State       string
	Pincode     string
	Country     string
}
