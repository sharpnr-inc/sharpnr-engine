package models

import "github.com/google/uuid"

// OrganizationBankAccount represents an organization's bank account
type OrganizationBankAccount struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BankName       string    `gorm:"size:100;not null"`
	AccountNumber  string    `gorm:"size:50;not null"`
	IFSCCode       string    `gorm:"size:20"`
	BranchName     string    `gorm:"size:100"`
	IsPrimary      bool      `gorm:"default:false"`
	IsSecondary    bool      `gorm:"default:false"`
}
