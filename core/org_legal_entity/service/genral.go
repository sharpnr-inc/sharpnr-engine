package org_legal_entity_service

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	org_legal_entity_contracts "sharpnr-engine/core/org_legal_entity/contracts"
	org_legal_entity_models "sharpnr-engine/core/org_legal_entity/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *orgLegalEntity) CreateNewLegalEntity(
	ctx context.Context,
	input org_legal_entity_contracts.NewLegalEntityInput,
) (legalEntityID uuid.UUID, err error) {

	err = r.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		legalEntity := &org_legal_entity_models.OrganizationLegalEntity{
			ID:             uuid.New(),
			OrganizationID: input.OrgID,
			LegalName:      input.LegalName,
			EntityType:     input.Type,
			IsActive:       false,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		if err := r.legalEntityRepo.CreateLegalEntity(txCtx, legalEntity); err != nil {
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

		if err := r.legalEntityRepo.AddAddressInLegalEntity(txCtx, legEntAddr); err != nil {
			return err
		}

		legalEntityID = legalEntity.ID
		return nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	return legalEntityID, nil
}

func (s *orgLegalEntity) GetLegalEntityByOrgID(
	ctx context.Context,
	orgID uuid.UUID,
) ([]*org_legal_entity_contracts.LegalEntityReader, error) {

	var results []*org_legal_entity_contracts.LegalEntityReader

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		// Fetch legal entities
		legalEntities, err := s.legalEntityRepo.GetLegalEntityByOrgID(txCtx, orgID)
		if err != nil {
			return err
		}

		results = make([]*org_legal_entity_contracts.LegalEntityReader, len(legalEntities))

		for i, le := range legalEntities {

			leAddr, err := s.legalEntityRepo.GetAddressOfLegalEntity(txCtx, le.ID)
			if err != nil {
				return err
			}

			results[i] = &org_legal_entity_contracts.LegalEntityReader{
				LegalEntityID: le.ID,
				Name:          le.LegalName,
				EntityType:    le.EntityType,
				GST:           le.GSTIN,
				CIN:           le.CIN,
				PAN:           le.PAN,
				Country:       leAddr.Country,
				AddedAt:       le.CreatedAt,
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil
}
