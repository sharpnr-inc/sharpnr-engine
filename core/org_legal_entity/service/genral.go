package org_legal_entity_service

import (
	"context"
	org_legal_entity_contracts "sharpnr-engine/core/org_legal_entity/contracts"
	org_legal_entity_models "sharpnr-engine/core/org_legal_entity/models"
	"time"

	"github.com/google/uuid"
)

func (r *orgLegalEntity) CreateNewLegalEntity(ctx context.Context, input org_legal_entity_contracts.NewLegalEntityInput) (legalEntityID uuid.UUID, err error) {

	legalEntity := &org_legal_entity_models.OrganizationLegalEntity{
		ID:             uuid.New(),
		OrganizationID: input.OrgID,
		LegalName:      input.LegalName,
		EntityType:     input.Type,
		IsActive:       false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := r.legalEntityRepo.CreateLegalEntity(ctx, legalEntity); err != nil {
		return uuid.Nil, err
	}

	legEntAddr := &org_legal_entity_models.LegalEntityAddress{
		ID:                        uuid.New(),
		OrganizationLegalEntityID: legalEntity.ID,
		Type:                      "OPERATION",
		AddressLine:               input.AddressLine,
		City:                      input.City,
		State:                     input.State,
		Pincode:                   input.Pincode,
		Country:                   input.Country,
		IsPrimary:                 true,
	}

	if err := r.legalEntityRepo.AddAddressInLegalEntity(ctx, legEntAddr); err != nil {
		return uuid.Nil, err
	}
	return legalEntity.ID, nil
}

func (s *orgLegalEntity) GetLegalEntityByOrgID(ctx context.Context, orgID uuid.UUID) ([]*org_legal_entity_contracts.LegalEntityReader, error) {
	// 1. Fetch from Repo
	legalEntities, err := s.legalEntityRepo.GetLegalEntityByOrgID(ctx, orgID)
	if err != nil {
		return nil, err
	}

	// 2. Initialize the result slice with the correct length
	results := make([]*org_legal_entity_contracts.LegalEntityReader, len(legalEntities))

	// 3. Map each repo model to the contract reader
	for i, le := range legalEntities {
		results[i] = &org_legal_entity_contracts.LegalEntityReader{
			LegalEntityID: le.ID,
			Name:          le.LegalName,
			EntityType:    le.EntityType,
			// ... map other fields like GST, PAN, etc.
		}
	}

	return results, nil
}
