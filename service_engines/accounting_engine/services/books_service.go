package accounting_service

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	accountingdtos "sharpnr-engine/service_engines/accounting_engine/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *accountingServices) GetAccountingBook(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) (accountingdtos.AccountingBook_Response_dto, error) {
	var response accountingdtos.AccountingBook_Response_dto

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// 1. Fetch the Book
		book, err := s.accRepo.GetAccountingBookByLegalEntityID(txCtx, legalEntityID)
		if err != nil {
			return err
		}

		// 2. Fetch the associated Groups
		accGroups, err := s.accRepo.GetAccountGroupsByBookID(txCtx, book.ID)
		if err != nil {
			return err
		}

		// 3. Map Book model to DTO
		response.ID = book.ID.String()
		response.Name = book.Name
		response.Status = string(book.Status)

		// 4. Map AccountGroup models to DTO slice
		response.Groups = make([]accountingdtos.AccountingGroupResponseDto, len(accGroups))
		for i, group := range accGroups {
			response.Groups[i] = accountingdtos.AccountingGroupResponseDto{
				ID:          group.ID,
				GroupName:   group.Name,
				Type:        string(group.Type),
				LastUpdated: group.UpdatedAt,
			}
		}

		return nil
	})

	if err != nil {
		return accountingdtos.AccountingBook_Response_dto{}, err
	}

	return response, nil
}
