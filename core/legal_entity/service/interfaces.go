package legal_entity_service

import (
	"context"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"

	"github.com/google/uuid"
)

type OrgLegalEntityServiceInterface interface {
	CreateNewLegalEntity(ctx context.Context, orgID, emplID uuid.UUID, input legal_entity_contracts.NewLegalEntityRequestDto) (legalEntityID uuid.UUID, err error)
	GetLegalEntityByID(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) (*legal_entity_contracts.LegalEntityResponseDto, error)
	GetLegalEntityByOrgID(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID) ([]legal_entity_contracts.LegalEntityResponseDto, error)
	GetLegalEntityDetails(ctx context.Context, orgID, emplId, legalEntityID uuid.UUID) (legal_entity_contracts.LegalEntityDetails, error)

	AddNewAddressInLegalEntity(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID, req legal_entity_contracts.CreateLegalEntityAddressRequest) (addressID uuid.UUID, err error)
	// GetLegalEntityAddresses(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) ([]string, error)

	LegalEntityTaxCongfigsServiceInterface
	LegalEntityTaxEngineAdapterInterface
	LegalEntityTaxCategoryService

	// Adds New Contact in Legal Entity
	AddNewContactInLegalEntity(
		ctx context.Context,
		orgID, emplID, legalEntityID uuid.UUID,
		req *legal_entity_contracts.CreateLegalEntityContactsRequestDto,
	) error

	AddNewBankAccounInLegalEntity(
		ctx context.Context,
		orgID, emplID, legalEntityID uuid.UUID,
		req *legal_entity_contracts.AddBankAccountInLegalEntityRequestDto,
	) error

	AddNewIdentifiersInLegalEntity(
		ctx context.Context,
		orgID, emplID, legalEntityID uuid.UUID,
		req *legal_entity_contracts.AddIdentifiersInLegalEntityRequestDto,
	) error
}

type LegalEntityTaxCategoryService interface {
	GetLegalEntityTaxCategory(ctx context.Context, orgID, emplID, LegalEntityID uuid.UUID) ([]legal_entity_contracts.LegalEntityTaxCategoryResponseDto, error)
}

/* ---------------------------------- <***> --------------------------------- */

type LegalEntityTaxCongfigsServiceInterface interface {
	GetTaxCategoryByID(ctx context.Context, taxCategoryID uuid.UUID) (legal_entity_contracts.TaxCategoryReader, error)
}

/* ---------------------------------- <***> --------------------------------- */

type LegalEntityTaxEngineAdapterInterface interface {
	GetTaxCategoryDetails(ctx context.Context, taxCategoryID uuid.UUID) (legal_entity_contracts.TaxCategoryReader, error)
}
