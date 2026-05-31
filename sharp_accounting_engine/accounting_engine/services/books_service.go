package accounting_service

import (
	"context"
	dbtx "sharpnr-engine/pkg/dbTx"
	accountingdtos "sharpnr-engine/sharp_accounting_engine/accounting_engine/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *accountingServices) GetAccountingBook(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
) (accountingdtos.AccountingBook_Response_dto, error) {

	var response accountingdtos.AccountingBook_Response_dto

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		// 1️⃣ Fetch the Book
		book, err := s.accRepo.GetAccountingBookByLegalEntityID(txCtx, legalEntityID)
		if err != nil {
			return err
		}

		// 2️⃣ Fetch Groups
		accGroups, err := s.accRepo.GetAccountGroupsByBookID(txCtx, book.ID)
		if err != nil {
			return err
		}

		// 3️⃣ Map Book → DTO
		response.ID = book.ID.String()
		response.Name = book.Name
		response.Status = string(book.Status)

		response.Groups = make([]accountingdtos.AccountingGroupResponseDto, len(accGroups))

		// 4️⃣ Map Groups
		for i, group := range accGroups {

			groupDTO := accountingdtos.AccountingGroupResponseDto{
				ID:          group.ID,
				GroupName:   group.Name,
				Type:        string(group.Type),
				LastUpdated: group.UpdatedAt,
			}

			// 5️⃣ Fetch Accounts inside this Group
			accounts, err := s.accRepo.GetAccountsByGroupID(txCtx, group.ID)
			if err != nil {
				return err
			}

			groupDTO.Accounts = make([]accountingdtos.AccountsResponseDto, len(accounts))

			for j, acc := range accounts {

				groupDTO.Accounts[j] = accountingdtos.AccountsResponseDto{
					ID:          acc.ID,
					AccountName: acc.Name,
					Code:        acc.Code,
					Status:      acc.Status,
					LastUpdated: acc.UpdatedAt,
				}
			}

			response.Groups[i] = groupDTO
		}

		return nil
	})

	if err != nil {
		return accountingdtos.AccountingBook_Response_dto{}, err
	}

	return response, nil
}
