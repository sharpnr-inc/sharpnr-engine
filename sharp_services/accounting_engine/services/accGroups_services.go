package accounting_service

import (
	"context"
	accountingdtos "sharpnr-engine/sharp_services/accounting_engine/dto"

	"github.com/google/uuid"
)

func (s *accountingServices) GetAccountingGroupOfBook(
	ctx context.Context,
	orgID, bookId uuid.UUID,
) (*accountingdtos.GetAccountingGroupOfBook_Response_dto, error) {

	// Fetch account groups for the book
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

		// 🔹 Fetch accounts for this group
		accounts, err := s.accRepo.GetAccountsByGroupID(ctx, grp.ID)
		if err != nil {
			return nil, err
		}

		// 🔹 Map accounts to DTO
		accountDTOs := make(
			[]accountingdtos.AccountsResponse_dto,
			0,
			len(accounts),
		)

		for _, acc := range accounts {
			accountDTOs = append(accountDTOs, accountingdtos.AccountsResponse_dto{
				ID:   acc.ID.String(),
				Name: acc.Name,
			})
		}

		// 🔹 Append group with accounts
		resp.AccountGroups = append(
			resp.AccountGroups,
			accountingdtos.AccountingGroup_Response_dto{
				ID:               grp.ID.String(),
				Name:             grp.Name,
				AccountingBookID: grp.AccountingBookID.String(),
				Type:             string(grp.Type),
				CreatedAt:        grp.CreatedAt,
				UpdatedAt:        grp.UpdatedAt,
				Accounts:         accountDTOs,
			},
		)
	}

	return resp, nil
}
