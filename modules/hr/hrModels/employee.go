package sharphr_models

import (
	"sharpnr-engine/core/constants"
	"sharpnr-engine/core/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	FirstName   string    `gorm:"size:50;not null"`
	NickName    string
	LastName    string `gorm:"size:50"`
	Email       string `gorm:"not null"`
	PhoneNumber string `gorm:"size:20"`

	ReportingMangerID string

	SkipLevelMangerID string

	OrganizationID uuid.UUID           `gorm:"type:uuid;not null;index"`
	Organization   models.Organization `gorm:"foreignKey:OrganizationID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt          `gorm:"index"`
	Role      constants.EmployeeRoles `gorm:"type:varchar(50);not null"`

	IsEmailVerified  bool
	Address          string
	EmergencyContact string
	IsIDVerified     string
	Salary           string

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

type EmployeeAuth struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID `gorm:"uniqueIndex"`

	Username string `gorm:"uniqueIndex"`
	Email    string `gorm:"uniqueIndex"`

	PasswordHash string

	IsActive    bool
	LastLoginAt *time.Time
}

type EmployeeSession struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID

	App        string // MOBILE, DASHBOARD, POS
	DeviceInfo string

	ExpiresAt time.Time
	RevokedAt *time.Time
}

type EmployeeProfile struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	EmployeeID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`

	Gender      string
	DateOfBirth *time.Time

	Address          string
	EmergencyContact string
	BloodGroup       string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type EmployeeBankAccount struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID `gorm:"type:uuid;uniqueIndex"`

	BankName      string
	AccountNumber string
	IFSC          string
	UPI           string

	IsPrimary bool
}

type EmployeeDocument struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID

	Type string // ID_PROOF, CONTRACT
	URL  string

	Verified bool
}

type EmployeeReporting struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID `gorm:"type:uuid;not null;index"`
	ManagerID  uuid.UUID `gorm:"type:uuid;index"`

	Level int // 1=manager, 2=skip-level

	CreatedAt time.Time
}

type EmployeeEmployment struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID `gorm:"type:uuid;uniqueIndex"`

	EmploymentType string // FULL_TIME, PART_TIME, CONTRACT
	Designation    string

	ConfirmationDate *time.Time

	CreatedAt time.Time
}

type EmployeeCompensation struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID uuid.UUID `gorm:"type:uuid;index"`

	BaseSalary float64
	Currency   string

	EffectiveFrom time.Time
	EffectiveTo   *time.Time
}

type EmployeeSalaryComponent struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	CompensationID uuid.UUID

	Component string // BASIC, HRA, BONUS, PF
	Amount    float64
}

//-------------------------------------------------//

/*********************************************
 * All the Tables Related to Employee Leaves *
 *********************************************/

type EmployeeLeaveType struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID
	Name           string // Casual, Sick, Paid
}

type EmployeeLeaveBalance struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID  uuid.UUID
	LeaveTypeID uuid.UUID

	Balance float64
}

type EmployeeLeaveRequest struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	EmployeeID  uuid.UUID
	LeaveTypeID uuid.UUID

	FromDate time.Time
	ToDate   time.Time

	Status string // PENDING, APPROVED, REJECTED

	ApprovedBy *uuid.UUID // Manager
	ApprovedAt *time.Time
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
