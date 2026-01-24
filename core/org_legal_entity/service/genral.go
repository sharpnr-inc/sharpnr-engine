package org_legal_entity_service

import (
	"context"
	org_legal_entity_contracts "sharpnr-engine/core/org_legal_entity/contracts"
	org_legal_entity_models "sharpnr-engine/core/org_legal_entity/models"
	"time"

	"github.com/google/uuid"
)

func (r *orgLegalEntity) CreateNewLegalEntity(ctx context.Context, input org_legal_entity_contracts.NewLegalEntityInput) error {

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
		return err
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
		return err
	}
	return nil
}
