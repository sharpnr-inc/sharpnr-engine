package orchestrator_service

import (
	"context"
	"errors"
	"fmt"
	dbtx "sharpnr-engine/core/dbTx"
	orchestrator_dtos "sharpnr-engine/core/orchestrator/dtos"
	org_property_contracts "sharpnr-engine/core/org_property/contracts"
	"sharpnr-engine/shared/utils"
	"strings"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
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
	g, ctx := errgroup.WithContext(ctx)

	// Mutex protects the 'modules' struct from concurrent writes
	var mu sync.Mutex
	var foundData bool

	// Temporary container
	modules := &orchestrator_dtos.HospitalityPropertyModules{}

	// --- 1. Hotel Module ---
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

			mu.Lock()
			modules.HotelDetails = &hotelDTOs
			foundData = true
			mu.Unlock()
		}
		return nil
	})

	// --- 2. Future: Bar Module ---
	/*
	   g.Go(func() error {
	       // Replace with actual Bar service call later
	       bars, err := c.hospitalityDomainSvc.BarService().GetBarDetails(ctx, orgID, propertyID)
	       if err != nil && !utils.IsNotFoundError(err) { return err }

	       if bars != nil && len(bars.Items) > 0 {
	           // ... map to DTOs
	           mu.Lock()
	           // modules.BarDetails = &barDTOs
	           foundData = true
	           mu.Unlock()
	       }
	       return nil
	   })
	*/

	// --- 3. Future: Kitchen Module ---

	g.Go(func() error {
		// Replace with actual Kitchen service call later
		kitchens, err := c.hospitalityDomainSvc.KitchenService().GetKitchensByPropertyID(ctx, propertyID)
		if err != nil && !utils.IsNotFoundError(err) {
			return err
		}

		if len(kitchens) > 0 {
			kitchenDTOs := make([]orchestrator_dtos.KitchenDetails_dto, len(kitchens))
			for i, h := range kitchens {
				kitchenDTOs[i] = orchestrator_dtos.KitchenDetails_dto{
					ID:     h.ID,
					Name:   h.Name,
					Status: h.Status,
				}
			}

			mu.Lock()
			modules.KitchenDetails = &kitchenDTOs
			foundData = true
			mu.Unlock()
		}
		return nil
	})

	// Wait for all background tasks to finish
	if err := g.Wait(); err != nil {
		return err
	}

	// ONLY attach the modules object if at least one service found data
	if foundData {
		res.Modules = modules
	}

	return nil
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

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		prt, err := s.orgPropertyService.GetPropertiesOfOrg(txCtx, orgID)
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

		for _, module := range req.Modules {
			switch module.Name {
			case "HOTEL":
				_, err := s.hospitalityDomainSvc.
					HotelService().
					SetupDefaultHotel(txCtx, orgID, propertyID)
				if err != nil {
					return err
				}

				if module.KitchenEnabled {
					if err := s.hospitalityDomainSvc.
						KitchenService().
						CreateDefaultKitchen(txCtx, orgID, propertyID, module.Name); err != nil {
						return err
					}
				}

			case "BAR":
				_, err := s.hospitalityDomainSvc.
					BarService().
					SetupDefaultBar(txCtx, orgID, propertyID)
				if err != nil {
					return err
				}

				if module.KitchenEnabled {
					if err := s.hospitalityDomainSvc.
						KitchenService().
						CreateDefaultKitchen(txCtx, orgID, propertyID, module.Name); err != nil {
						return err
					}
				}

			case "RESTAURANT":
				_, err := s.hospitalityDomainSvc.
					RestaurantService().
					SetupDefaultRestaurant(txCtx, orgID, propertyID)
				if err != nil {
					return err
				}

				if module.KitchenEnabled {
					if err := s.hospitalityDomainSvc.
						KitchenService().
						CreateDefaultKitchen(txCtx, orgID, propertyID, module.Name); err != nil {
						return err
					}
				}

			default:
				return fmt.Errorf("unsupported module: %s", module.Name)

			}

		}

		return nil
	})

}
