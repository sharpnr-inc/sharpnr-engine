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

func (s *orgLegalEntity) AddNewContactInLegalEntity(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
	req *legal_entity_contracts.CreateLegalEntityContactsRequestDto,
) error {
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		if err := s.legalEntityRepo.AddLegalEntityContacts(txCtx, legal_entity_models.LegalEntityContact{
			ID:            uuid.New(),
			LegalEntityID: legalEntityID,
			Type:          req.Type,
			FirstName:     req.FirstName,
			LastName:      req.LastName,
			Designation:   req.Designation,
			Email:         req.Email,
			Mobile:        req.Mobile,
			DIN:           req.DIN,
			IsPrimary:     req.IsPrimary,
			IsActive:      true,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}); err != nil {
			return err
		}

		return nil
	})

}
