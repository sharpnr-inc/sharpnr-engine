package accounting_service

import (
	"context"
	"errors"
	accountingdtos "sharpnr-engine/modules/accounting/accounting_dtos"
	"sharpnr-engine/modules/accounting/accounting_models"
	"sharpnr-engine/modules/accounting/accounting_repo"
	"time"

	"github.com/google/uuid"
)

func (s *accountingServices) CreateAccount(
	ctx context.Context,
	orgID uuid.UUID,
	req accountingdtos.CreateAccountRequest,
) error {
	return s.accRepo.WithTx(ctx, func(txRepo accounting_repo.AccountingRepositoryInterface) error {

		// 1️⃣ Step 1 --> Get the

		book, err := txRepo.GetAccountingBookByID(ctx, req.AccountingBookID)
		if err != nil {
			return err
		}
		if book.OrgID != orgID {
			return errors.New("accounting does not belong to organization")
		}

		//2️⃣ Step --> Get the Account Group

		group, err := txRepo.GetAccountGroupByID(ctx, req.AccountGroupID)
		if err != nil {
			return err
		}

		if group.AccountingBookID != book.ID {
			return errors.New("account group does not belong to accounting book")
		}

		// 3️⃣ Step --> Create Account

		acc := &accounting_models.Account{
			ID:             uuid.New(),
			Name:           req.Name,
			Code:           req.Code,
			AccountGroupID: req.AccountGroupID,
			AccountGroup:   *group,
			Type:           group.Type,
			OpeningBalance: req.OpeningBalance,
			OpeningIsDebit: req.IsDebit,
			IsDefault:      req.IsDefault,
			Status:         "ACTIVE",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		if !req.IsDebit {
			acc.OpeningBalance = -acc.OpeningBalance
		}

		return txRepo.CreateAccount(ctx, acc)
	})
}
