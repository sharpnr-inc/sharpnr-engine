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

func (s *inventoryEngine) AddNewInventoryAccount(ctx context.Context, orgID, emplID uuid.UUID, request *stockwaveEngineDtos.NewInventoryAccountRequestDto) error {
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		newInventoryAcc := &stockwaveEngineModels.InventoryAccount{
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

func (s *inventoryEngine) ListInventoryAccountOfOrg(ctx context.Context, orgID, emplID uuid.UUID) ([]stockwaveEngineModels.InventoryAccount, error) {

	var accounts []stockwaveEngineModels.InventoryAccount

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

func (s *inventoryEngine) GetAllSharedInventoryAccount(ctx context.Context, orgID, emplID uuid.UUID) ([]stockwaveEngineModels.InventoryAccount, error) {
	var accounts []stockwaveEngineModels.InventoryAccount

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
