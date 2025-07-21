package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Terminal struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	Name        string `gorm:"size:100;not null"`
	Description string `gorm:"type:text"`

	TerminalOperations []TerminalOperations `gorm:"type:text[]"`
	Permissions        pq.StringArray       `gorm:"type:text[]"`

	IsActive  bool `gorm:"default:true"`
	IsDefault bool `gorm:"default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relationships
	Organization Organization `gorm:"foreignKey:OrganizationID"`
	Employees    []Employee   `gorm:"many2many:terminal_employees;"` // Many-to-many
}

type TerminalEmployee struct {
	TerminalID uuid.UUID `gorm:"type:uuid;primary_key"`
	EmployeeID uuid.UUID `gorm:"type:uuid;primary_key"`
	CanAccess  bool      `gorm:"default:true"`
	CreatedAt  time.Time
}

type TerminalOperations struct {
	TerminalID  uuid.UUID
	OperationID uuid.UUID
	CreatedAt   time.Time
}
