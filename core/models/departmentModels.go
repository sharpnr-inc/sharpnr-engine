package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Department struct {
	ID             uuid.UUID `gorm:"primaryKey;"`
	OrganizationID string    `gorm:"type:varchar(256);not null;index"`

	Name        string `gorm:"size:100;not null"`
	Description string `gorm:"type:text"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relationships - FIXED
	Organization Organization `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Employees    []Employee   `gorm:"foreignKey:DepartmentID;references:ID" json:"employees"`
}

// TableName specifies the table name
func (Department) TableName() string {
	return "departments"
}
