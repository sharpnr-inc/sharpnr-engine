package inventory_engine_services

import (
	"context"
	inventory_engine_contracts "sharpnr-engine/sharp_inventory_engine/contracts"
	inventory_dtos "sharpnr-engine/sharp_inventory_engine/dtos"
	inventory_engine_models "sharpnr-engine/sharp_inventory_engine/models"

	"github.com/google/uuid"
)

type InventoryEngineServicesInterface interface {
	AddNewInventoryAccount(ctx context.Context, orgID, emplID uuid.UUID, request *inventory_dtos.NewInventoryAccountRequestDto) error
	ListInventoryAccountOfOrg(ctx context.Context, orgID, emplID uuid.UUID) ([]inventory_engine_models.InventoryAccount, error)
	GetInventoryAccountID(ctx context.Context, orgID, emplID, inventoryAccountID uuid.UUID) (*inventory_engine_models.InventoryAccount, error)

	InventoryServiceInterface
	InventoryItemServiceInterface
	InventoryItemPriceServiceInterface
	InventoryItemHistoryServiceInterface
}

type InventoryServiceInterface interface {
	AddNewInventory(ctx context.Context, orgID, emplID, inventoryAccountID uuid.UUID, request *inventory_dtos.NewInventoryRequestDto) error
	ListAllInventoryOfInventoryAccount(ctx context.Context, orgID, emplID, inventoryAccountID uuid.UUID) ([]inventory_engine_models.Inventory, error)
	GetAllSharedInventoryAccount(ctx context.Context, orgID, emplID uuid.UUID) ([]inventory_engine_models.InventoryAccount, error)
	GetInventoriesOfBusinessUnit(ctx context.Context, orgID, emplID, businessUnitID uuid.UUID, inventoryAccountID *uuid.UUID) ([]inventory_engine_models.Inventory, error)
	GetInventoriesOfLegalEntity(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) ([]inventory_engine_models.Inventory, error)
	GetInventoryByID(ctx context.Context, orgID, emplID, inventoryID uuid.UUID) (*inventory_engine_models.Inventory, error)
}

type InventoryItemServiceInterface interface {
	AddItemsInInventory(ctx context.Context, input inventory_engine_contracts.AddItemsInInventoryInput) error
}

type InventoryItemPriceServiceInterface interface {
}

type InventoryItemHistoryServiceInterface interface {
}
