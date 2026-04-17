package accounting_engine_contracts

import (
	"time"

	"github.com/google/uuid"
)

type LegalEntitySuppliersResponseDto struct {
	SuppliersID   uuid.UUID `json:"suppliers_id"`
	SuppliersName string    `json:"suppliers_name"`
	GSTNumber     string    `json:"gst_number"`
	PANNumber     string    `json:"pan_number"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	AddressLine   string    `json:"address_line"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	Pincode       string    `json:"pincode"`
	Country       string    `json:"country"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
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
