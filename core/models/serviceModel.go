package models

import (
	"time"

	"gorm.io/gorm"
)

type Service struct {
	ID          string  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string  `gorm:"type:varchar(100);not null"`
	Description string  `gorm:"type:text"`
	Duration    int     `gorm:"not null"` // duration in minutes
	Price       float64 `gorm:"not null"`

	OrganizationID string       `gorm:"type:uuid;not null;index"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
