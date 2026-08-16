package stockwaveEngineServices

import (
	"context"
	stockwaveEngineContracts "sharpnr-engine/stockwave_engine/contracts"
	stockwaveEngineDtos "sharpnr-engine/stockwave_engine/dtos"
	stockwaveEngineModels "sharpnr-engine/stockwave_engine/models"

	"github.com/google/uuid"
)

type InventoryEngineServicesInterface interface {
	AddNewInventoryAccount(ctx context.Context, orgID, emplID uuid.UUID, request *stockwaveEngineDtos.NewInventoryAccountRequestDto) error
	ListInventoryAccountOfOrg(ctx context.Context, orgID, emplID uuid.UUID) ([]stockwaveEngineModels.InventoryAccount, error)
	GetInventoryAccountID(ctx context.Context, orgID, emplID, inventoryAccountID uuid.UUID) (*stockwaveEngineModels.InventoryAccount, error)

	InventoryServiceInterface
	InventoryItemServiceInterface
	InventoryItemPriceServiceInterface
	InventoryItemHistoryServiceInterface
}

type InventoryServiceInterface interface {
	AddNewInventory(ctx context.Context, orgID, emplID, inventoryAccountID uuid.UUID, request *stockwaveEngineDtos.NewInventoryRequestDto) error
	ListAllInventoryOfInventoryAccount(ctx context.Context, orgID, emplID, inventoryAccountID uuid.UUID) ([]stockwaveEngineModels.Inventory, error)
	GetAllSharedInventoryAccount(ctx context.Context, orgID, emplID uuid.UUID) ([]stockwaveEngineModels.InventoryAccount, error)
	GetInventoriesOfBusinessUnit(ctx context.Context, orgID, emplID, businessUnitID uuid.UUID, inventoryAccountID *uuid.UUID) ([]stockwaveEngineModels.Inventory, error)
	GetInventoriesOfLegalEntity(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) ([]stockwaveEngineModels.Inventory, error)
	GetInventoryByID(ctx context.Context, orgID, emplID, inventoryID uuid.UUID) (*stockwaveEngineModels.Inventory, error)
}

type InventoryItemServiceInterface interface {
	AddItemsInInventory(ctx context.Context, input stockwaveEngineContracts.AddItemsInInventoryInput) error
}

type InventoryItemPriceServiceInterface interface {
}

type InventoryItemHistoryServiceInterface interface {
}
