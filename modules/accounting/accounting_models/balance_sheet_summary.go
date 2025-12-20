package accounting_models

import "github.com/google/uuid"

type BalanceSheetSummary struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	BookID      uuid.UUID
	FiscalYear  string
	TotalAssets float64
	TotalLiab   float64
	Capital     float64
}
