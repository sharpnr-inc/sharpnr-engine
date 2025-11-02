package hospitalitymodels

import (
	"time"

	"sharpnr-engine/core/models"
)

type Bar struct {
	ID             string              `gorm:"type:uuid;primaryKey" json:"id"`
	Name           string              `gorm:"size:100;not null" json:"name"`
	OrganizationID string              `gorm:"type:uuid;not null" json:"organization_id"`
	Organization   models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;" json:"organization"`

	OpeningTime time.Time `json:"opening_time"`
	ClosingTime time.Time `json:"closing_time"`

	Tables []BarTable `gorm:"foreignKey:BarID;constraint:OnDelete:CASCADE;" json:"tables"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BarTable represents an individual table inside a bar.
type BarTable struct {
	ID    string `gorm:"type:uuid;primaryKey" json:"id"`
	BarID string `gorm:"type:uuid;not null" json:"bar_id"`
	Bar   Bar    `gorm:"foreignKey:BarID;constraint:OnDelete:CASCADE;" json:"bar"`

	TableName string `gorm:"size:50;not null" json:"table_name"`
	Seats     int    `gorm:"default:4" json:"seats"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
