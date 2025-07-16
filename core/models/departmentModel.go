// core/models/department.go
package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/sharpnr-inc/sharpnr-engine/core/constants"
	"gorm.io/gorm"
)

// Department represents organizational departments
type Department struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	Name        string                       `gorm:"size:100;not null"`
	Code        string                       `gorm:"size:20;not null"`
	Type        constants.BaseDepartmentType `gorm:"size:50;not null"` // Use the enum type
	Description string                       `gorm:"type:text"`

	IsDefault bool `gorm:"default:false"` // Default departments can't be deleted
	IsActive  bool `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relationships
	Organization Organization `gorm:"foreignKey:OrganizationID"`
	Terminals    []Terminal
	Employees    []Employee `gorm:"default:null"`
}

// TableName specifies the table name
func (Department) TableName() string {
	return "departments"
}

// BeforeCreate hook
func (d *Department) BeforeCreate(tx *gorm.DB) error {
	if d.ID == uuid.Nil {
		d.ID = uuid.New()
	}
	return nil
}
