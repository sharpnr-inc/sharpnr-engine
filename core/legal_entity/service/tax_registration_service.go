package legal_entity_service

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *orgLegalEntity) AddLegalEntityTaxRegistration(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
	request legal_entity_contracts.CreateLegalEntityTaxRegistrationRequestDto,
) error {
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		taxRegistration := &legal_entity_models.LegalEntityTaxRegistration{
			ID:                 uuid.New(),
			LegalEntityID:      request.LegalEntityID,
			TaxType:            request.TaxType,
			RegistrationNumber: request.RegistrationNumber,
			State:              request.State,
			Country:            request.Country,
			EffectiveFrom:      request.EffectiveFrom,
			EffectiveTo:        request.EffectiveTo,
			IsActive:           true,
			CreatedAt:          time.Now(),
		}

		if err := s.legalEntityRepo.CreateTaxRegistration(txCtx, taxRegistration); err != nil {
			return err
		}

		return nil
	})
}

func (s *orgLegalEntity) GetAllLegelEntityTaxRegistration(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
) ([]legal_entity_contracts.LegalEntityTaxRegistrationResponseDto, error) {

	var response []legal_entity_contracts.LegalEntityTaxRegistrationResponseDto

	taxRegistrations, err := s.legalEntityRepo.GetTaxRegistrations(ctx, legalEntityID)
	if err != nil {
		return response, err
	}

	for _, tr := range taxRegistrations {

		dto := legal_entity_contracts.LegalEntityTaxRegistrationResponseDto{
			ID:                 tr.ID,
			TaxType:            tr.TaxType,
			RegistrationNumber: tr.RegistrationNumber,
			Country:            tr.Country,
			State:              tr.State,
			EffectiveFrom:      tr.EffectiveFrom,
			EffectiveTo:        tr.EffectiveTo,
			IsActive:           tr.IsActive,
			LastUpdatedAt:      tr.CreatedAt,
		}

		response = append(response, dto)
	}

	return response, nil
}
