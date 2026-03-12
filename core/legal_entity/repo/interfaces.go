package legal_entity_repo

import (
	"context"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"

	"github.com/google/uuid"
)

type OrgLegalEntityRepoInterface interface {
	CreateLegalEntity(ctx context.Context, legalEntity *legal_entity_models.LegalEntity) error
	GetLegalEntityByOrgID(ctx context.Context, orgID uuid.UUID) ([]*legal_entity_models.LegalEntity, error)
	GetLegalEntityByID(ctx context.Context, legalEntityID uuid.UUID) (*legal_entity_models.LegalEntity, error)
	UpdateLegalEntityStatus(ctx context.Context, legalEntityID uuid.UUID, isActive bool) error

	AddAddressInLegalEntity(ctx context.Context, address *legal_entity_models.LegalEntityAddress) error
	GetAddressOfLegalEntity(ctx context.Context, legalEntityID uuid.UUID) (*legal_entity_models.LegalEntityAddress, error)
	UpdateLegalEntityAddress(ctx context.Context, addressID uuid.UUID, updates map[string]interface{}) error

	CreateLegalEntityCompliance(ctx context.Context, compliance *legal_entity_models.LegalEntityCompliance) error
	GetLegalCompliancesByLegalEntityID(ctx context.Context, legalEntityID uuid.UUID) ([]*legal_entity_models.LegalEntityCompliance, error)

	// Binding Tables
	AddPropertyInLegalEntity(ctx context.Context, propertyLegalEntityRelation *legal_entity_models.PropertyLegalEntity) error
	LegalEntityTaxRepoInterface
}

type LegalEntityTaxRepoInterface interface {

	// --- Tax Registration (GSTIN/VAT details) ---
	CreateTaxRegistration(ctx context.Context, reg *legal_entity_models.LegalEntityTaxRegistration) error
	GetTaxRegistrations(ctx context.Context, entityID uuid.UUID) ([]legal_entity_models.LegalEntityTaxRegistration, error)

	// --- Tax Categories (The "What": FOOD, ROOM, LIQUOR) ---
	CreateTaxCategory(ctx context.Context, category *legal_entity_models.LegalEntityTaxCategory) error
	ListTaxCategories(ctx context.Context, entityID uuid.UUID) ([]legal_entity_models.LegalEntityTaxCategory, error)
	GetTaxCategoryByCode(ctx context.Context, entityID uuid.UUID, code string) (*legal_entity_models.LegalEntityTaxCategory, error)

	// --- Tax Rules & Components (The "How Much") ---
	// Atomic creation of a rule and its constituent parts (CGST, SGST, etc.)
	CreateTaxRuleOfTaxCategory(
		ctx context.Context,
		rule *legal_entity_models.LegalEntityTaxRule,
	) error

	CreateTaxComponentOfTaxRule(
		ctx context.Context,
		tax_components *legal_entity_models.LegalEntityTaxComponent,
	) error

	// Fetching active rules for the Accounting Engine to calculate tax
	GetActiveRulesByCategory(ctx context.Context, categoryID uuid.UUID) ([]legal_entity_models.LegalEntityTaxRule, error)
	GetTaxRuleComponents(ctx context.Context, ruleID uuid.UUID) ([]legal_entity_models.LegalEntityTaxComponent, error)

	// --- Accounting Mapping (The "Where" in the Ledger) ---
	// Links a Tax Category to specific Ledger Accounts in a specific Book

	CreateTaxCategoryAccountMapping(ctx context.Context, mapping *legal_entity_models.TaxCategoryAccountMapping) error
	GetTaxCategoryAccountMappingByTaxCategoryID(ctx context.Context, LegalEntityTaxCategoryID uuid.UUID) (*legal_entity_models.TaxCategoryAccountMapping, error)

	// --- Maintenance ---
	ToggleRuleStatus(ctx context.Context, ruleID uuid.UUID, isActive bool) error
	// DeleteTaxRule(ctx context.Context, ruleID uuid.UUID) error // Use carefully for audit trails
}
