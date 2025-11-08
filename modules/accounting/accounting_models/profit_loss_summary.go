package accounting_models

import "github.com/google/uuid"

type ProfitLossSummary struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	BookID       uuid.UUID
	FiscalYear   string
	TotalIncome  float64
	TotalExpense float64
	NetProfit    float64
}
