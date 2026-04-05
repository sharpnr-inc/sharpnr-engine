package inventory_engine_repo

import (
	"context"
	inventory_engine_models "sharpnr-engine/service_engines/inventory_engine/models"

	"github.com/google/uuid"
)

type InventoryEngineRepositoryInterface interface {
	CreateInventoryAccount(ctx context.Context, inventoryAccont *inventory_engine_models.InventoryAccount) error
	GetInventoryAccountsByOrgID(ctx context.Context, orgID uuid.UUID) ([]inventory_engine_models.InventoryAccount, error)
	GetSharedInventoryAccountsByOrgID(ctx context.Context, orgID uuid.UUID) ([]inventory_engine_models.InventoryAccount, error)

	CreateInventory(ctx context.Context, inventory *inventory_engine_models.Inventory) error
	GetInventoryByBusinessUnitID(ctx context.Context, businessUnitID uuid.UUID, inventoryAccountID *uuid.UUID) ([]inventory_engine_models.Inventory, error)
	GetInventoryByLegalEntityID(ctx context.Context, legalEntityID uuid.UUID) ([]inventory_engine_models.Inventory, error)
	GetInventoryByID(ctx context.Context, inventoryID uuid.UUID) ([]inventory_engine_models.Inventory, error)
	GetInventoryByInventoryAccountID(ctx context.Context, inventoryAccountID uuid.UUID) ([]inventory_engine_models.Inventory, error)

	CreateInventoryItems(ctx context.Context, item *inventory_engine_models.InventoryItem) error
	GetInventoryItemsByID(ctx context.Context, itemID *inventory_engine_models.InventoryItem) (*inventory_engine_models.InventoryItem, error)
	GetInventoryItemsByInventoryID(ctx context.Context, inventoryID uuid.UUID) ([]inventory_engine_models.InventoryItem, error)
}

