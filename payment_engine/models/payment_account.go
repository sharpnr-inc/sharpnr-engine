package payment_engine_models

import (
	"time"

	"github.com/google/uuid"
)

type PaymentAccount struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	TenantID      uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"` // maps to OrgID in monolith
	AccountRef    string    `gorm:"type:varchar(100);not null"`     // human ref e.g. "acme-hotels"
	LegalEntityID uuid.UUID `gorm:"type:uuid;not null;index"`
	CurrencyCode  string    `gorm:"type:char(3);not null;default:'INR'"`
	IsActive      bool      `gorm:"default:true"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (PaymentAccount) TableName() string { return "payments.payment_accounts" }
