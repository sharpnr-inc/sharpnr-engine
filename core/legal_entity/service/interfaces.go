package legal_entity_service

import (
	"context"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"

	"github.com/google/uuid"
)

type OrgLegalEntityServiceInterface interface {

	// Creates a New Legal Entity
	CreateNewLegalEntity(ctx context.Context, input legal_entity_contracts.NewLegalEntityInput) (legalEntityID uuid.UUID, err error)

	// Returns the Legal Entity By ID
	GetLegalEntityByID(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) (*legal_entity_contracts.LegalEntityResponseDto, error)

	// Returns A Slice of Pointers of Legal Entities
	GetLegalEntityByOrgID(ctx context.Context, orgID uuid.UUID) ([]*legal_entity_contracts.LegalEntityReader, error)

	// This Returns the A Cumelative response of Legal Enitity
	GetLegalEntityDetails(ctx context.Context, orgID, emplId, legalEntityID uuid.UUID) (*legal_entity_contracts.LegalEntityReader, error)

	/* -------------------------------------------------------------------------- */
	/*                          Other Children Interfaces                         */
	/* -------------------------------------------------------------------------- */
	LegalEntityTaxCongfigsServiceInterface
	LegalEntityTaxEngineAdapterInterface
}

/* ---------------------------------- <***> --------------------------------- */

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

	GetTaxCategoryByID(ctx context.Context, taxCategoryID uuid.UUID) (legal_entity_contracts.TaxCategoryReader, error)
}

/* ---------------------------------- <***> --------------------------------- */

type LegalEntityTaxEngineAdapterInterface interface {
	GetTaxCategoryDetails(
		ctx context.Context, taxCategoryID uuid.UUID,
	) (legal_entity_contracts.TaxCategoryReader, error)
}
