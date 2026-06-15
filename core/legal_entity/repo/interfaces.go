package legal_entity_repo

import (
	"context"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"

	"github.com/google/uuid"
)

type OrgLegalEntityRepoInterface interface {

	// Creates a New Legal Entity
	CreateLegalEntity(ctx context.Context, legalEntity *legal_entity_models.LegalEntity) error
	GetLegalEntityByID(ctx context.Context, legalEntityID uuid.UUID) (*legal_entity_models.LegalEntity, error)
	// Updates Legal Entity on Allowed
	// the models it updates is legal_entity_models.LegalEntity
	UpdateLegalEntity(ctx context.Context, legalEntityID uuid.UUID, req legal_entity_contracts.UpdateLegalEntityNamesRequestDto) error

	// List all the Legal Entity of the Organization,
	// Returns the Slices if No record found it returns the gorm Err Record Nt found
	GetLegalEntitiesByOrgID(ctx context.Context, orgID uuid.UUID) ([]legal_entity_models.LegalEntity, error)
	// Switch on/Off for the Legal Entity
	UpdateLegalEntityStatus(ctx context.Context, legalEntityID uuid.UUID, isActive bool) error

	LegalEntityBankAccountRepoInterface
	LegalEntityComplianceRepoInterface
	LegalEntityTaxCategoryRepoInterface
	LegalEntityAddressRepoInterface
	LegalEntityIdentifiersRepoInterface
	LegalEntityContactsRepoInterface
	LegalEntityTaxRegistrationRepoInterface
}

/* ---------------------------------- <***> --------------------------------- */

type LegalEntityBankAccountRepoInterface interface {
	AddBankAccountInLegalEntity(ctx context.Context, request legal_entity_models.LegalEntityBankAccount) (bankAccountID uuid.UUID, err error)
	UpdateBankAccountDetails(ctx context.Context, bankAccountID uuid.UUID, updateReq legal_entity_contracts.UpdateLegalEntityBankAccountDetailsRequest) error
	FindBankAccountByID(ctx context.Context, entryID uuid.UUID) (*legal_entity_models.LegalEntityBankAccount, error)
	FindBankAccountsByLegalEntityID(ctx context.Context, legalEntityID uuid.UUID) ([]legal_entity_models.LegalEntityBankAccount, error)
}

/* ---------------------------------- <***> --------------------------------- */

type LegalEntityComplianceRepoInterface interface {
	CreateLegalEntityCompliance(ctx context.Context, compliance *legal_entity_models.LegalEntityCompliance) error
	FindAllLegalCompliancesByLegalEntityID(ctx context.Context, legalEntityID uuid.UUID) ([]*legal_entity_models.LegalEntityCompliance, error)
	UpdateLegalEntityCompliance(ctx context.Context, complianceID uuid.UUID) error
}

/* ---------------------------------- <***> --------------------------------- */

type LegalEntityAddressRepoInterface interface {
	CreateAddressOfLegalEntity(ctx context.Context, address *legal_entity_models.LegalEntityAddress) error
	GetAddressesOfLegalEntity(ctx context.Context, legalEntityID uuid.UUID) ([]legal_entity_models.LegalEntityAddress, error)
	UpdateLegalEntityAddress(ctx context.Context, legalEntityID, addressID uuid.UUID, req legal_entity_contracts.UpdateLegalEntityAddressRequestDto) error
}

/* ---------------------------------- <***> --------------------------------- */

type LegalEntityIdentifiersRepoInterface interface {
	AddLegalEntityIdentifier(ctx context.Context, req legal_entity_models.LegalEntityIdentifier) error
	FindAllIdentitifiersOfLegalEntity(ctx context.Context, legalEntityID uuid.UUID) ([]legal_entity_models.LegalEntityIdentifier, error)
	GetIdentifiersByID(ctx context.Context, identifierID uuid.UUID) (*legal_entity_models.LegalEntityIdentifier, error)
}

/* ---------------------------------- <***> --------------------------------- */

type LegalEntityContactsRepoInterface interface {
	AddLegalEntityContacts(ctx context.Context, req legal_entity_models.LegalEntityContact) error
	FindAddAllContactsOfLegalEntity(ctx context.Context, legalEntityID uuid.UUID) ([]legal_entity_models.LegalEntityContact, error)
}

/* ---------------------------------- <***> --------------------------------- */

type LegalEntityTaxRegistrationRepoInterface interface {
	AddNewTaxRegistration(ctx context.Context, req *legal_entity_models.LegalEntityTaxRegistration) error
	FindAllTaxRegistrationOfLegalEntity(ctx context.Context, legalEntityID uuid.UUID) ([]legal_entity_models.LegalEntityTaxRegistration, error)
}

/* ---------------------------------- <***> --------------------------------- */

type LegalEntityTaxCategoryRepoInterface interface {

	// --- Tax Categories (The "What": FOOD, ROOM, LIQUOR) ---
	CreateTaxCategory(ctx context.Context, category *legal_entity_models.LegalEntityTaxCategory) error
	ListTaxCategories(ctx context.Context, entityID uuid.UUID) ([]legal_entity_models.LegalEntityTaxCategory, error)
	GetTaxCategoryByID(ctx context.Context, taxCategoryID uuid.UUID) (legal_entity_models.LegalEntityTaxCategory, error)
	GetTaxCategoryByCode(ctx context.Context, entityID uuid.UUID, code string) (*legal_entity_models.LegalEntityTaxCategory, error)

	// --- Tax Rules & Components (The "How Much") ---
	// Atomic creation of a rule and its constituent parts (CGST, SGST, etc.)
	CreateTaxRuleOfTaxCategory(ctx context.Context, rule *legal_entity_models.LegalEntityTaxRule) error
	CreateTaxComponentOfTaxRule(ctx context.Context, tax_components *legal_entity_models.LegalEntityTaxComponent) error

	// Fetching active rules for the Accounting Engine to calculate tax
	GetActiveRulesByCategory(ctx context.Context, categoryID uuid.UUID) ([]legal_entity_models.LegalEntityTaxRule, error)
	FindAllTaxCategoryRules(ctx context.Context, taxCategoryID uuid.UUID) ([]legal_entity_models.LegalEntityTaxRule, error)
	GetTaxRuleComponents(ctx context.Context, ruleID uuid.UUID) ([]legal_entity_models.LegalEntityTaxComponent, error)

	// --- Accounting Mapping (The "Where" in the Ledger) ---
	// Links a Tax Category to specific Ledger Accounts in a specific Book

	CreateTaxCategoryAccountMapping(ctx context.Context, mapping *legal_entity_models.TaxCategoryAccountMapping) error
	GetTaxCategoryAccountMappingByTaxCategoryID(ctx context.Context, LegalEntityTaxCategoryID uuid.UUID) (*legal_entity_models.TaxCategoryAccountMapping, error)

	// --- Maintenance ---
	ToggleRuleStatus(ctx context.Context, ruleID uuid.UUID, isActive bool) error
}
