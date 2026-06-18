package orchestrator_service

import (
	"context"
	"fmt"
	hospitality_utilities "sharpnr-engine/business_ops/hospitality/utilities"
	org_property_contracts "sharpnr-engine/core/org_property/contracts"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/* ---------------------------------- <***> --------------------------------- */

func (s *orchestratorService) GetAllPropertiesOfOrg(ctx context.Context, orgID, emplID uuid.UUID) (*orchestrator_dtos.PropertiesOfOrganization_ResDto, error) {

	// make sure to verify the organization status and its exists
	// emplID can be checked user have the access for the fetching the data

	res, err := s.orgPropertyService.GetPropertiesOfOrg(ctx, orgID)
	if err != nil {
		return nil, err
	}

	readers := make([]orchestrator_dtos.PropertiesReaderSlice, 0, len(res))

	for _, p := range res {
		readers = append(readers, orchestrator_dtos.PropertiesReaderSlice{
			ID:          p.ID,
			Name:        p.Name,
			City:        p.City,
			State:       p.State,
			Code:        p.Code,
			AddressLine: p.AddressLine,
			Country:     p.Country,
		})
	}

	return &orchestrator_dtos.PropertiesOfOrganization_ResDto{
		Properties: readers,
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (c *orchestratorService) GetPropertyDetail(ctx context.Context, orgID, propertyID uuid.UUID) (*orchestrator_dtos.PropertyDetails_res_dto, error) {

	propDetail, err := c.orgPropertyService.GetPropertyDetailsByID(ctx, orgID, propertyID)
	if err != nil {
		return nil, err
	}

	// Map the base details to your response DTO
	response := &orchestrator_dtos.PropertyDetails_res_dto{
		Name:        propDetail.Name,
		Code:        propDetail.Code,
		City:        propDetail.City,
		AddressLine: propDetail.AddressLine,
		State:       propDetail.State,
		Country:     propDetail.Country,
	}

	return response, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *orchestratorService) CreateProperty(ctx context.Context, orgID uuid.UUID, req *orchestrator_dtos.CreatePropertyReq_dto) error {
	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		prInp := org_property_contracts.CreateOrgPropertyInput{
			OrgID:       orgID,
			Name:        req.Name,
			Code:        req.Code,
			AddressLine: req.AddressLine,
			City:        req.City,
			State:       req.State,
			Country:     req.Country,
			Pincode:     req.Pincode,
		}

		if err := s.orgPropertyService.CreatePropertyDefault(txCtx, prInp); err != nil {
			return err
		}

		status, err := s.organizationService.GetStatus(txCtx, orgID)
		if err != nil {
			return err
		}

		if status == "NOT_CONFIGURED" {
			if err := s.organizationService.UpdateStatus(txCtx, orgID, "CONFIGURED"); err != nil {
				return err
			}
		}

		return nil

	})

}

/* ---------------------------------- <***> --------------------------------- */

func (s *orchestratorService) AddBusinessOperationModulesInProperty(
	ctx context.Context,
	orgID, emplID, propertyID uuid.UUID,
	req orchestrator_dtos.AddBusinessOpModulesRequestDto,
) error {

	// Request validation is done in the controller, so the module types and
	// quantities here are already known-valid.

	// Healthcare is on hold for now — req.HealthcareModules is intentionally
	// ignored until that domain is wired up.

	for _, m := range req.HospitalityModules {
		for i := 0; i < m.Quantity; i++ {
			if err := s.setupHospitalityModule(ctx, orgID, propertyID, m.Module); err != nil {
				return fmt.Errorf(
					"setup hospitality module %q (instance %d of %d) for property %s: %w",
					m.Module, i+1, m.Quantity, propertyID, err,
				)
			}
		}
	}

	return nil
}

// setupHospitalityModule provisions a single default instance of the given
// hospitality module for the property.
func (s *orchestratorService) setupHospitalityModule(
	ctx context.Context,
	orgID, propertyID uuid.UUID,
	module hospitality_utilities.HospitalityModulesType,
) error {
	switch module {
	case hospitality_utilities.Hotel:
		_, err := s.hospitalityDomainSvc.HotelService().SetupDefaultHotel(ctx, orgID, propertyID)
		return err

	case hospitality_utilities.Bar:
		_, err := s.hospitalityDomainSvc.BarService().SetupDefaultBar(ctx, orgID, propertyID)
		return err

	case hospitality_utilities.Restaurant:
		_, err := s.hospitalityDomainSvc.RestaurantService().SetupDefaultRestaurant(ctx, orgID, propertyID)
		return err

	case hospitality_utilities.NightClub,
		hospitality_utilities.Resort,
		hospitality_utilities.CloudKitchen_Takeaways:
		// Declared module types that don't have a setup service yet.
		return fmt.Errorf("hospitality module %q is not yet supported", module)

	default:
		return fmt.Errorf("unknown hospitality module %q", module)
	}
}

/* ---------------------------------- <***> --------------------------------- */

func (s *orchestratorService) AddInfraModulesInProperty(
	ctx context.Context,
	orgID, emplID, propertyID uuid.UUID,
	request orchestrator_dtos.AddInfraModulesRequestDto,
) error {

	// Scope for now: only Hospitality infra is provisioned, and within it only
	// the Kitchen module. General infra and Healthcare infra are intentionally
	// ignored even if present in the request body, and any other hospitality
	// infra types are skipped until their setup services exist.

	for _, m := range request.HospitalityInfra {
		if m.Module != hospitality_utilities.Kitchen {
			// Not yet supported — skip rather than fail the whole request.
			continue
		}

		for i := 0; i < m.Quantity; i++ {
			err := s.hospitalityDomainSvc.KitchenService().
				CreateDefaultKitchen(ctx, orgID, propertyID, string(m.Module))
			if err != nil {
				return fmt.Errorf(
					"create default kitchen (instance %d of %d) for property %s: %w",
					i+1, m.Quantity, propertyID, err,
				)
			}
		}
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *orchestratorService) ListAllBusinessOpsModules(ctx context.Context, orgID, emplID, propertyID uuid.UUID) error {
	return nil
}

func (s *orchestratorService) ListAllInfraModulesOfProperty(ctx context.Context, orgID, emplID, propertyID uuid.UUID) error {
	return nil
}
