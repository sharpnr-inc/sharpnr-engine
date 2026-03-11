package org_legal_entity_models

import (
	"time"

	"github.com/google/uuid"
)

type LegalEntityAccountingMapping struct {
	ID                        uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationLegalEntityID uuid.UUID `gorm:"type:uuid;index;not null"`
	LegalEntityTaxCategoryID  uuid.UUID `gorm:"type:uuid;index;not null"`
	AccountingBookID          uuid.UUID `gorm:"type:uuid;index;not null"`

	DebitAccountCode  string
	CreditAccountCode string
	TaxAccountCode    string

	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (LegalEntityAccountingMapping) TableName() string {
	return "legal_entity.legal_entity_accounting_mapping"
}
