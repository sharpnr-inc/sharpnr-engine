package orchestrator_service

import (
	"context"
	"errors"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	dbtx "sharpnr-engine/pkg/dbTx"
	accounting_engine_contracts "sharpnr-engine/sharp_accounting_engine/accounting_engine/contracts"
	accountingdtos "sharpnr-engine/sharp_accounting_engine/accounting_engine/dto"
	accounting_service "sharpnr-engine/sharp_accounting_engine/accounting_engine/services"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// legalEntityOrchestrator implements LegalEntityOrchestrationServiceInterface.
// Creating a legal entity also provisions its draft accounting setup, which is
// why the accounting engine lives here too.
type legalEntityOrchestrator struct {
	txManager               dbtx.TxManager
	legalEntityService      legal_entity_service.OrgLegalEntityServiceInterface
	accountingEngineService accounting_service.AccountingServiceInterface
}

func (c *legalEntityOrchestrator) GetAllLegalEntity(ctx context.Context, orgID, emplID uuid.UUID) ([]legal_entity_contracts.LegalEntityResponseDto, error) {

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

func (c *legalEntityOrchestrator) CreateLegalEntity(ctx context.Context, orgID, emplID uuid.UUID, req *legal_entity_contracts.NewLegalEntityRequestDto) error {

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

func (s *legalEntityOrchestrator) GetDetailsOfLegalEntity(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) (legal_entity_contracts.LegalEntityDetails, error) {
	// Simply invoke the downstream service directly using the request context.
	// If you implemented errgroup inside GetLegalEntityDetails, it will safely spawn sub-connections here.
	res, err := s.legalEntityService.GetLegalEntityDetails(ctx, orgID, emplID, legalEntityID)
	if err != nil {
		return legal_entity_contracts.LegalEntityDetails{}, err // Properly bubble up the error
	}

	return res, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *legalEntityOrchestrator) AddContactsInLegalEntity(
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

/* ---------------------------------- <***> --------------------------------- */

func (s *legalEntityOrchestrator) AddAddressInLegalEntity(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
	req *legal_entity_contracts.CreateLegalEntityAddressRequest,
) error {
	if req == nil {
		return errors.New("request body should not be empty")
	}
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		_, err := s.legalEntityService.AddNewAddressInLegalEntity(
			txCtx,
			orgID,
			emplID,
			legalEntityID,
			*req,
		)

		if err != nil {
			return err
		}

		return nil
	})
}

/* ---------------------------------- <***> --------------------------------- */

func (s *legalEntityOrchestrator) AddBankAccountDetailsInLegalEntity(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
	req *legal_entity_contracts.AddBankAccountInLegalEntityRequestDto,
) error {
	if req == nil {
		return errors.New("request body should not be empty")
	}
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		return s.legalEntityService.AddNewBankAccounInLegalEntity(
			txCtx,
			orgID,
			emplID,
			legalEntityID,
			req,
		)
	})
}

/* ---------------------------------- <***> --------------------------------- */
func (s *legalEntityOrchestrator) AddIdentifiersInLegalEntity(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
	req *legal_entity_contracts.AddIdentifiersInLegalEntityRequestDto,
) error {
	if req == nil {
		return errors.New("request body should not be empty")
	}
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		return s.legalEntityService.AddNewIdentifiersInLegalEntity(
			txCtx,
			orgID,
			emplID,
			legalEntityID,
			req,
		)
	})
}

/* ---------------------------------- <***> --------------------------------- */

func (s *legalEntityOrchestrator) AddLicenseOrComplianceInLegalEntity(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
	req legal_entity_contracts.AddNewLicenseAndComplianceRequestDto,
) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		return s.legalEntityService.AddNewLicenseOrComplianceInLegalEntity(txCtx, orgID, emplID, legalEntityID, req)
	})
}

/* ---------------------------------- <***> --------------------------------- */

func (s *legalEntityOrchestrator) AddTaxRegistrationInLegalEntity(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
	req legal_entity_contracts.AddNewTaxRegistrationInLegalEntityRequestDto,
) error {
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		return s.legalEntityService.AddNewTaxRegistration(txCtx, orgID, emplID, legalEntityID, req)
	})
}

/* -------------------------------------------------------------------------- */
/*                       Legal Entity Accounting Service                      */
/* -------------------------------------------------------------------------- */

func (s *legalEntityOrchestrator) GetAccountingBooksOfLegalEntity(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) (accountingdtos.AccountingBookOfLegalEntity_Response_dto, error) {

	panic("")
}

func (s *legalEntityOrchestrator) GetAllAccountBalancesOfLegalEntity(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) ([]accountingdtos.AccountGroupBalanceResponseDto, error) {
	panic("")
}
