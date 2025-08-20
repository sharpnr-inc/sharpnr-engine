package models

import (
	"sharpnr-engine/core/constants"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// This is Master Model For the Organisation
type Organization struct {
	ID              uuid.UUID         `gorm:"not null;unique"`
	Name            string            `gorm:"size:100;not null"`
	SharpnrOrgID    string            `gorm:"not null;unique"`
	OrgOprationType constants.OrgType `gorm:"type:varchar(20);not null"`
	OrgDescription  string            `gorm:"type:text"`

	// Address fields
	Address string `gorm:"type:varchar(255);not null"`
	State   string `gorm:"type:varchar(100);not null" `
	City    string `gorm:"type:varchar(100);not null"`
	Pincode string `gorm:"type:varchar(20);not null" `
	Country string `gorm:"type:varchar(100);not null"`

	// Contact fields
	PublicContact     string  `gorm:"type:varchar(100);not null"`
	ContactNumber     string  `gorm:"type:varchar(20);not null"`
	ContactNumber2    *string `gorm:"type:varchar(20)"`
	MasterPhoneNumber string  `gorm:"type:varchar(20)"`
	RootEmail         string  `gorm:"type:varchar(100);not null"`
	SecondaryEmail    *string `gorm:"type:varchar(100)"`

	// Verification and registration
	IsPhysicalVerified  bool    `gorm:"default:false;not null"`
	ISRegistered        bool    `gorm:"default:false;not null"`
	CinNumber           *string `gorm:"type:varchar(50)"`
	GstNumber           *string `gorm:"type:varchar(50)"`
	OrganizationRegType *string `gorm:"type:varchar(100)"`

	// Banking details
	Banker         *string `gorm:"type:varchar(100)"`
	BankAccNumber  string  `gorm:"type:varchar(30);not null"`
	IfscCode       *string `gorm:"type:varchar(20)"`
	BankBranchName *string `gorm:"type:varchar(100)"`
	UPIID          *string `gorm:"type:varchar(100)"`

	// Security
	Password       string `gorm:"type:char(60);not null"`
	MasterPassword string `gorm:"type:char(60);not null"`

	// Timestamps
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relationships
	Products    []Product    `gorm:"foreignKey:OrganizationID"`
	Services    []Service    `gorm:"foreignKey:OrganizationID"`
	Employees   []Employee   `gorm:"foreignKey:OrganizationID"`
	Departments []Department `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// TableName specifies the table name
func (Organization) TableName() string {
	return "organizations"
}
