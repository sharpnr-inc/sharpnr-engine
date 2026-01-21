package accounting_service

import (
	"context"
	accountingdtos "sharpnr-engine/sharp_services/accounting_engine/dto"

	"github.com/google/uuid"
)

func (s *accountingServices) GetAccountingBooksOfOrg(
	ctx context.Context,
	orgID uuid.UUID,
) (*accountingdtos.GetAccountBooksOfOrg_Response_dto, error) {

	accBooks, err := s.accRepo.GetAccountingBooksByOrgID(ctx, orgID)
	if err != nil {
		return nil, err
	}

	resp := &accountingdtos.GetAccountBooksOfOrg_Response_dto{
		AccountBooks: make([]accountingdtos.AccountingBook_Response_dto, 0, len(accBooks)),
	}

	for _, ac := range accBooks {
		resp.AccountBooks = append(resp.AccountBooks, accountingdtos.AccountingBook_Response_dto{
			ID:        ac.ID.String(),
			Name:      ac.Name,
			Module:    ac.Module,
			SubModule: ac.SubModule,
			Status:    string(ac.Status),
		})
	}

	return resp, nil
}
