package legal_entity_service

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	legal_entity_constants "sharpnr-engine/core/legal_entity/constants"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *orgLegalEntity) AddNewAddressInLegalEntity(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
	req legal_entity_contracts.CreateLegalEntityAddressRequest,
) (addressID uuid.UUID, err error) {

	err = s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)
		addressID = uuid.New()

		if err := s.legalEntityRepo.CreateAddressOfLegalEntity(txCtx, &legal_entity_models.LegalEntityAddress{
			ID:            addressID,
			LegalEntityID: legalEntityID,
			Type:          legal_entity_constants.LegalEntityAddressType(req.AddressType.Label()),
			AddressLine1:  req.AddressLine1,
			AddressLine2:  req.AddressLine2,
			Landmark:      req.Landmark,
			City:          req.City,
			State:         req.State,
			GSTStateCode:  req.GSTStateCode,
			Pincode:       req.Pincode,
			Country:       req.Country,
			IsPrimary:     false,
			PropertyID:    &uuid.Nil,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	return addressID, nil
}
