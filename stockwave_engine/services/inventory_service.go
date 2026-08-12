package stockwaveEngineServices

import (
	"context"
	dbtx "sharpnr-engine/pkg/dbTx"
	stockwaveEngineDtos "sharpnr-engine/stockwave_engine/dtos"
	stockwaveEngineModels "sharpnr-engine/stockwave_engine/models"

	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *inventoryEngine) AddNewInventory(ctx context.Context, orgID, emplID, inventoryAccountID uuid.UUID, request *stockwaveEngineDtos.NewInventoryRequestDto) error {
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		newInventory := &stockwaveEngineModels.Inventory{
			ID:                 uuid.New(),
			InventoryAccountID: inventoryAccountID,
			BusinessUnitID:     *request.BusinessUnitID,
			LegalEntityID:      *request.LegalEntityID,
			Description:        request.Description,
			Location:           request.Location,
			Status:             request.Status,
			Name:               request.Name,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}

		if err := s.repo.CreateInventory(txCtx, newInventory); err != nil {
			return err
		}

		return nil
	})
}

/* ---------------------------------- <***> --------------------------------- */

func (s *inventoryEngine) ListAllInventoryOfInventoryAccount(ctx context.Context, orgID, emplID, inventoryAccountID uuid.UUID) ([]stockwaveEngineModels.Inventory, error) {
	var inventories []stockwaveEngineModels.Inventory

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// TODO: Optional Security Check
		// if err := s.verifyAccountOwnership(txCtx, orgID, inventoryAccountID); err != nil {
		//     return err
		// }

		res, err := s.repo.GetInventoryByInventoryAccountID(txCtx, inventoryAccountID)
		if err != nil {
			return err
		}

		inventories = res
		return nil
	})

	if err != nil {
		return nil, err // Return nil on failure
	}

	return inventories, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *inventoryEngine) GetInventoriesOfLegalEntity(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, legalEntityID uuid.UUID) ([]stockwaveEngineModels.Inventory, error) {
	var inventories []stockwaveEngineModels.Inventory

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		res, err := s.repo.GetInventoryByLegalEntityID(txCtx, legalEntityID)

		if err != nil {
			return err
		}

		inventories = res
		return nil

	})

	if err != nil {
		return nil, err // Return nil on failure
	}

	return inventories, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *inventoryEngine) GetInventoriesOfBusinessUnit(ctx context.Context, orgID, emplID, businessUnitID uuid.UUID, inventoryAccountID *uuid.UUID) ([]stockwaveEngineModels.Inventory, error) {

	var result []stockwaveEngineModels.Inventory
	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		inventories, err := s.repo.GetInventoryByBusinessUnitID(txCtx, businessUnitID, inventoryAccountID)
		if err != nil {
			return nil
		}

		result = inventories
		return nil
	})

	if err != nil {
		return result, err
	}

	return result, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *inventoryEngine) GetInventoryItems(ctx context.Context) error {
	panic("")
}
