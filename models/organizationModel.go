package models

import (
	"github.com/google/uuid"
)

type Organization struct {
	Name           string    `gorm:"not null"`                // Name must not be null
	SharpnrOrgID   string    `gorm:"size:16;unique;not null"` // Public-facing ID (e.g. SHR000000000001)
	SharpnrOrgUUID uuid.UUID `gorm:"type:uuid;primaryKey"`    // Internal UUID for linking and security

	// Type of organization: Hotel, Salon, Retail, etc.
	OrgOprationType string `gorm:"type:varchar(50)"`

	CreatedAt int64 `gorm:"autoCreateTime"` // Timestamp of creation
	UpdatedAt int64 `gorm:"autoUpdateTime"` // Timestamp of last update

	// Address details
	Address string `gorm:"type:varchar(255);not null"`
	State   string `gorm:"type:varchar(100);not null"`
	City    string `gorm:"type:varchar(100);not null"`
	Pincode string `gorm:"type:varchar(20);not null"`
	Country string `gorm:"type:varchar(100);not null"`

	// Contact details
	PublicContact  int     `gorm:"type:varchar(15);not null"`     // Assuming it's a contact number
	ContactNumber  string  `gorm:"type:varchar(15);not null"`     // Primary contact
	ContactNumber2 *string `gorm:"type:varchar(15);default:null"` // Optional
	MasterContact  *string `gorm:"type:varchar(15);default:null"` // Optional
	RootEmail      string  `gorm:"unique;not null"`
	SecondaryEmail *string `gorm:"type:varchar(100);default:null"`

	// Verification & registration
	IsPhysicalVerified  bool    `gorm:"default:false"`
	ISRegistered        bool    `gorm:"default:false"`
	CinNumber           *string `gorm:"type:varchar(30);unique"`
	GstNumber           *string `gorm:"type:varchar(20);unique"`
	OrganizationRegType *string `gorm:"type:varchar(50)"`

	// Banking
	Banker         *string `gorm:"type:varchar(100)"`
	BankAccNumber  string  `gorm:"type:varchar(30);unique;not null"`
	IfscCode       *string `gorm:"type:varchar(20)"`
	BankBranchName *string `gorm:"type:varchar(100)"`
	UPIID          *string `gorm:"type:varchar(100);unique"`

	// Root details
	Password       string `gorm:"not null"`
	MasterPassword string `gorm:"not null"`

	// Organisation System Details
	// TerminalID *string `gorm:"type:varchar(50);unique"`       // Unique Terminal ID for organization
	// AccessKey  *string `gorm:"type:varchar(50);unique"`       // Unique Access Key for organization
	// OrgOS      *string `gorm:"type:varchar(50);default:null"` // Operating System of the organization (e.g. Android, iOS, Web)

	// Relationships (fixed references to SharpnrOrgUUID)
	Products  []Product  `gorm:"foreignKey:OrganizationID;references:SharpnrOrgUUID"`
	Service   []Service  `gorm:"foreignKey:OrganizationID;references:SharpnrOrgUUID"`
	Employees []Employee `gorm:"foreignKey:OrganizationID;references:SharpnrOrgUUID"`
}
