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
}
