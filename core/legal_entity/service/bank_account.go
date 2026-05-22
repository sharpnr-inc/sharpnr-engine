package legal_entity_service

import (
	"context"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"
	dbtx "sharpnr-engine/pkg/dbTx"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *orgLegalEntity) AddNewBankAccounInLegalEntity(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
	req *legal_entity_contracts.AddBankAccountInLegalEntityRequestDto,
) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		_, err := s.legalEntityRepo.AddBankAccountInLegalEntity(txCtx, legal_entity_models.LegalEntityBankAccount{
			ID:                uuid.New(),
			LegalEntityID:     legalEntityID,
			AccountHolderName: req.AccountHolderName,
			AccountNumber:     req.AccountNumber,
			IFSC:              req.IFSC,
			BankName:          req.BankName,
			BranchName:        req.BranchName,
			MICRCode:          req.MICRCode,
			AccountType:       req.AccountType,
			IsPrimary:         req.IsPrimary,
			IsVerified:        false,
			IsActive:          true,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		})
		if err != nil {
			return err
		}

		return nil

	})
}
