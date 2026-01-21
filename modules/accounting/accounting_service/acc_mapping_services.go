package accounting_service

import (
	"context"
	accountingdtos "sharpnr-engine/modules/accounting/accounting_dtos"
	accounting_models "sharpnr-engine/modules/accounting/models"
	"time"

	"github.com/google/uuid"
)

func (s *accountingServices) GetAccountMapping(
	ctx context.Context,
	orgID uuid.UUID,
) ([]accountingdtos.AccountMappingDTO, error) {

	mappings, err := s.accRepo.GetAccountMappingsByOrg(ctx, orgID)
	if err != nil {
		return nil, err
	}

	resp := make([]accountingdtos.AccountMappingDTO, 0, len(mappings))

	for _, m := range mappings {
		resp = append(resp, accountingdtos.AccountMappingDTO{
			ID:          m.ID,
			Module:      m.Module,
			SourceType:  m.SourceType,
			SourceCode:  m.SourceCode,
			AccountID:   m.AccountID,
			AccountName: m.Account.Name,
			AccountCode: m.Account.Code,
		})
	}

	return resp, nil
}

func (s *accountingServices) CreateAccountMapping(
	ctx context.Context,
	orgId uuid.UUID,
	req accountingdtos.CreateAccountMappingRequest_dto,
) error {

	account, err := s.accRepo.GetAccountByID(ctx, req.AccountID)
	if err != nil {
		return err
	}

	mapping := &accounting_models.AccountMapping{
		ID:         uuid.New(),
		OrgID:      orgId,
		Module:     "HOSPITALITY",
		SubModule:  "HOTEL",
		SourceType: req.SourceType,
		SourceCode: req.SourceCode,
		AccountID:  account.ID,
		Account:    *account,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := s.accRepo.CreateAccountMapping(ctx, mapping); err != nil {
		return err
	}

	return nil
}
