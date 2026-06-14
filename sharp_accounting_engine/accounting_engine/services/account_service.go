package accounting_service

import (
	"context"
	accountingdtos "sharpnr-engine/sharp_accounting_engine/accounting_engine/dto"
	accounting_models "sharpnr-engine/sharp_accounting_engine/accounting_engine/models"
	"time"

	"github.com/google/uuid"
)

func (s *accountingServices) CreateAccount(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
	accountRequest accountingdtos.CreateAccountRequest,
) (accountID uuid.UUID, err error) {

	// if err := accountRequest.ValidateRequest(); err != nil {
	// 	return uuid.Nil, err
	// }

	account := &accounting_models.Account{
		ID:             uuid.New(),
		AccountGroupID: accountRequest.AccountGroupID,
		Name:           accountRequest.Name,
		Code:           accountRequest.Code,
		OpeningBalance: accountRequest.OpeningBalance,
		OpeningIsDebit: accountRequest.IsDebit,
		Status:         "ACTIVE",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := s.accRepo.CreateAccount(ctx, account); err != nil {
		return uuid.Nil, err
	}

	return account.ID, nil
}

/* ------ <-- This method is used to fetch the leger of the account --> ----- */

func (s *accountingServices) GetAccountLeger(ctx context.Context, orgID, emplID, legalEntityID, accountID uuid.UUID) (accountingdtos.AccountLedgerDto, error) {

	from := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Now()

	account, err := s.accRepo.GetAccountByID(ctx, accountID)
	if err != nil {
		return accountingdtos.AccountLedgerDto{}, err
	}

	openingBalance := account.OpeningBalance
	if account.OpeningIsDebit {
		openingBalance = -openingBalance
	}

	lines, _ := s.accRepo.GetLedger(accountID, from, to)
	for i := range lines {
		lines[i].RunningBalance += openingBalance
	}

	// 5. closing balance = last line's running balance
	closing := openingBalance
	if len(lines) > 0 {
		closing = lines[len(lines)-1].RunningBalance
	}

	return accountingdtos.AccountLedgerDto{
		AccountID:      account.ID,
		AccountName:    account.Name,
		AccountCode:    account.Code,
		OpeningBalance: account.OpeningBalance,
		OpeningIsDebit: account.OpeningIsDebit,
		Lines:          lines,
		ClosingBalance: closing,
	}, nil

}
