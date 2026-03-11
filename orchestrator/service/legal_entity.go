package orchestrator_service

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	legal_entity_contracts "sharpnr-engine/core/org_legal_entity/contracts"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	accounting_engine_contracts "sharpnr-engine/service_engines/accounting_engine/contracts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// This service method creates the Legal entity
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

func (c *orchestratorService) UpdateLegalEntity() {

}

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
