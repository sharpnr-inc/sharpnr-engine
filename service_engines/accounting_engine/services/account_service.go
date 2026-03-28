package accounting_service

import (
	"context"
	accountingdtos "sharpnr-engine/service_engines/accounting_engine/dto"
	accounting_models "sharpnr-engine/service_engines/accounting_engine/models"
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
