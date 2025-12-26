package accounting_service

import (
	"context"
	accountingdtos "sharpnr-engine/modules/accounting/accounting_dtos"

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
