package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model

	OrganizationID     string `gorm:"not null"` // Foreign key
	ServiceName        string `gorm:"not null"` // Foreign key
	ServicePrice       int
	PricingType        string
	ServiceDescription string
}
