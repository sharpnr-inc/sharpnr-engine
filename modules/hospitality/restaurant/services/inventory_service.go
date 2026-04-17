package restaurantservices

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	restaurant_contracts "sharpnr-engine/modules/hospitality/restaurant/contracts"
	inventory_dtos "sharpnr-engine/service_engines/inventory_engine/dtos"
	inventory_engine_models "sharpnr-engine/service_engines/inventory_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *restaurantService) GetAllInventoryAccounts(ctx context.Context, orgID, emplID uuid.UUID) ([]inventory_engine_models.InventoryAccount, error) {
	var inventoryAcconts []inventory_engine_models.InventoryAccount

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)
		accounts, err := s.inventoryEngine.GetAllSharedInventoryAccount(txCtx, orgID, emplID)
		if err != nil {
			return err
		}

		inventoryAcconts = accounts

		return nil
	})

	if err != nil {
		return inventoryAcconts, err
	}

	return inventoryAcconts, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *restaurantService) AddNewInventoryInRestaurant(ctx context.Context, orgID, emplID, restaurantID uuid.UUID, request *restaurant_contracts.AddInventoryRequestDto) error {
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		restaurantDetails, err := s.restaurantRepo.GetRestaurantByID(txCtx, restaurantID)
		if err != nil {
			return nil
		}

		mappedReq := &inventory_dtos.NewInventoryRequestDto{
			BusinessUnitID: &restaurantID,
			LegalEntityID:  &restaurantDetails.LegalEntityID,
			Name:           request.Name,
			Description:    request.Description,
			Location:       request.Location,
			Status:         "ACTIVE",
		}

		if err := s.inventoryEngine.AddNewInventory(txCtx, orgID, emplID, request.InventoryAccountID, mappedReq); err != nil {
			return err
		}

		return nil
	})
}

/* ---------------------------------- <***> --------------------------------- */

func (s *restaurantService) GetRestaurantAllInventories(ctx context.Context, orgID, emplID, restaurantID uuid.UUID, inventoryAccountID *uuid.UUID) ([]restaurant_contracts.RestuarantInventoriesResponseDto, error) {
	var inventories []restaurant_contracts.RestuarantInventoriesResponseDto

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// 1. Fetch data from the inventory engine
		invent, err := s.inventoryEngine.GetInventoriesOfBusinessUnit(txCtx, orgID, emplID, restaurantID, inventoryAccountID)
		if err != nil {
			return err
		}

		// 2. Map the domain models to Response DTOs
		for _, item := range invent {
			inventories = append(inventories, restaurant_contracts.RestuarantInventoriesResponseDto{
				ID:                 item.ID,
				InventoryAccountID: item.InventoryAccountID,
				BusinessUnitID:     item.BusinessUnitID,
				LegalEntityID:      item.LegalEntityID,
				Name:               item.Name,
				Description:        item.Description,
				Location:           item.Location,
				Status:             item.Status,
				CreatedAt:          item.CreatedAt,
				UpdatedAt:          item.UpdatedAt,
			})
		}

		return nil
	})

	if err != nil {
		return nil, err // Return nil slice on error for cleaner API handling
	}

	return inventories, nil
}
