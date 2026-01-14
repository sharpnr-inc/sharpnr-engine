package acc_sere_wf_adapters

import (
	accountingconstants "sharpnr-engine/modules/accounting/accounting_constants"
	"sharpnr-engine/modules/accounting/accounting_models"
	"sharpnr-engine/modules/accounting/accounting_repo"
	"sharpnr-engine/sere"
	"time"

	"github.com/google/uuid"
)

// So this is bad architecture we will not create accounting during the application

var accountGroups = []struct {
	Name string
	Type accountingconstants.AccountGroupType
}{
	{"Assets", accountingconstants.AccountAsset},
	{"Liabilities", accountingconstants.AccountLiability},
	{"Equity", accountingconstants.AccountEquity},
	{"Income", accountingconstants.AccountIncome},
	{"Expenses", accountingconstants.AccountExpense},
}

func defaultAccountName(
	module string,
	t accountingconstants.AccountGroupType,
) string {

	if t != accountingconstants.AccountIncome {
		switch t {
		case accountingconstants.AccountAsset:
			return "Cash"
		case accountingconstants.AccountLiability:
			return "Customer Advance"
		case accountingconstants.AccountEquity:
			return "Owner Capital"
		case accountingconstants.AccountExpense:
			return "General Expense"
		default:
			return "Unknown"
		}
	}

	// Income account differs by module
	switch module {
	case "BAR":
		return "Bar Revenue"
	case "HOTEL":
		return "Hotel Revenue"
	case "RESTAURANT":
		return "Restaurant Revenue"
	default:
		return "Revenue"
	}
}

func barAccountingDefaultSetup(
	accountingRepo accounting_repo.AccountingRepositoryInterface,
) sere.WorkflowHandler {

	return func(wf *sere.Workflow) error {

		// 0️⃣ Check if BAR selected
		modules := wf.Data["sector_modules"].([]string)
		hasBar := false
		for _, m := range modules {
			if m == "BAR" {
				hasBar = true
				break
			}
		}
		if !hasBar {
			return nil
		}

		orgID := wf.Data["org_id"].(uuid.UUID)

		// 1️⃣ Idempotency
		books, err := accountingRepo.GetAccountingBooksByOrgID(wf.Ctx, orgID)
		if err != nil {
			return err
		}
		for _, b := range books {
			if b.Module != nil && b.SubModule != nil &&
				*b.Module == "HOSPITALITY" && *b.SubModule == "BAR" {
				return nil
			}
		}

		// 2️⃣ Create AccountingBook
		module := "HOSPITALITY"
		subModule := "BAR"

		book := &accounting_models.AccountingBook{
			ID:        uuid.New(),
			OrgID:     orgID,
			Name:      "Bar Accounting",
			Module:    &module,
			SubModule: &subModule,
			Status:    accountingconstants.Active,
			CreatedAt: time.Now(),
		}

		if err := accountingRepo.CreateAccountingBook(wf.Ctx, book); err != nil {
			return err
		}

		// 3️⃣ Groups + Accounts
		for _, g := range accountGroups {

			group := &accounting_models.AccountGroup{
				ID:               uuid.New(),
				Name:             g.Name,
				AccountingBookID: book.ID,
				Type:             g.Type,
				CreatedAt:        time.Now(),
			}
			if err := accountingRepo.CreateAccountGroup(wf.Ctx, group); err != nil {
				return err
			}

			account := &accounting_models.Account{
				ID:             uuid.New(),
				AccountGroupID: group.ID,
				Name:           defaultAccountName("BAR", g.Type),
				IsDefault:      true,
				Status:         "ACTIVE",
				Code:           "DEFAULT-01",
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			}
			if err := accountingRepo.CreateAccount(wf.Ctx, account); err != nil {
				return err
			}
		}

		// 4️⃣ Fiscal Year
		now := time.Now()
		fyStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
		fyEnd := fyStart.AddDate(1, 0, -1)

		fy := &accounting_models.FiscalYear{
			ID:               uuid.New(),
			OrgID:            orgID,
			AccountingBookID: book.ID,
			Name:             "FY " + fyStart.Format("2006") + "-" + fyEnd.Format("2006"),
			StartDate:        fyStart,
			EndDate:          fyEnd,
			Status:           accountingconstants.Open,
		}

		return accountingRepo.CreateFiscalYear(wf.Ctx, fy)
	}
}

func hotelAccountingDefaultSetup(
	accountingRepo accounting_repo.AccountingRepositoryInterface,
) sere.WorkflowHandler {

	return func(wf *sere.Workflow) error {

		modules := wf.Data["sector_modules"].([]string)
		hasHotel := false
		for _, m := range modules {
			if m == "HOTEL" {
				hasHotel = true
				break
			}
		}
		if !hasHotel {
			return nil
		}

		orgID := wf.Data["org_id"].(uuid.UUID)

		books, err := accountingRepo.GetAccountingBooksByOrgID(wf.Ctx, orgID)
		if err != nil {
			return err
		}
		for _, b := range books {
			if b.Module != nil && b.SubModule != nil &&
				*b.Module == "HOSPITALITY" && *b.SubModule == "HOTEL" {
				return nil
			}
		}

		module := "HOSPITALITY"
		subModule := "HOTEL"

		book := &accounting_models.AccountingBook{
			ID:        uuid.New(),
			OrgID:     orgID,
			Name:      "Hotel Accounting",
			Module:    &module,
			SubModule: &subModule,
			Status:    accountingconstants.Active,
			CreatedAt: time.Now(),
		}

		if err := accountingRepo.CreateAccountingBook(wf.Ctx, book); err != nil {
			return err
		}

		for _, g := range accountGroups {

			group := &accounting_models.AccountGroup{
				ID:               uuid.New(),
				Name:             g.Name,
				AccountingBookID: book.ID,
				Type:             g.Type,
				CreatedAt:        time.Now(),
			}
			if err := accountingRepo.CreateAccountGroup(wf.Ctx, group); err != nil {
				return err
			}

			account := &accounting_models.Account{
				ID:             uuid.New(),
				AccountGroupID: group.ID,
				Name:           defaultAccountName("HOTEL", g.Type),
				IsDefault:      true,
				Status:         "ACTIVE",
				Code:           "DEFAULT-01",
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			}
			if err := accountingRepo.CreateAccount(wf.Ctx, account); err != nil {
				return err
			}
		}

		now := time.Now()
		fyStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
		fyEnd := fyStart.AddDate(1, 0, -1)

		fy := &accounting_models.FiscalYear{
			ID:               uuid.New(),
			OrgID:            orgID,
			AccountingBookID: book.ID,
			Name:             "FY " + fyStart.Format("2006") + "-" + fyEnd.Format("2006"),
			StartDate:        fyStart,
			EndDate:          fyEnd,
			Status:           accountingconstants.Open,
		}

		return accountingRepo.CreateFiscalYear(wf.Ctx, fy)
	}
}

func restaurantAccountingDefaultSetup(
	accountingRepo accounting_repo.AccountingRepositoryInterface,
) sere.WorkflowHandler {

	return func(wf *sere.Workflow) error {

		modules := wf.Data["sector_modules"].([]string)
		hasRestaurant := false
		for _, m := range modules {
			if m == "RESTAURANT" {
				hasRestaurant = true
				break
			}
		}
		if !hasRestaurant {
			return nil
		}

		orgID := wf.Data["org_id"].(uuid.UUID)

		books, err := accountingRepo.GetAccountingBooksByOrgID(wf.Ctx, orgID)
		if err != nil {
			return err
		}
		for _, b := range books {
			if b.Module != nil && b.SubModule != nil &&
				*b.Module == "HOSPITALITY" && *b.SubModule == "RESTAURANT" {
				return nil
			}
		}

		module := "HOSPITALITY"
		subModule := "RESTAURANT"

		book := &accounting_models.AccountingBook{
			ID:        uuid.New(),
			OrgID:     orgID,
			Name:      "Restaurant Accounting",
			Module:    &module,
			SubModule: &subModule,
			Status:    accountingconstants.Active,
			CreatedAt: time.Now(),
		}

		if err := accountingRepo.CreateAccountingBook(wf.Ctx, book); err != nil {
			return err
		}

		for _, g := range accountGroups {

			group := &accounting_models.AccountGroup{
				ID:               uuid.New(),
				Name:             g.Name,
				AccountingBookID: book.ID,
				Type:             g.Type,
				CreatedAt:        time.Now(),
			}
			if err := accountingRepo.CreateAccountGroup(wf.Ctx, group); err != nil {
				return err
			}

			account := &accounting_models.Account{
				ID:             uuid.New(),
				AccountGroupID: group.ID,
				Name:           defaultAccountName("RESTAURANT", g.Type),
				IsDefault:      true,
				Status:         "ACTIVE",
				Code:           "DEFAULT-01",
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			}
			if err := accountingRepo.CreateAccount(wf.Ctx, account); err != nil {
				return err
			}
		}

		now := time.Now()
		fyStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
		fyEnd := fyStart.AddDate(1, 0, -1)

		fy := &accounting_models.FiscalYear{
			ID:               uuid.New(),
			OrgID:            orgID,
			AccountingBookID: book.ID,
			Name:             "FY " + fyStart.Format("2006") + "-" + fyEnd.Format("2006"),
			StartDate:        fyStart,
			EndDate:          fyEnd,
			Status:           accountingconstants.Open,
		}

		return accountingRepo.CreateFiscalYear(wf.Ctx, fy)
	}
}
