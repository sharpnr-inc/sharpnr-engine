package accounting_service

import (
	"context"
	"fmt"
	dbtx "sharpnr-engine/core/dbTx"
	accountingconstants "sharpnr-engine/sharp_services/accounting_engine/constants"
	accounting_engine_contracts "sharpnr-engine/sharp_services/accounting_engine/contracts"
	accounting_engine_helpers "sharpnr-engine/sharp_services/accounting_engine/helpers"
	accounting_models "sharpnr-engine/sharp_services/accounting_engine/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/******************************************
 * These are methods for Accounting Setup *
 ******************************************/

func (r *accountingServices) SetupAccountingForLegalEntity() {
	panic("unimplemented")
}

//----------------------------------------------------------------//

// this methods creates the Drafts Accounting Setup
// the Drafts Accounting setup is the 5 Accounting group
// within 5 accounting group we create the 1 Default Account
// with Current date to 31 March 2027
func (r *accountingServices) CreateDraftAccountingSetup(ctx context.Context, input *accounting_engine_contracts.DraftAccountingSetupInput) error {
	// 1. txManager now automatically decides: JOIN existing tx or START new tx
	return r.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		// 2. We MUST inject 'tx' back into the context.
		// Even if 'tx' was already there, this ensures the child calls (repos)
		// definitely have access to the current transaction instance.
		txCtx := dbtx.InjectTx(ctx, tx)
		now := time.Now()

		// Step One: Create The Accounting book
		accBookInput := &accounting_models.AccountingBook{
			ID:                        uuid.New(),
			OrganizationID:            input.OrgID,
			OrganizationLegalEntityID: input.LegalEntityID,
			Name:                      accbookname(input.Name_of_legalEntity),
			Status:                    accountingconstants.Active,
			CreatedAt:                 now,
			UpdatedAt:                 now,
		}

		// Pass txCtx so repo can ExtractTx
		if err := r.accRepo.CreateAccountingBook(txCtx, accBookInput); err != nil {
			return fmt.Errorf("failed to create accounting book: %w", err)
		}

		// Step Two: Create Fiscal year
		fy := accounting_engine_helpers.GenerateIndianFiscalYear(accBookInput.ID)
		if err := r.accRepo.CreateFiscalYear(txCtx, fy); err != nil {
			return fmt.Errorf("failed to create fiscal year: %w", err)
		}

		// Step Three: Create Groups and Default Accounts
		for _, ag := range accountingconstants.AllAccountGroups {
			accgrp := &accounting_models.AccountGroup{
				ID:               uuid.New(),
				Name:             fmt.Sprintf("Default %s Group", ag),
				AccountingBookID: accBookInput.ID,
				Type:             ag,
				CreatedAt:        now,
				UpdatedAt:        now,
			}

			if err := r.accRepo.CreateAccountGroup(txCtx, accgrp); err != nil {
				return fmt.Errorf("failed to create account group %s: %w", ag, err)
			}

			// Determine normal balance nature
			isDebit := (ag == accountingconstants.AccountAsset || ag == accountingconstants.AccountExpense)

			defaultAcc := &accounting_models.Account{
				ID:             uuid.New(),
				AccountGroupID: accgrp.ID,
				Name:           fmt.Sprintf("General %s Ledger", ag),
				Code:           fmt.Sprintf("%s-001", ag),
				OpeningBalance: 0,
				OpeningIsDebit: isDebit,
				Status:         string(accountingconstants.Active),
				CreatedAt:      now,
				UpdatedAt:      now,
			}

			if err := r.accRepo.CreateAccount(txCtx, defaultAcc); err != nil {
				return fmt.Errorf("failed to create default account for %s: %w", ag, err)
			}
		}

		return nil
	})
}

func accbookname(name string) string {
	return name + " Accounting Book"
}
