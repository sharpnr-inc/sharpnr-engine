package models

import "time"

type Product struct {
	ProductID      uint   `gorm:"primaryKey;autoIncrement"`
	Name           string `gorm:"not null"`
	Description    *string
	SKU            *string      `gorm:"unique"`
	Price          float64      `gorm:"default:0"`
	OrganizationID string       `gorm:"size:16;not null"` // FK to Organization.SharpnrOrgID
	Organization   Organization `gorm:"foreignKey:OrganizationID;references:SharpnrOrgID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
