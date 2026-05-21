package orchestrator_service

import (
	"context"
	"errors"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	dbtx "sharpnr-engine/pkg/dbTx"
	accounting_engine_contracts "sharpnr-engine/service_engines/accounting_engine/contracts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (c *orchestratorService) GetAllLegalEntity(ctx context.Context, orgID, emplID uuid.UUID) ([]legal_entity_contracts.LegalEntityResponseDto, error) {

	var response []legal_entity_contracts.LegalEntityResponseDto

	err := c.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		les, err := c.legalEntityService.GetLegalEntityByOrgID(txCtx, orgID, emplID)
		if err != nil {
			return err
		}

		response = les

		return nil
	})

	if err != nil {
		return []legal_entity_contracts.LegalEntityResponseDto{}, err
	}

	return response, nil
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

func (s *orchestratorService) GetDetailsOfLegalEntity(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) (legal_entity_contracts.LegalEntityDetails, error) {
	// Simply invoke the downstream service directly using the request context.
	// If you implemented errgroup inside GetLegalEntityDetails, it will safely spawn sub-connections here.
	res, err := s.legalEntityService.GetLegalEntityDetails(ctx, orgID, emplID, legalEntityID)
	if err != nil {
		return legal_entity_contracts.LegalEntityDetails{}, err // Properly bubble up the error
	}

	return res, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *orchestratorService) AddContactsInLegalEntity(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
	req *legal_entity_contracts.CreateLegalEntityContactsRequestDto,
) error {
	if req == nil {
		return errors.New("request body should not be empty")
	}

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		return s.legalEntityService.AddNewContactInLegalEntity(
			txCtx,
			orgID,
			emplID,
			legalEntityID,
			req,
		)
	})
}
