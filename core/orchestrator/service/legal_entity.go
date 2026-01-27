package orchestrator_service

import (
	"context"
	orchestrator_dtos "sharpnr-engine/core/orchestrator/dtos"
	org_legal_entity_contracts "sharpnr-engine/core/org_legal_entity/contracts"

	"github.com/google/uuid"
)

// This service method creates the Legal entity
func (c *orchestratorService) CreateLegalEntity(ctx context.Context, orgID uuid.UUID, req *orchestrator_dtos.AddLegalEntity_REQ_DTO) error {
	// *This is orchestrator level Service
	// *So orchestrator Responsible for the Creating Legal Entity and its Accounting book
	// Each Legal Entity have the One Accounting book with fiscal Year
	// This is Not Simple as seems
	// So Legal Entity Owns the Accounting book
	// When Create the Legal Entity We also Create the Accounting book related to the Legal Entity
	// *Since everything is NOT Tightly connected we create the Accounting book via orchestrator

	// *Step 1 to call the LegalEntity Service for Creating the legal entity
	// For more details on the Legal entity creation Service --> go the service implementation

	//--> Here we have to map the request Legal Entity Service input contract
	input := org_legal_entity_contracts.NewLegalEntityInput{
		OrgID:       orgID,
		LegalName:   req.Name,
		Type:        req.Type,
		AddressLine: req.Address_line,
		City:        req.City,
		State:       req.State,
		Country:     req.Country,
		Pincode:     req.Pincode,
	}

	if err := c.legalEntityService.CreateNewLegalEntity(ctx, input); err != nil {
		return err
	}

	// ** This Creates the Legal Entity basic version.

	// Now we Get the Legal Entity By ORGID by legal Entity service we will get the legal entity
	// when we have the Legal Entity ID and Details We will Create the Accounting books

	//This section gonna creates the Accounting books

	// ** First map the legal entity reader to input of Accounting , needed for creating the Accounting
	// for more details go to methods implementations
	c.accountingEngineService.SetupAccountingForLegalEntity()

	// Step 3 then we add Accounting mapping

	return nil
}

func (c *orchestratorService) UpdateLegalEntity() {

}
