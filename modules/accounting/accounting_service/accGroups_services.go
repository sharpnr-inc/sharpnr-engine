package accounting_service

import (
	"context"
	accountingdtos "sharpnr-engine/modules/accounting/accounting_dtos"

	"github.com/google/uuid"
)

func (s *accountingServices) GetAccountingGroupOfBook(
	ctx context.Context,
	orgID, bookId uuid.UUID,
) (*accountingdtos.GetAccountingGroupOfBook_Response_dto, error) {

	// Fetch groups (repo should already enforce book → org relationship ideally)
	acGrps, err := s.accRepo.GetAccountGroupsByBookID(ctx, bookId)
	if err != nil {
		return nil, err
	}

	resp := &accountingdtos.GetAccountingGroupOfBook_Response_dto{
		AccountGroups: make(
			[]accountingdtos.AccountingGroup_Response_dto,
			0,
			len(acGrps),
		),
	}

	for _, grp := range acGrps {
		resp.AccountGroups = append(resp.AccountGroups, accountingdtos.AccountingGroup_Response_dto{
			ID:               grp.ID.String(),
			Name:             grp.Name,
			AccountingBookID: grp.AccountingBookID.String(),
			Type:             string(grp.Type),
			CreatedAt:        grp.CreatedAt,
			UpdatedAt:        grp.UpdatedAt,
		})
	}

	return resp, nil
}
