package accounting_service

import (
	"context"
	"fmt"
	"math"
	accountingdtos "sharpnr-engine/sharp_accounting_engine/accounting_engine/dto"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (s *accountingServices) GenerateProfitAndLossReport(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
	periodFrom, periodTo string,
) (accountingdtos.ProfitAndLossReportResponseDto, error) {

	empty := accountingdtos.ProfitAndLossReportResponseDto{}

	// ── 1. parse dates ──
	from, err := time.Parse("2006-01-02", periodFrom)
	if err != nil {
		return empty, fmt.Errorf("invalid period_from: %w", err)
	}
	to, err := time.Parse("2006-01-02", periodTo)
	if err != nil {
		return empty, fmt.Errorf("invalid period_to: %w", err)
	}
	// make `to` inclusive (end of day)
	to = to.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	// ── 2. get accounting book ──
	book, err := s.accRepo.GetAccountingBookByLegalEntityID(ctx, legalEntityID)
	if err != nil {
		return empty, fmt.Errorf("accounting book not found: %w", err)
	}

	// ── 3. fetch all account movements for the period ──
	rows, err := s.accRepo.GetAccountMovementsForPeriod(ctx, book.ID, from, to)
	if err != nil {
		return empty, fmt.Errorf("GetAccountMovementsForPeriod: %w", err)
	}

	// ── 4. separate into INCOME and EXPENSE buckets ──
	// key = group_name, value = lines
	revMap := make(map[string][]accountingdtos.ReportLineItemDto)
	expMap := make(map[string][]accountingdtos.ReportLineItemDto)

	for _, row := range rows {
		switch strings.ToUpper(row.GroupType) {

		case "INCOME":
			// income is Cr-normal: grows when credited
			netAmount := row.OpeningBalance + row.TotalCredit - row.TotalDebit
			if !row.OpeningIsDebit {
				netAmount = row.OpeningBalance + row.TotalCredit - row.TotalDebit
			}
			if netAmount <= 0 {
				continue // skip zero-balance income accounts
			}
			revMap[row.GroupName] = append(revMap[row.GroupName], accountingdtos.ReportLineItemDto{
				Code:        row.AccountCode,
				AccountName: row.AccountName,
				Amount:      netAmount,
			})

		case "EXPENSE":
			// expense is Dr-normal: grows when debited
			netAmount := row.OpeningBalance + row.TotalDebit - row.TotalCredit
			if netAmount <= 0 {
				continue // skip zero-balance expense accounts
			}
			expMap[row.GroupName] = append(expMap[row.GroupName], accountingdtos.ReportLineItemDto{
				Code:        row.AccountCode,
				AccountName: row.AccountName,
				Amount:      netAmount,
			})
		}
	}

	// ── 5. build sections ──
	revSections, totalRevenue := buildPnLSections(revMap)
	expSections, totalExpenses := buildPnLSections(expMap)

	// ── 6. compute bottom line ──
	netProfit := totalRevenue - totalExpenses
	isProfit := netProfit >= 0

	var margin float64
	if totalRevenue > 0 {
		margin = math.Round((netProfit/totalRevenue)*10000) / 100 // 2 decimal places
	}

	return accountingdtos.ProfitAndLossReportResponseDto{
		EntityName:      book.Name, // or fetch from legal entity
		PeriodFrom:      periodFrom,
		PeriodTo:        periodTo,
		GeneratedAt:     time.Now().Format(time.RFC3339),
		RevenueSections: revSections,
		TotalRevenue:    math.Round(totalRevenue*100) / 100,
		ExpenseSections: expSections,
		TotalExpenses:   math.Round(totalExpenses*100) / 100,
		NetProfit:       math.Round(netProfit*100) / 100,
		IsProfit:        isProfit,
		ProfitMargin:    margin,
	}, nil
}

func buildPnLSections(
	m map[string][]accountingdtos.ReportLineItemDto,
) ([]accountingdtos.PnLSectionDto, float64) {

	// sort group names for deterministic order
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sections []accountingdtos.PnLSectionDto
	var grandTotal float64

	for _, label := range keys {
		lines := m[label]
		sectionTotal := 0.0
		for _, l := range lines {
			sectionTotal += l.Amount
		}
		// sort lines by code within section
		sort.Slice(lines, func(i, j int) bool {
			return lines[i].Code < lines[j].Code
		})
		sections = append(sections, accountingdtos.PnLSectionDto{
			Label: label,
			Lines: lines,
			Total: math.Round(sectionTotal*100) / 100,
		})
		grandTotal += sectionTotal
	}

	return sections, grandTotal
}
