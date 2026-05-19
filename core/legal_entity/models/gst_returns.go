package legal_entity_models

import (
	"time"

	"github.com/google/uuid"
)

// LegalEntityGSTReturn stores the summary of GST returns filed per period.
// Used for three-way reconciliation: your invoice data vs this table vs
// GSTR-2B auto-populated data from the GSTN portal.
type LegalEntityGSTReturn struct {
	ID                           uuid.UUID `gorm:"type:uuid;primaryKey"`
	LegalEntityID                uuid.UUID `gorm:"type:uuid;not null;index"`
	LegalEntityTaxRegistrationID uuid.UUID `gorm:"type:uuid;not null;index"`

	// ReturnType: GSTR1 | GSTR3B | GSTR9 | GSTR9C | CMP08
	ReturnType  string `gorm:"size:10;not null"`
	PeriodStart time.Time
	PeriodEnd   time.Time

	// Status: DRAFT | FILED | FILED_WITH_LATE_FEE | REVISED
	Status string `gorm:"size:25;not null;default:'DRAFT'"`

	ARN          *string `gorm:"size:50"` // Acknowledgment Reference Number from GSTN
	FiledAt      *time.Time
	LateFeePaid  float64 `gorm:"default:0"`
	InterestPaid float64 `gorm:"default:0"`

	// TaxableTurnover / TotalTaxLiability are the figures as filed —
	// used to cross-check against your internal ledger totals.
	TaxableTurnover   float64 `gorm:"default:0"`
	TotalTaxLiability float64 `gorm:"default:0"`
	ITCClaimed        float64 `gorm:"default:0"`
	TaxPaid           float64 `gorm:"default:0"`

	FiledByUserID *uuid.UUID `gorm:"type:uuid"`
	DocumentURL   *string    `gorm:"size:512"` // filed JSON / PDF from GSTN

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (LegalEntityGSTReturn) TableName() string {
	return "legal_entity.legal_entity_gst_return"
}
