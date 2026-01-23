package employee_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// -------------------------------
// 1️⃣ CORE EMPLOYEE
// -------------------------------
type Employee struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	FirstName string `gorm:"type:varchar(50);not null"`
	LastName  string `gorm:"type:varchar(50)"`
	NickName  string `gorm:"type:varchar(50)"`

	Email       string `gorm:"type:varchar(150);not null;index"`
	PhoneNumber string `gorm:"type:varchar(20)"`

	// Role comes from your constants package
	// Status: ACTIVE, INACTIVE, ON_LEAVE, TERMINATED
	Role   string `gorm:"type:varchar(50);not null"`
	Status string `gorm:"type:varchar(20);default:'ACTIVE'"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Employee) TableName() string {
	return "employee.employee"
}

// -------------------------------
// 2️⃣ VERIFICATION & AUTH
// -------------------------------
type EmployeeVerification struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID uuid.UUID `gorm:"type:uuid;index"`
	Email      string    `gorm:"type:varchar(150);index"`

	Type      string `gorm:"type:varchar(20)"` // EMAIL, SMS
	TokenHash string `gorm:"type:text"`

	ExpiresAt  time.Time
	VerifiedAt *time.Time
	CreatedAt  time.Time
}

func (EmployeeVerification) TableName() string {
	return "employee.employee_verification"
}

type EmployeeSession struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID uuid.UUID `gorm:"type:uuid;index"`

	App        string `gorm:"type:varchar(50)"` // MOBILE, POS, DASHBOARD
	DeviceInfo string `gorm:"type:text"`

	ExpiresAt time.Time
	RevokedAt *time.Time
	CreatedAt time.Time
}

func (EmployeeSession) TableName() string {
	return "employee.employee_session"
}

// -------------------------------
// 3️⃣ REPORTING HIERARCHY
// -------------------------------
type EmployeeReporting struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID uuid.UUID `gorm:"type:uuid;not null;index"`
	ManagerID  uuid.UUID `gorm:"type:uuid;index"`

	Level int `gorm:"type:smallint"` // 1=Manager, 2=Skip-level

	CreatedAt time.Time
}

func (EmployeeReporting) TableName() string {
	return "employee.employee_reporting"
}

// -------------------------------
// 4️⃣ HOSPITALITY OPERATIONAL PROFILE
// -------------------------------
type EmployeeHospitalityProfile struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`

	// OperationalRole: KITCHEN, CONCIERGE, HOUSEKEEPING, WAITER
	OperationalRole string `gorm:"type:varchar(50)"`

	// Capability Flags for RBAC in the UI
	CanTakeOrders  bool `gorm:"default:false"`
	CanPrepareFood bool `gorm:"default:false"`
	CanAssignRooms bool `gorm:"default:false"`
	CanCloseBills  bool `gorm:"default:false"`

	// Preferences for auto-assignment
	DefaultOutlet string `gorm:"type:varchar(50)"` // BAR, RESTAURANT
	ServiceZone   string `gorm:"type:varchar(50)"` // FLOOR_1, SECTION_A

	IsActive bool `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (EmployeeHospitalityProfile) TableName() string {
	return "employee.employee_hospitality_profile"
}
