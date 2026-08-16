package stockwaveEngineRepo

import (
	"context"

	stockwaveEngineConstants "sharpnr-engine/stockwave_engine/constants"
	stockwaveEngineContracts "sharpnr-engine/stockwave_engine/contracts"
	stockwaveEngineModels "sharpnr-engine/stockwave_engine/models"

	"time"

	"github.com/google/uuid"
)

type InventoryEngineRepositoryInterface interface {
	CreateInventoryAccount(ctx context.Context, account *stockwaveEngineModels.InventoryAccount) error
	GetInventoryAccountsByOrgID(ctx context.Context, orgID uuid.UUID) ([]stockwaveEngineModels.InventoryAccount, error)
	GetSharedInventoryAccountsByOrgID(ctx context.Context, orgID uuid.UUID) ([]stockwaveEngineModels.InventoryAccount, error)
	GetInventoryAccountByID(ctx context.Context, inventoryAccountID uuid.UUID) (*stockwaveEngineModels.InventoryAccount, error)
	UpdateInventoryAccount(ctx context.Context, inventoryAccountID uuid.UUID, updatesdto stockwaveEngineContracts.UpdateInventoryAccountParams) error
	DeleteInventoryAccount(ctx context.Context, inventoryAccountID uuid.UUID) error

	InventoryRepositoryInterface
	InventoryItemsRepositoryInterface
	InventoryItemPriceRepositoryInterface
	InventoryItemHistoryRepositoryInterface
	InventoryPurchaseOrderRepositoryInterface
}

/* ----------------------------------- Inventory ---------------------------------- */

type InventoryRepositoryInterface interface {
	CreateInventory(ctx context.Context, inventory *stockwaveEngineModels.Inventory) error
	GetInventoryByID(ctx context.Context, inventoryID uuid.UUID) (*stockwaveEngineModels.Inventory, error)
	GetInventoryByBusinessUnitID(ctx context.Context, businessUnitID uuid.UUID, inventoryAccountID *uuid.UUID) ([]stockwaveEngineModels.Inventory, error)
	GetInventoryByLegalEntityID(ctx context.Context, legalEntityID uuid.UUID) ([]stockwaveEngineModels.Inventory, error)
	GetInventoryByInventoryAccountID(ctx context.Context, inventoryAccountID uuid.UUID) ([]stockwaveEngineModels.Inventory, error)
	UpdateInventory(ctx context.Context, inventoryID uuid.UUID, params stockwaveEngineContracts.UpdateInventoryParams) error
	DeleteInventory(ctx context.Context, inventoryID uuid.UUID) error
}

/* ----------------------------------- InventoryItem ---------------------------------- */

type InventoryItemsRepositoryInterface interface {
	CreateInventoryItem(ctx context.Context, item *stockwaveEngineModels.InventoryItem) error
	GetInventoryItemByID(ctx context.Context, itemID uuid.UUID) (*stockwaveEngineModels.InventoryItem, error)
	GetInventoryItemsByInventoryID(ctx context.Context, inventoryID uuid.UUID) ([]stockwaveEngineModels.InventoryItem, error)
	GetInventoryItemBySKU(ctx context.Context, sku string) (*stockwaveEngineModels.InventoryItem, error)
	GetLowStockItemsByInventoryID(ctx context.Context, inventoryID uuid.UUID) ([]stockwaveEngineModels.InventoryItem, error) // WHERE quantity <= min_stock_level
	UpdateInventoryItem(ctx context.Context, itemID uuid.UUID, params stockwaveEngineContracts.UpdateInventoryItemParams) error
	AdjustInventoryItemQuantity(ctx context.Context, itemID uuid.UUID, delta float64) error // atomic: quantity = quantity + delta
	DeleteInventoryItem(ctx context.Context, itemID uuid.UUID) error
}

/* ----------------------------------- InventoryItemPrice ---------------------------------- */

type InventoryItemPriceRepositoryInterface interface {
	CreateInventoryItemPrice(ctx context.Context, price *stockwaveEngineModels.InventoryItemPrice) error
	GetInventoryItemPriceByID(ctx context.Context, priceID uuid.UUID) (*stockwaveEngineModels.InventoryItemPrice, error)
	GetInventoryItemPricesByItemID(ctx context.Context, itemID uuid.UUID) ([]stockwaveEngineModels.InventoryItemPrice, error)                                                        // full history
	GetActiveInventoryItemPrice(ctx context.Context, itemID uuid.UUID, priceType stockwaveEngineConstants.InventoryItemPriceType) (*stockwaveEngineModels.InventoryItemPrice, error) // WHERE effective_to IS NULL
	GetActiveInventoryItemPriceBySupplier(ctx context.Context, itemID uuid.UUID, supplierID uuid.UUID, priceType stockwaveEngineConstants.InventoryItemPriceType) (*stockwaveEngineModels.InventoryItemPrice, error)
	DeactivateInventoryItemPrice(ctx context.Context, priceID uuid.UUID, effectiveTo time.Time) error // SET effective_to = $2, never DELETE
}

/* ----------------------------------- InventoryItemHistory ---------------------------------- */

type InventoryItemHistoryRepositoryInterface interface {
	CreateInventoryItemHistory(ctx context.Context, history *stockwaveEngineModels.InventoryItemHistory) error
	GetInventoryItemHistoryByItemID(ctx context.Context, itemID uuid.UUID) ([]stockwaveEngineModels.InventoryItemHistory, error)
	GetInventoryItemHistoryByItemIDAndType(ctx context.Context, itemID uuid.UUID, changeType string) ([]stockwaveEngineModels.InventoryItemHistory, error)
	GetInventoryItemHistoryByReferenceID(ctx context.Context, referenceID uuid.UUID) ([]stockwaveEngineModels.InventoryItemHistory, error) // trace back to PO / consumption
	GetInventoryItemHistoryByDateRange(ctx context.Context, itemID uuid.UUID, from, to time.Time) ([]stockwaveEngineModels.InventoryItemHistory, error)
}

/* ----------------------------------- PurchaseOrder (stub) ---------------------------------- */

type InventoryPurchaseOrderRepositoryInterface interface{}
