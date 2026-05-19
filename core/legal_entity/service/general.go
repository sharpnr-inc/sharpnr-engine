package legal_entity_service

import (
	"context"
	"fmt"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/shared/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Create a new Legal Entity for the organization
// and returns the LegalEntityID
func (r *orgLegalEntity) CreateNewLegalEntity(
	ctx context.Context,
	orgID, emplID uuid.UUID,
	input legal_entity_contracts.NewLegalEntityRequestDto,
) (legalEntityID uuid.UUID, err error) {

	err = r.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		incorporationDate, err := utils.ConvertDate(input.IncorporationDate)
		if err != nil {
			return err
		}

		legalEntity := &legal_entity_models.LegalEntity{
			ID:                uuid.New(),
			OrganizationID:    orgID,
			LegalName:         input.LegalName,
			DisplayName:       input.DisplayName,
			EntityType:        input.Type,
			IsActive:          false,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			FYStartMonth:      input.FiscalYearMonth,
			IncorporationDate: &incorporationDate,
		}

		if err := r.legalEntityRepo.CreateLegalEntity(txCtx, legalEntity); err != nil {
			return err
		}

		if len(input.Addresses) > 0 {
			for _, a := range input.Addresses {
				_, err = r.AddNewAddressInLegalEntity(txCtx, orgID, emplID, legalEntity.ID, a)
				if err != nil {
					return err
				}
			}
		}

		legalEntityID = legalEntity.ID
		return nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	return legalEntityID, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *orgLegalEntity) GetLegalEntityDetails(ctx context.Context, orgID, emplId, legalEntityID uuid.UUID) (*legal_entity_contracts.LegalEntityResponseDto, error) {
	var result *legal_entity_contracts.LegalEntityResponseDto

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// 1. Fetch the core Legal Entity details
		le, err := s.legalEntityRepo.GetLegalEntityByID(txCtx, legalEntityID)
		if err != nil {
			return err
		}

		// 2. Fetch the associated Address
		leAddr, err := s.legalEntityRepo.GetAddressesOfLegalEntity(txCtx, le.ID)
		if err != nil {
			return err
		}

		fmt.Println(leAddr)

		// 3. Map the data to the contract
		result = &legal_entity_contracts.LegalEntityResponseDto{
			ID:          le.ID,
			DisplayName: le.LegalName,
			EntityType:  le.EntityType,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *orgLegalEntity) GetLegalEntityByID(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) (*legal_entity_contracts.LegalEntityResponseDto, error) {

	legalEntity, err := s.legalEntityRepo.GetLegalEntityByID(ctx, legalEntityID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
	}

	return &legal_entity_contracts.LegalEntityResponseDto{
		ID:          legalEntity.ID,
		DisplayName: legalEntity.DisplayName,
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *orgLegalEntity) GetLegalEntityByOrgID(
	ctx context.Context,
	orgID uuid.UUID,
	emplID uuid.UUID,
) ([]legal_entity_contracts.LegalEntityResponseDto, error) {

	var results []legal_entity_contracts.LegalEntityResponseDto

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		// Fetch legal entities
		legalEntities, err := s.legalEntityRepo.GetLegalEntitiesByOrgID(txCtx, orgID)
		if err != nil {
			return err
		}

		results = make([]legal_entity_contracts.LegalEntityResponseDto, len(legalEntities))

		for i, entity := range legalEntities {
			var incDate time.Time
			if entity.IncorporationDate != nil {
				incDate = *entity.IncorporationDate
			}

			results[i] = legal_entity_contracts.LegalEntityResponseDto{
				ID:                entity.ID,
				DisplayName:       entity.DisplayName,
				LegalName:         entity.LegalName,
				Status:            entity.IsActive, // Maps IsActive to Status
				EntityType:        entity.EntityType,
				IncorporationDate: incDate,
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil
}

/* ---------------------------------- <***> --------------------------------- */
