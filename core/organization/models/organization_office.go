package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrganizationOffice struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;index;not null"`

	Code string `gorm:"type:varchar(30)"`           // HQ, BLR-01
	Name string `gorm:"type:varchar(100);not null"` // Head Office

	IsHQ     bool `gorm:"default:false"`
	IsActive bool `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type OrganizationOfficeAddress struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OfficeID uuid.UUID `gorm:"type:uuid;index;not null"`

	Type string `gorm:"type:varchar(30)"`
	// REGISTERED | COMMUNICATION | BILLING

	AddressLine string `gorm:"type:varchar(255);not null"`
	City        string `gorm:"type:varchar(100);not null"`
	State       string `gorm:"type:varchar(100);not null"`
	Pincode     string `gorm:"type:varchar(20);not null"`
	Country     string `gorm:"type:varchar(100);not null"`

	IsPrimary bool `gorm:"default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrganizationOfficeContact struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OfficeID uuid.UUID `gorm:"type:uuid;index"`

	Role string // HR, FINANCE, SUPPORT

	Email *string
	Phone *string

	CreatedAt time.Time
}

type OrganizationOfficeCompliance struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OfficeID uuid.UUID `gorm:"type:uuid;uniqueIndex"`

	GSTIN string
	CIN   string

	EffectiveFrom time.Time
	EffectiveTo   *time.Time
}
