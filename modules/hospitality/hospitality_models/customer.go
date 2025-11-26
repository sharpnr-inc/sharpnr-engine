package hospitalitymodels

import (
	"github.com/google/uuid"
)

type HospitalityCustomerAndGuest struct {
	ID           uuid.UUID
	Name         string
	DOB          string
	MobileNumber string
	EmailAddress string
	AadhaarID    string
	PANCardID    string
	// Purchases
	// ADD (correct, without renaming tables)
	BarPurchases        []HospitalityCustomerBarPurchase        `gorm:"foreignKey:HospitalityCustomerAndGuestID"`
	HotelPurchases      []HospitalityCustomerRestaurantPurchase `gorm:"foreignKey:HospitalityCustomerAndGuestID"`
	RestaurantPurchases []HospitalityCustomerRestaurantPurchase `gorm:"foreignKey:HospitalityCustomerAndGuestID"`
	// Meta
	CreatedAt string
	UpdatedAt string
}

type HospitalityCustomerBarPurchase struct {
	ID                            uuid.UUID
	HospitalityCustomerAndGuestID uuid.UUID
	InvoiceID                     uuid.UUID
}

type HospitalityCustomerRestaurantPurchase struct {
	ID                            uuid.UUID
	HospitalityCustomerAndGuestID uuid.UUID
	InvoiceID                     uuid.UUID
}

type HospitalityGuestRestaurantPurchase struct {
	ID                            uuid.UUID
	HospitalityCustomerAndGuestID uuid.UUID
	InvoiceID                     uuid.UUID
}
