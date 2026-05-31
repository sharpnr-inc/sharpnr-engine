package inventory_engine_services

import (
	"context"
	dbtx "sharpnr-engine/pkg/dbTx"
	inventory_dtos "sharpnr-engine/sharp_inventory_engine/dtos"
	inventory_engine_models "sharpnr-engine/sharp_inventory_engine/models"

	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *inventoryEngine) AddNewInventoryAccount(ctx context.Context, orgID, emplID uuid.UUID, request *inventory_dtos.NewInventoryAccountRequestDto) error {
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		newInventoryAcc := &inventory_engine_models.InventoryAccount{
			ID:             uuid.New(),
			OrganizationID: orgID,
			Name:           request.Name,
			Type:           request.Type,
			IsShared:       request.IsShared,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		if err := s.repo.CreateInventoryAccount(txCtx, newInventoryAcc); err != nil {
			return err
		}

		return nil
	})
}

func (s *inventoryEngine) ListInventoryAccountOfOrg(ctx context.Context, orgID, emplID uuid.UUID) ([]inventory_engine_models.InventoryAccount, error) {

	var accounts []inventory_engine_models.InventoryAccount

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)
		inventoryAccounts, err := s.repo.GetInventoryAccountsByOrgID(txCtx, orgID)
		if err != nil {
			return err
		}

		accounts = inventoryAccounts

		return nil
	})

	if err != nil {
		return nil, err
	}

	return accounts, err
}

func (s *inventoryEngine) GetAllSharedInventoryAccount(ctx context.Context, orgID, emplID uuid.UUID) ([]inventory_engine_models.InventoryAccount, error) {
	var accounts []inventory_engine_models.InventoryAccount

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)
		inventoryAccounts, err := s.repo.GetSharedInventoryAccountsByOrgID(txCtx, orgID)
		if err != nil {
			return err
		}

		accounts = inventoryAccounts

		return nil
	})

	if err != nil {
		return nil, err
	}

	return accounts, err
}
