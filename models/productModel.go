package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	OrganizationID     string `gorm:"not null"` // Foreign key
	ProductName        string `gorm:"not null"`
	ProductPrice       int
	ProductImage       string
	ProductDescription string
}
