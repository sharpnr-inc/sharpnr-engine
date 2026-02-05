package employee_models

import (
	"time"

	"github.com/google/uuid"
)

type EmployeeAuth struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// Links to the core Employee identity
	EmployeeID uuid.UUID `gorm:"type:uuid;uniqueIndex;not null"`

	Username string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Email    string `gorm:"type:varchar(150);uniqueIndex;not null"`

	PasswordHash string `gorm:"type:text;not null"`

	IsActive      bool `gorm:"default:true"`
	IsProvisioned bool `gorm:"default:false"` // True once IT setup is complete

	LastLoginAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (EmployeeAuth) TableName() string {
	return "employee.employee_auth"
}
