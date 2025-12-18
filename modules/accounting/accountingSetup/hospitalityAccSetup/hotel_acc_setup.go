package hospitality_acc_setup

import (
	"context"
	accountingconstants "sharpnr-engine/modules/accounting/accounting_constants"
	"sharpnr-engine/modules/accounting/accounting_models"
	"sharpnr-engine/modules/accounting/accounting_repo"
	"time"

	"github.com/google/uuid"
)

type HotelAccountingSetupStep struct{}

var hotelAccountGroups = []struct {
	Name string
	Type accountingconstants.AccountGroupType
}{
	{"Assets", accountingconstants.AccountAsset},
	{"Liabilities", accountingconstants.AccountLiability},
	{"Equity", accountingconstants.AccountEquity},
	{"Income", accountingconstants.AccountIncome},
	{"Expenses", accountingconstants.AccountExpense},
}

func defaultHotelAccountName(t accountingconstants.AccountGroupType) string {
	switch t {
	case accountingconstants.AccountAsset:
		return "Cash"
	case accountingconstants.AccountLiability:
		return "Customer Advance"
	case accountingconstants.AccountEquity:
		return "Owner Capital"
	case accountingconstants.AccountIncome:
		return "Hotel Revenue"
	case accountingconstants.AccountExpense:
		return "General Expense"
	default:
		return "Unknown"
	}
}

func (s *HotelAccountingSetupStep) Name() string {
	return "BarAccountingSetup"
}

func (s *HotelAccountingSetupStep) Execute(
	ctx context.Context,
	c AccSetupContext,
) error {
	return c.AccountingRepo.WithTx(ctx, func(txRepo accounting_repo.AccountingRepositoryInterface) error {

		// ------------------------------------------------
		// 1️⃣ Idempotency: Check if HOTEL book already exists
		// ------------------------------------------------
		books, err := txRepo.GetAccountingBooksByOrgID(ctx, c.Org.ID)
		if err != nil {
			return err
		}

		for _, b := range books {
			if b.Module != nil && b.SubModule != nil &&
				*b.Module == "HOSPITALITY" &&
				*b.SubModule == "HOTEL" {
				return nil // already set up
			}
		}

		// ------------------------------------------------
		// 2️⃣ Create AccountingBook (NO FiscalYear here)
		// ------------------------------------------------
		module := "HOSPITALITY"
		subModule := "HOTEL"

		book := &accounting_models.AccountingBook{
			ID:        uuid.New(),
			OrgID:     c.Org.ID,
			Name:      "Hotel Accounting",
			Module:    &module,
			SubModule: &subModule,
			Status:    accountingconstants.Active,
		}

		if err := txRepo.CreateAccountingBook(ctx, book); err != nil {
			return err
		}

		// ------------------------------------------------
		// 3️⃣ Create Account Groups + Default Accounts
		// ------------------------------------------------
		for _, g := range hotelAccountGroups {

			group := &accounting_models.AccountGroup{
				ID:               uuid.New(),
				AccountingBookID: book.ID,
				Name:             g.Name,
				Type:             g.Type,
			}

			if err := txRepo.CreateAccountGroup(ctx, group); err != nil {
				return err
			}

			account := &accounting_models.Account{
				ID:             uuid.New(),
				AccountGroupID: group.ID,
				Name:           defaultHotelAccountName(g.Type),
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
