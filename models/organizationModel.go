package models

import (
	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model

	Name                string `gorm:"unique;not null"`         // Name must be unique and not null
	OrganizationID      string `gorm:"size:12;unique;not null"` // 12-digit unique number
	Address             string
	Phone               int        `gorm:"not null"`                   // Assuming phone number is required
	Email               string     `gorm:"unique;not null"`            // Email should be unique and not null
	Employees           []Employee `gorm:"foreignKey:OrganizationID"`  // One-to-many relationship
	IsPhysicalVerified  bool       `gorm:"default:false"`              // Boolean, defaults to false
	ISRegistered        bool       `gorm:"default:false"`              // Boolean, defaults to false
	CinNumber           string     `gorm:"unique"`                     // CIN number, if registered, should be unique
	GstNumber           string     `gorm:"unique"`                     // GST number must be unique
	OrganizationRegType string     `gorm:"type:varchar(50)"`           // Registration type (Sole Proprietor, LLP, etc.)
	Banker              string     `gorm:"type:varchar(100)"`          // Name of the banker
	BankAccNumber       string     `gorm:"unique;not null"`            // Bank account number must be unique
	IfscCode            string     `gorm:"type:varchar(20)"`           // IFSC code
	BankBranchName      string     `gorm:"type:varchar(100)"`          // Branch name
	Password            string     `gorm:"not null"`                   // Admin password
	AdminName           string     `gorm:"type:varchar(100);not null"` // Admin name
	Products            []Product  `gorm:"foreignKey:OrganizationID;references:OrganizationID"`
	Service             []Service  `gorm:"foreignKey:OrganizationID;references:OrganizationID"`
}
