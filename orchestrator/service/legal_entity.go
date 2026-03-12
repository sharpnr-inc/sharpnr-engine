package orchestrator_service

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	accounting_engine_contracts "sharpnr-engine/service_engines/accounting_engine/contracts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (c *orchestratorService) CreateLegalEntity(ctx context.Context, orgID uuid.UUID, req *orchestrator_dtos.AddLegalEntity_REQ_DTO) error {

	return c.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		input := legal_entity_contracts.NewLegalEntityInput{
			OrgID:       orgID,
			LegalName:   req.Name,
			Type:        req.Type,
			AddressLine: req.Address_line,
			City:        req.City,
			State:       req.State,
			Country:     req.Country,
			Pincode:     req.Pincode,
		}

		legalEntityID, err := c.legalEntityService.CreateNewLegalEntity(txCtx, input)
		if err != nil {
			return err
		}

		accountInput := accounting_engine_contracts.DraftAccountingSetupInput{
			OrgID:               orgID,
			LegalEntityID:       legalEntityID,
			Name_of_legalEntity: input.LegalName,
		}
		if err := c.accountingEngineService.CreateDraftAccountingSetup(txCtx, &accountInput); err != nil {
			return err
		}

		return nil
	})
}

/* ---------------------------------- <***> --------------------------------- */

func (c *orchestratorService) GetAllLegalEntity(ctx context.Context, orgID uuid.UUID) ([]*orchestrator_dtos.LegalEntity_Response_dto, error) {
	les, err := c.legalEntityService.GetLegalEntityByOrgID(ctx, orgID)
	if err != nil {
		return nil, err
	}

	results := make([]*orchestrator_dtos.LegalEntity_Response_dto, len(les))

	for i, le := range les {

		results[i] = &orchestrator_dtos.LegalEntity_Response_dto{
			Name:        le.Name,
			ID:          le.LegalEntityID,
			Type:        le.EntityType,
			AddressLine: le.Country,
			GST:         le.GST,
			CIN_Number:  le.CIN,
			Country:     le.Country,
			State:       le.Country,
			City:        le.Country,
			Pincode:     le.Country,
			PAN_Number:  le.PAN,
		}
	}

	return results, nil

}

/* ---------------------------------- <***> --------------------------------- */

func (s *orchestratorService) AddTaxRegsitrationInLegalEntity(
	ctx context.Context,
	orgID, emplID, legalEntity uuid.UUID,
	req legal_entity_contracts.CreateLegalEntityTaxRegistrationRequestDto,
) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := s.legalEntityService.AddLegalEntityTaxRegistration(
			txCtx, orgID, emplID, legalEntity, req,
		); err != nil {
			return err
		}

		return nil
	})
}

/* ---------------------------------- <***> --------------------------------- */

func (s *orchestratorService) GetAllTaxRegistrationOfLegalEntity(
	ctx context.Context,
	orgID, emplID, legalEntity uuid.UUID,
) ([]legal_entity_contracts.LegalEntityTaxRegistrationResponseDto, error) {

	var res []legal_entity_contracts.LegalEntityTaxRegistrationResponseDto

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)
		taxRegistration, err := s.legalEntityService.GetAllLegelEntityTaxRegistration(txCtx, orgID, emplID, legalEntity)
		if err != nil {
			return err
		}

		res = taxRegistration

		return nil
	})

	if err != nil {
		return res, err
	}

	return res, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (c *orchestratorService) AddTaxCategoryInLegalEntiy(
	ctx context.Context,
	orgID, emplID, legalEntity uuid.UUID,
	req legal_entity_contracts.CreateLegalEntityTaxCategory,
) error {

	return c.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)
		if err := c.legalEntityService.AddLegalEntityTaxCategory(txCtx, orgID, emplID, legalEntity, req); err != nil {
			return err
		}
		return nil
	})
}

/* ---------------------------------- <***> --------------------------------- */

func (c *orchestratorService) ListAllTaxCategoryOfLegalEntity(
	ctx context.Context,
	orgID, emplID, legalEntity uuid.UUID,
) ([]legal_entity_contracts.LegalEntityTaxCategoryResponseDto, error) {

	var response []legal_entity_contracts.LegalEntityTaxCategoryResponseDto

	err := c.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		taxCategories, err := c.legalEntityService.GetAllLegalEntityTaxCategory(txCtx, orgID, emplID, legalEntity)
		if err != nil {
			return err
		}

		response = taxCategories
		return nil

	})

	if err != nil {
		return response, err
	}

	return response, nil

}

/* ---------------------------------- <***> --------------------------------- */
