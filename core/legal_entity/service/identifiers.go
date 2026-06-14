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

func (s *orgLegalEntity) AddNewIdentifiersInLegalEntity(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
	req *legal_entity_contracts.AddIdentifiersInLegalEntityRequestDto,
) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// Map structural state into model instance
		identifierModel := legal_entity_models.LegalEntityIdentifier{
			ID:            uuid.New(),
			LegalEntityID: legalEntityID,
			Type:          req.Type,
			Value:         req.Value,
			ExpiresAt:     nil,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		// Ensure we capture, evaluate, and return database error signatures
		if err := s.legalEntityRepo.AddLegalEntityIdentifier(txCtx, identifierModel); err != nil {
			return err
		}

		return nil
	})
}
