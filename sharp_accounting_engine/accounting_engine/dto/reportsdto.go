package accountingdtos

import "github.com/google/uuid"

// ── shared line item ──
type ReportLineItemDto struct {
	Code        string  `json:"code"`
	AccountName string  `json:"account_name"`
	Amount      float64 `json:"amount"`
}

// ── P&L section (e.g. "Operating Revenue", "Cost of Goods Sold") ──
type PnLSectionDto struct {
	Label string              `json:"label"`
	Lines []ReportLineItemDto `json:"lines"`
	Total float64             `json:"total"` // sum of lines, computed backend
}

// ── top level P&L ──
type ProfitAndLossReportResponseDto struct {
	// meta
	EntityName  string `json:"entity_name"`
	PeriodFrom  string `json:"period_from"` // "2025-04-01"
	PeriodTo    string `json:"period_to"`   // "2026-03-29"
	GeneratedAt string `json:"generated_at"`

	// revenue
	RevenueSections []PnLSectionDto `json:"revenue_sections"`
	TotalRevenue    float64         `json:"total_revenue"`

	// expenses
	ExpenseSections []PnLSectionDto `json:"expense_sections"`
	TotalExpenses   float64         `json:"total_expenses"`

	// bottom line
	NetProfit    float64 `json:"net_profit"` // can be negative = loss
	IsProfit     bool    `json:"is_profit"`
	ProfitMargin float64 `json:"profit_margin"` // percentage, e.g. 24.5
}

/* ---------------------------------- <***> --------------------------------- */

// AccountMovementRow is an internal DTO used by the repo layer.
// It joins Account + AccountGroup + journal line aggregates in one query.
type AccountMovementRow struct {
	// from Account
	AccountID      uuid.UUID
	AccountName    string
	AccountCode    string
	OpeningBalance float64
	OpeningIsDebit bool

	// from AccountGroup
	GroupName string
	GroupType string // "INCOME", "EXPENSE", "ASSET", "LIABILITY", "EQUITY"

	// aggregated from JournalLine in the period
	TotalDebit  float64
	TotalCredit float64
}
