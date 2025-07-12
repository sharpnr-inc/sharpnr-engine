package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	TransactionID  string
	OrganizationID string `gorm:"not null"` // Tenant
	ProductID      *uint  // Nullable, if it's a product
	ServiceID      *uint  // Nullable, if it's a service

	Amount          float64 `gorm:"not null"`
	PaymentStatus   string  `gorm:"type:varchar(20);default:'pending'"` // paid, pending, failed
	PaymentMethod   string  // cash, card, UPI, etc.
	TransactionType string  `gorm:"type:varchar(20)"` // sale, refund, etc.
	Notes           string  // Optional comments

	// Optional: Reference to User/Employee who initiated the transaction
	InitiatedByID *uint
}
