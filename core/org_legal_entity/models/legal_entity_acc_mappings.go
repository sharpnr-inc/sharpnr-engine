package org_legal_entity_models

import "github.com/google/uuid"

type LegalEntityAccountingMapping struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationLegalEntityID uuid.UUID `gorm:"type:uuid;index;not null"`
	TaxCategoryID             uuid.UUID `gorm:"type:uuid;index;not null"`

	AccountingBookID uuid.UUID `gorm:"type:uuid;index;not null"`

	DebitAccountCode  string
	CreditAccountCode string
	TaxAccountCode    string

	IsActive bool
}
