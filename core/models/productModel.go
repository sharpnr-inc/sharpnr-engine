package models

import (
	"time"

	"gorm.io/gorm"
)

// Every Organisation Will gonna have either products or service
// Every Organisation will gonna products or service
// Or Gonna have only
// this is Model for the Products
type Product struct {
	ID          string  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string  `gorm:"type:varchar(100);not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"not null"`
	SKU         string  `gorm:"type:varchar(50);uniqueIndex"`

	OrganizationID string       `gorm:"type:uuid;not null;index"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
