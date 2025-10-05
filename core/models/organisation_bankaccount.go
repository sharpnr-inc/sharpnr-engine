package models

import "github.com/google/uuid"

type Organization_Bank_Accounts struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;index"`
	BankName       string    `gorm:"size:100;not null"`
	AccountNumber  string    `gorm:"size:50;not null"`
	IFSCCode       string    `gorm:"size:20"`
	BranchName     string    `gorm:"size:20"`
	IsPrimary      bool      `gorm:"default:false"`
	IsSecondary    bool      `gorm:"default:false"`
}
