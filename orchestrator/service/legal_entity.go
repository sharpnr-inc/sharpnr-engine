package orchestrator_service

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	accounting_engine_contracts "sharpnr-engine/service_engines/accounting_engine/contracts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (c *orchestratorService) GetAllLegalEntity(ctx context.Context, orgID, emplID uuid.UUID) ([]legal_entity_contracts.LegalEntityResponseDto, error) {

	les, err := c.legalEntityService.GetLegalEntityByOrgID(ctx, orgID)
	if err != nil {
		return nil, err
	}

	return les, nil

}

/* ---------------------------------- <***> --------------------------------- */

func (c *orchestratorService) CreateLegalEntity(ctx context.Context, orgID, emplID uuid.UUID, req *legal_entity_contracts.NewLegalEntityRequestDto) error {

	return c.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		legalEntityID, err := c.legalEntityService.CreateNewLegalEntity(txCtx, orgID, emplID, *req)
		if err != nil {
			return err
		}

		// Whenever a new Legal Entity is Created we create an Accounting books
		// with proper defaults Accounts and Accounts Groups

		accountInput := accounting_engine_contracts.DraftAccountingSetupInput{
			OrgID:               orgID,
			LegalEntityID:       legalEntityID,
			Name_of_legalEntity: req.LegalName,
		}
		if err := c.accountingEngineService.CreateDraftAccountingSetup(txCtx, &accountInput); err != nil {
			return err
		}

		return nil
	})
}

/* ---------------------------------- <***> --------------------------------- */
