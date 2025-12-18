package hospitality_acc_setup

import (
	"context"
	accountingconstants "sharpnr-engine/modules/accounting/accounting_constants"
	"sharpnr-engine/modules/accounting/accounting_models"
	"sharpnr-engine/modules/accounting/accounting_repo"
	"time"

	"github.com/google/uuid"
)

type BarAccountingSetupStep struct{}

var barAccountGroups = []struct {
	Name string
	Type accountingconstants.AccountGroupType
}{
	{"Assets", accountingconstants.AccountAsset},
	{"Liabilities", accountingconstants.AccountLiability},
	{"Equity", accountingconstants.AccountEquity},
	{"Income", accountingconstants.AccountIncome},
	{"Expenses", accountingconstants.AccountExpense},
}

func defaultBarAccountName(t accountingconstants.AccountGroupType) string {
	switch t {
	case accountingconstants.AccountAsset:
		return "Cash"
	case accountingconstants.AccountLiability:
		return "Customer Advance"
	case accountingconstants.AccountEquity:
		return "Owner Capital"
	case accountingconstants.AccountIncome:
		return "Bar Revenue"
	case accountingconstants.AccountExpense:
		return "General Expense"
	default:
		return "Unknown"
	}
}

func (s *BarAccountingSetupStep) Name() string {
	return "BarAccountingSetup"
}

func (s *BarAccountingSetupStep) Execute(
	ctx context.Context,
	c AccSetupContext,
) error {
	return c.AccountingRepo.WithTx(ctx, func(txRepo accounting_repo.AccountingRepositoryInterface) error {
		// ------------------------------------------------
		// 1️⃣ Idempotency: Check if BAR AccountingBook exists
		// ------------------------------------------------

		books, err := txRepo.GetAccountingBooksByOrgID(ctx, c.Org.ID)
		if err != nil {
			return err
		}

		for _, b := range books {
			if b.Module != nil && b.SubModule != nil &&
				*b.Module == "HOSPITALITY" &&
				*b.SubModule == "BAR" {
				return nil // already setup
			}
		}

		// ------------------------------------------------
		// 2️⃣ Create AccountingBook (NO FiscalYear here)
		// ------------------------------------------------
		module := "HOSPITALITY"
		subModule := "BAR"

		book := &accounting_models.AccountingBook{
			ID:        uuid.New(),
			OrgID:     c.Org.ID,
			Name:      "Bar Accounting",
			Module:    &module,
			SubModule: &subModule,
			Status:    accountingconstants.Active,
			CreatedAt: time.Now(),
		}

		if err := txRepo.CreateAccountingBook(ctx, book); err != nil {
			return err
		}

		// ------------------------------------------------
		// 3️⃣ Create Account Groups + Default Accounts
		// ------------------------------------------------

		for _, g := range barAccountGroups {
			group := &accounting_models.AccountGroup{
				ID:               uuid.New(),
				Name:             g.Name,
				AccountingBookID: book.ID,
				Type:             g.Type,
				CreatedAt:        time.Now(),
			}
			if err := txRepo.CreateAccountGroup(ctx, group); err != nil {
				return err
			}

			account := &accounting_models.Account{
				ID:             uuid.New(),
				AccountGroupID: group.ID,
				Name:           defaultBarAccountName(g.Type),
				IsDefault:      true,
				Status:         "ACTIVE",
			}

			if err := txRepo.CreateAccount(ctx, account); err != nil {
				return err
			}
		}

		// ------------------------------------------------
		// 4️⃣ Create CURRENT OPEN Fiscal Year (Optional but practical)
		// ------------------------------------------------

		now := time.Now()
		fyStart := time.Date(now.Year(), 4, 1, 0, 0, 0, 0, time.UTC)
		fyEnd := fyStart.AddDate(1, 0, -1)

		fy := &accounting_models.FiscalYear{
			ID:               uuid.New(),
			AccountingBookID: book.ID,
			Name:             "FY " + fyStart.Format("2006") + "-" + fyEnd.Format("2006"),
			StartDate:        fyStart,
			EndDate:          fyEnd,
			Status:           accountingconstants.Open,
		}

		if err := txRepo.CreateFiscalYear(ctx, fy); err != nil {
			return err
		}

		return nil
	})

}
