package inventory_engine_services

import (
	"context"
	inventory_dtos "sharpnr-engine/service_engines/inventory_engine/dtos"
	inventory_engine_models "sharpnr-engine/service_engines/inventory_engine/models"

	"github.com/google/uuid"
)

type InventoryEngineServicesInterface interface {
	AddNewInventoryAccount(ctx context.Context, orgID, emplID uuid.UUID, request *inventory_dtos.NewInventoryAccountRequestDto) error
	ListInventoryAccountOfOrg(ctx context.Context, orgID, emplID uuid.UUID) ([]inventory_engine_models.InventoryAccount, error)

	AddNewInventory(ctx context.Context, orgID, emplID, inventoryAccountID uuid.UUID, request *inventory_dtos.NewInventoryRequestDto) error
	ListAllInventoryOfInventoryAccount(ctx context.Context, orgID, emplID, inventoryAccountID uuid.UUID) ([]inventory_engine_models.Inventory, error)
}
