package orchestrator_service

import (
	"context"
	"errors"
	"fmt"
	orchestrator_dtos "sharpnr-engine/core/orchestrator/dtos"
	org_property_contracts "sharpnr-engine/core/org_property/contracts"
	"sharpnr-engine/shared/utils"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

func (s *orchestratorService) GetAllPropertiesOfOrg(ctx context.Context, orgID uuid.UUID) (*orchestrator_dtos.PropertiesOfOrganization_ResDto, error) {

	res, err := s.orgPropertyService.GetPropertiesOfOrg(ctx, orgID)
	if err != nil {
		return nil, err
	}

	readers := make([]orchestrator_dtos.PropertiesReaderSlice, 0, len(res.Properties))

	for _, p := range res.Properties {
		readers = append(readers, orchestrator_dtos.PropertiesReaderSlice{
			ID:     p.ID,
			Name:   p.Name,
			City:   p.City,
			State:  p.State,
			Domain: p.Domain,
			Code:   p.Code,
		})
	}

	return &orchestrator_dtos.PropertiesOfOrganization_ResDto{
		Properties: readers,
	}, nil
}

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
		Domain:      &propDetail.Domain,
	}

	// Case 1: Empty Domain - return the basic details immediately
	if propDetail.Domain == "" {
		return response, nil
	}

	// Case 2: Hospitality Domain - enrich with sub-modules
	if propDetail.Domain == "HOSPITALITY" {
		err := c.enrichHospitalityData(ctx, orgID, propertyID, response)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func (c *orchestratorService) enrichHospitalityData(
	ctx context.Context,
	orgID, propertyID uuid.UUID,
	res *orchestrator_dtos.PropertyDetails_res_dto,
) error {
	// 1. Initialize the Modules container
	res.Modules = &orchestrator_dtos.HospitalityPropertyModules{}

	g, ctx := errgroup.WithContext(ctx)

	// --- Hotel Module ---
	g.Go(func() error {
		hotel, err := c.hospitalityDomainSvc.HotelService().GetHotelDetails(ctx, orgID, propertyID)
		if err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return err
		}

		if hotel != nil && len(hotel.Hotels) > 0 {
			hotelDTOs := make([]orchestrator_dtos.HotelDetails_dto, len(hotel.Hotels))

			for i, h := range hotel.Hotels {
				hotelDTOs[i] = orchestrator_dtos.HotelDetails_dto{
					ID:     h.ID,
					Name:   h.Name,
					Status: h.Status,
				}
			}

			// Since the struct expects *[]HotelDetails_dto, we take the address of the slice
			res.Modules.HotelDetails = &hotelDTOs
		}

		return nil
	})

	// --- Future: Add Bar and Restaurant g.Go blocks here ---

	// 2. Wait for all background tasks to finish
	return g.Wait()
}

func (s *orchestratorService) ConfigurePropertyToDomain(
	ctx context.Context,
	orgID uuid.UUID,
	propertyID uuid.UUID,
	req *orchestrator_dtos.AddDomainToPropertyReq_DTO,
) error {

	verify, err := s.organizationService.VerifyOrgByID(ctx, orgID)
	if err != nil {
		return err
	}

	if !verify {
		return errors.New("org invalid")
	}

	return s.orgPropertyService.AddDomain(ctx, orgID, propertyID, req.Domain)
}

//=---------------------------------------------------------------//

func (s *orchestratorService) AddModulesInProperty(
	ctx context.Context,
	orgID uuid.UUID,
	propertyID uuid.UUID,
	req *orchestrator_dtos.PropertyModuleConfigurationDTO,
) error {

	// *Step One we have to vErify if user have the domain configured or not
	// if they have then they must have respective modules
	// for example if they have the Hospitality then the modules must be
	// "HOTEL","BAR","RESTAURANT"
	// then we create the Switch case if Restaurant then we create Restaurant ...

	prt, err := s.orgPropertyService.GetPropertiesOfOrg(ctx, orgID)
	if err != nil {
		return err
	}

	var property *org_property_contracts.PropertiesReader

	for i := range prt.Properties {
		if prt.Properties[i].ID == propertyID {
			property = &prt.Properties[i]
			break
		}
	}

	if property == nil {
		return errors.New("property not found")
	}

	if strings.TrimSpace(property.Domain) == "" {
		return errors.New("domain not configured for property")
	}

	// Now we will add the modules that we receive in the request

	for _, module := range req.Modules {
		switch module.Name {
		case "HOTEL":
			_, err := s.hospitalityDomainSvc.
				HotelService().
				SetupDefaultHotel(ctx, orgID, propertyID)
			if err != nil {
				return err
			}

			if module.KitchenEnabled {
				if err := s.hospitalityDomainSvc.
					KitchenService().
					CreateDefaultKitchen(ctx, orgID, propertyID, module.Name); err != nil {
					return err
				}
			}

		case "BAR":
			_, err := s.hospitalityDomainSvc.
				BarService().
				SetupDefaultBar(ctx, orgID, propertyID)
			if err != nil {
				return err
			}

			if module.KitchenEnabled {
				if err := s.hospitalityDomainSvc.
					KitchenService().
					CreateDefaultKitchen(ctx, orgID, propertyID, module.Name); err != nil {
					return err
				}
			}

		case "RESTAURANT":
			_, err := s.hospitalityDomainSvc.
				RestaurantService().
				SetupDefaultRestaurant(ctx, orgID, propertyID)
			if err != nil {
				return err
			}

			if module.KitchenEnabled {
				if err := s.hospitalityDomainSvc.
					KitchenService().
					CreateDefaultKitchen(ctx, orgID, propertyID, module.Name); err != nil {
					return err
				}
			}

		default:
			return fmt.Errorf("unsupported module: %s", module.Name)

		}

	}

	return nil
}
