package org_legal_entity_repo

import (
	"context"
	org_legal_entity_models "sharpnr-engine/core/org_legal_entity/models"

	"github.com/google/uuid"
)

type OrgLegalEntityRepoInterface interface {

	// this method creates the legal entity
	CreateLegalEntity(
		ctx context.Context,
		legalEntity org_legal_entity_models.OrganizationLegalEntity,
	) error

	// This method returns the slices of all the entity
	// of organization if there is no legal entity it will
	// return empty [] slices
	GetLegalEntityByOrgID(
		ctx context.Context,
		orgID uuid.UUID,
	) ([]org_legal_entity_models.OrganizationLegalEntity, error)

	// this method return the legal entity By Id
	GetLegalEntityByID(
		ctx context.Context,
		legalEntityID uuid.UUID,
	) (org_legal_entity_models.OrganizationLegalEntity, error)

	AddAddressInLegalEntity(
		ctx context.Context,
	) error
}
