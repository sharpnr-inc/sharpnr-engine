package accounting_models

import "github.com/google/uuid"

type LedgerBalance struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	AccountID   uuid.UUID
	BookID      uuid.UUID
	FiscalYear  string
	OpeningBal  float64
	DebitTotal  float64
	CreditTotal float64
	ClosingBal  float64
}
