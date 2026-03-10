package org_legal_entity_repo

import (
	"context"
	org_legal_entity_models "sharpnr-engine/core/org_legal_entity/models"

	"github.com/google/uuid"
)

type OrgLegalEntityRepoInterface interface {
	CreateLegalEntity(ctx context.Context, legalEntity *org_legal_entity_models.OrganizationLegalEntity) error
	GetLegalEntityByOrgID(ctx context.Context, orgID uuid.UUID) ([]*org_legal_entity_models.OrganizationLegalEntity, error)
	GetLegalEntityByID(ctx context.Context, legalEntityID uuid.UUID) (*org_legal_entity_models.OrganizationLegalEntity, error)
	UpdateLegalEntityStatus(ctx context.Context, legalEntityID uuid.UUID, isActive bool) error

	AddAddressInLegalEntity(ctx context.Context, address *org_legal_entity_models.LegalEntityAddress) error
	GetAddressOfLegalEntity(ctx context.Context, legalEntityID uuid.UUID) (*org_legal_entity_models.LegalEntityAddress, error)
	UpdateLegalEntityAddress(ctx context.Context, addressID uuid.UUID, updates map[string]interface{}) error

	CreateLegalEntityCompliance(ctx context.Context, compliance *org_legal_entity_models.LegalEntityCompliance) error
	GetLegalCompliancesByLegalEntityID(ctx context.Context, legalEntityID uuid.UUID) ([]*org_legal_entity_models.LegalEntityCompliance, error)

	// Binding Tables
	AddPropertyInLegalEntity(ctx context.Context, propertyLegalEntityRelation *org_legal_entity_models.PropertyLegalEntity) error
	LegalEntityTaxRepoInterface
}

type LegalEntityTaxRepoInterface interface {

	// --- Tax Registration (GSTIN/VAT details) ---
	CreateTaxRegistration(ctx context.Context, reg *org_legal_entity_models.LegalEntityTaxRegistration) error
	GetTaxRegistrations(ctx context.Context, entityID uuid.UUID) ([]org_legal_entity_models.LegalEntityTaxRegistration, error)

	// --- Tax Categories (The "What": FOOD, ROOM, LIQUOR) ---
	CreateTaxCategory(ctx context.Context, category *org_legal_entity_models.LegalEntityTaxCategory) error
	ListTaxCategories(ctx context.Context, entityID uuid.UUID) ([]org_legal_entity_models.LegalEntityTaxCategory, error)
	GetTaxCategoryByCode(ctx context.Context, entityID uuid.UUID, code string) (*org_legal_entity_models.LegalEntityTaxCategory, error)

	// --- Tax Rules & Components (The "How Much") ---
	// Atomic creation of a rule and its constituent parts (CGST, SGST, etc.)
	CreateTaxRuleWithComponents(
		ctx context.Context,
		rule *org_legal_entity_models.LegalEntityTaxRule,
		components []org_legal_entity_models.LegalEntityTaxComponent,
	) error

	// Fetching active rules for the Accounting Engine to calculate tax
	GetActiveRulesByCategory(ctx context.Context, categoryID uuid.UUID) ([]org_legal_entity_models.LegalEntityTaxRule, error)
	GetRuleComponents(ctx context.Context, ruleID uuid.UUID) ([]org_legal_entity_models.LegalEntityTaxComponent, error)

	// --- Accounting Mapping (The "Where" in the Ledger) ---
	// Links a Tax Category to specific Ledger Accounts in a specific Book
	UpsertAccountingMapping(ctx context.Context, mapping *org_legal_entity_models.LegalEntityAccountingMapping) error
	GetMappingByCategory(ctx context.Context, categoryID uuid.UUID, bookID uuid.UUID) (*org_legal_entity_models.LegalEntityAccountingMapping, error)

	// --- Maintenance ---
	ToggleRuleStatus(ctx context.Context, ruleID uuid.UUID, isActive bool) error
	// DeleteTaxRule(ctx context.Context, ruleID uuid.UUID) error // Use carefully for audit trails
}
