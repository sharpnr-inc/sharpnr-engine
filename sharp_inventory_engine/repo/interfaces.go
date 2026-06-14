package inventory_engine_repo

import (
	"context"

	inventory_constants "sharpnr-engine/sharp_inventory_engine/constants"
	inventory_engine_contracts "sharpnr-engine/sharp_inventory_engine/contracts"
	inventory_engine_models "sharpnr-engine/sharp_inventory_engine/models"

	"time"

	"github.com/google/uuid"
)

type InventoryEngineRepositoryInterface interface {
	CreateInventoryAccount(ctx context.Context, account *inventory_engine_models.InventoryAccount) error
	GetInventoryAccountsByOrgID(ctx context.Context, orgID uuid.UUID) ([]inventory_engine_models.InventoryAccount, error)
	GetSharedInventoryAccountsByOrgID(ctx context.Context, orgID uuid.UUID) ([]inventory_engine_models.InventoryAccount, error)
	GetInventoryAccountByID(ctx context.Context, inventoryAccountID uuid.UUID) (*inventory_engine_models.InventoryAccount, error)
	UpdateInventoryAccount(ctx context.Context, inventoryAccountID uuid.UUID, updatesdto inventory_engine_contracts.UpdateInventoryAccountParams) error
	DeleteInventoryAccount(ctx context.Context, inventoryAccountID uuid.UUID) error

	InventoryRepositoryInterface
	InventoryItemsRepositoryInterface
	InventoryItemPriceRepositoryInterface
	InventoryItemHistoryRepositoryInterface
	InventoryPurchaseOrderRepositoryInterface
}

/* ----------------------------------- Inventory ---------------------------------- */

type InventoryRepositoryInterface interface {
	CreateInventory(ctx context.Context, inventory *inventory_engine_models.Inventory) error
	GetInventoryByID(ctx context.Context, inventoryID uuid.UUID) (*inventory_engine_models.Inventory, error)
	GetInventoryByBusinessUnitID(ctx context.Context, businessUnitID uuid.UUID, inventoryAccountID *uuid.UUID) ([]inventory_engine_models.Inventory, error)
	GetInventoryByLegalEntityID(ctx context.Context, legalEntityID uuid.UUID) ([]inventory_engine_models.Inventory, error)
	GetInventoryByInventoryAccountID(ctx context.Context, inventoryAccountID uuid.UUID) ([]inventory_engine_models.Inventory, error)
	UpdateInventory(ctx context.Context, inventoryID uuid.UUID, params inventory_engine_contracts.UpdateInventoryParams) error
	DeleteInventory(ctx context.Context, inventoryID uuid.UUID) error
}

/* ----------------------------------- InventoryItem ---------------------------------- */

type InventoryItemsRepositoryInterface interface {
	CreateInventoryItem(ctx context.Context, item *inventory_engine_models.InventoryItem) error
	GetInventoryItemByID(ctx context.Context, itemID uuid.UUID) (*inventory_engine_models.InventoryItem, error)
	GetInventoryItemsByInventoryID(ctx context.Context, inventoryID uuid.UUID) ([]inventory_engine_models.InventoryItem, error)
	GetInventoryItemBySKU(ctx context.Context, sku string) (*inventory_engine_models.InventoryItem, error)
	GetLowStockItemsByInventoryID(ctx context.Context, inventoryID uuid.UUID) ([]inventory_engine_models.InventoryItem, error) // WHERE quantity <= min_stock_level
	UpdateInventoryItem(ctx context.Context, itemID uuid.UUID, params inventory_engine_contracts.UpdateInventoryItemParams) error
	AdjustInventoryItemQuantity(ctx context.Context, itemID uuid.UUID, delta float64) error // atomic: quantity = quantity + delta
	DeleteInventoryItem(ctx context.Context, itemID uuid.UUID) error
}

/* ----------------------------------- InventoryItemPrice ---------------------------------- */

type InventoryItemPriceRepositoryInterface interface {
	CreateInventoryItemPrice(ctx context.Context, price *inventory_engine_models.InventoryItemPrice) error
	GetInventoryItemPriceByID(ctx context.Context, priceID uuid.UUID) (*inventory_engine_models.InventoryItemPrice, error)
	GetInventoryItemPricesByItemID(ctx context.Context, itemID uuid.UUID) ([]inventory_engine_models.InventoryItemPrice, error)                                                   // full history
	GetActiveInventoryItemPrice(ctx context.Context, itemID uuid.UUID, priceType inventory_constants.InventoryItemPriceType) (*inventory_engine_models.InventoryItemPrice, error) // WHERE effective_to IS NULL
	GetActiveInventoryItemPriceBySupplier(ctx context.Context, itemID uuid.UUID, supplierID uuid.UUID, priceType inventory_constants.InventoryItemPriceType) (*inventory_engine_models.InventoryItemPrice, error)
	DeactivateInventoryItemPrice(ctx context.Context, priceID uuid.UUID, effectiveTo time.Time) error // SET effective_to = $2, never DELETE
}

/* ----------------------------------- InventoryItemHistory ---------------------------------- */

type InventoryItemHistoryRepositoryInterface interface {
	CreateInventoryItemHistory(ctx context.Context, history *inventory_engine_models.InventoryItemHistory) error
	GetInventoryItemHistoryByItemID(ctx context.Context, itemID uuid.UUID) ([]inventory_engine_models.InventoryItemHistory, error)
	GetInventoryItemHistoryByItemIDAndType(ctx context.Context, itemID uuid.UUID, changeType string) ([]inventory_engine_models.InventoryItemHistory, error)
	GetInventoryItemHistoryByReferenceID(ctx context.Context, referenceID uuid.UUID) ([]inventory_engine_models.InventoryItemHistory, error) // trace back to PO / consumption
	GetInventoryItemHistoryByDateRange(ctx context.Context, itemID uuid.UUID, from, to time.Time) ([]inventory_engine_models.InventoryItemHistory, error)
}

/* ----------------------------------- PurchaseOrder (stub) ---------------------------------- */

type InventoryPurchaseOrderRepositoryInterface interface{}
