package legal_entity_service

import (
	"context"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"

	"github.com/google/uuid"
)

type OrgLegalEntityServiceInterface interface {
	CreateNewLegalEntity(ctx context.Context, input legal_entity_contracts.NewLegalEntityInput) (legalEntityID uuid.UUID, err error)
	GetLegalEntityByOrgID(ctx context.Context, orgID uuid.UUID) ([]*legal_entity_contracts.LegalEntityReader, error)

	LegalEntityTaxCongfigsServiceInterface
	LegalEntityTaxEngineAdapterInterface
}

type LegalEntityTaxCongfigsServiceInterface interface {
	AddLegalEntityTaxRegistration(
		ctx context.Context,
		orgID, emplID, legalEntityID uuid.UUID,
		request legal_entity_contracts.CreateLegalEntityTaxRegistrationRequestDto,
	) error

	GetAllLegelEntityTaxRegistration(
		ctx context.Context,
		orgID, emplID, legalEntityID uuid.UUID,
	) ([]legal_entity_contracts.LegalEntityTaxRegistrationResponseDto, error)

	AddLegalEntityTaxCategory(
		ctx context.Context,
		orgID, emplID, legalEntityID uuid.UUID,
		request legal_entity_contracts.CreateLegalEntityTaxCategory,
	) error

	GetAllLegalEntityTaxCategory(
		ctx context.Context,
		orgID, emplID, legalEntityID uuid.UUID,
	) ([]legal_entity_contracts.LegalEntityTaxCategoryResponseDto, error)
}

type LegalEntityTaxEngineAdapterInterface interface {
	GetTaxCategoryDetails(ctx context.Context, taxCategoryID uuid.UUID, LegalEntityID uuid.UUID) (legal_entity_contracts.LegalEntityTaxCategoryResponseDto, error)
}
