package employee_models

import (
	"sharpnr-engine/core/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	FirstName string `gorm:"size:50;not null"`
	LastName  string `gorm:"size:50"`
	NickName  string

	Email       string `gorm:"not null"`
	PhoneNumber string `gorm:"size:20"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Role constants.EmployeeRoles `gorm:"type:varchar(50);not null"`

	// Need to Add constants
	Status string
}

type EmployeeVerification struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID `gorm:"index"`
	Email      string    `gorm:"index"`

	Type string // EMAIL

	TokenHash string

	ExpiresAt  time.Time
	VerifiedAt *time.Time

	CreatedAt time.Time
}

type EmployeeReporting struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID `gorm:"type:uuid;not null;index"`
	ManagerID  uuid.UUID `gorm:"type:uuid;index"`

	Level int // 1=manager, 2=skip-level

	CreatedAt time.Time
}

//-------------------------------------------------//

type EmployeeHospitalityProfile struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`

	// Operational role inside hospitality
	OperationalRole string // KITCHEN, CONCIERGE, HOUSEKEEPING, WAITER

	// Capability flags (TEMP – will later map to policy)
	CanTakeOrders  bool
	CanPrepareFood bool
	CanAssignRooms bool
	CanCloseBills  bool

	// Work preferences
	DefaultOutlet string // BAR, RESTAURANT, HOTEL
	ServiceZone   string // FLOOR_1, KITCHEN_A

	IsActive bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type EmployeeSession struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID

	App        string // MOBILE, DASHBOARD, POS
	DeviceInfo string

	ExpiresAt time.Time
	RevokedAt *time.Time
}
