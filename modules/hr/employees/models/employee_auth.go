package employee_models

import (
	"time"

	"github.com/google/uuid"
)

type EmployeeAuth struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID `gorm:"uniqueIndex"`

	Username string `gorm:"uniqueIndex"`
	Email    string `gorm:"uniqueIndex"`

	PasswordHash string

	IsActive      bool
	IsProvisioned bool

	CanAccessMobile    bool
	CanAccessDashboard bool
	CanAccessPOS       bool

	LastLoginAt *time.Time
}
