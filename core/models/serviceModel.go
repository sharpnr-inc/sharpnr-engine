package models

import (
	"time"
)

type Service struct {
	ServiceID      uint   `gorm:"primaryKey;autoIncrement"`
	Name           string `gorm:"not null"`
	Description    *string
	Fee            float64      `gorm:"default:0"`
	OrganizationID string       `gorm:"size:16;not null"` // FK to Organization.SharpnrOrgID
	Organization   Organization `gorm:"foreignKey:OrganizationID;references:SharpnrOrgID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
