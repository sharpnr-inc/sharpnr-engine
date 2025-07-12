package models

import (
	"time"
	"github.com/google/uuid"
)

type Product struct {
	ProductID      uint      `gorm:"primaryKey;autoIncrement"`
	Name           string    `gorm:"not null"`
	Description    *string
	SKU            *string   `gorm:"unique"`
	Price          float64   `gorm:"default:0"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null"` // FK to Organization.SharpnrOrgUUID
	Organization   Organization `gorm:"foreignKey:OrganizationID;references:SharpnrOrgUUID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}