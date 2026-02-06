package orchestrator_service

import (
	"context"
	"errors"
	"fmt"
	dbtx "sharpnr-engine/core/dbTx"
	orchestrator_dtos "sharpnr-engine/core/orchestrator/dtos"
	org_property_contracts "sharpnr-engine/core/org_property/contracts"
	"sharpnr-engine/shared/utils"
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
			ID:           p.ID,
			Name:         p.Name,
			City:         p.City,
			State:        p.State,
			Domain:       p.Domain,
			Code:         p.Code,
			IsConfigured: p.IsConfigured,
		})
	}

	return &orchestrator_dtos.PropertiesOfOrganization_ResDto{
		Properties: readers,
	}, nil
}

//=---------------------------------------------------------------//

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

	var (
		mu      sync.Mutex
		modules []orchestrator_dtos.ModulesDetails_dto
		infra   []orchestrator_dtos.PropertyInfra_res_dto
	)

	// --- 1. Hotel Module ---
	g.Go(func() error {
		hotel, err := c.hospitalityDomainSvc.
			HotelService().
			GetHotelDetails(ctx, orgID, propertyID)

		if err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return err
		}

		if hotel == nil || len(hotel.Hotels) == 0 {
			return nil
		}

		hotelDTOs := make([]orchestrator_dtos.ModulesDetails_dto, 0, len(hotel.Hotels))
		for _, h := range hotel.Hotels {
			hotelDTOs = append(hotelDTOs, orchestrator_dtos.ModulesDetails_dto{
				ID:           h.ID,
				Name:         h.Name,
				Status:       h.Status,
				IsConfigured: h.IsConfigured,
				Module_Type:  "HOTEL",
			})
		}

		mu.Lock()
		modules = append(modules, hotelDTOs...)
		mu.Unlock()

		return nil
	})

	// --- 2. Restaurant Module ---
	g.Go(func() error {
		restaurant, err := c.hospitalityDomainSvc.
			RestaurantService().
			GetAllRestaurantOfProperty(ctx, propertyID)

		if err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return err
		}

		if len(restaurant) == 0 {
			return nil
		}

		dTOs := make([]orchestrator_dtos.ModulesDetails_dto, 0, len(restaurant))
		for _, h := range restaurant {
			dTOs = append(dTOs, orchestrator_dtos.ModulesDetails_dto{
				ID:           h.ID,
				Name:         h.Name,
				Status:       h.Status,
				IsConfigured: h.IsConfigured,
				Module_Type:  "RESTAURANT",
			})
		}

		mu.Lock()
		modules = append(modules, dTOs...)
		mu.Unlock()

		return nil
	})

	// --- 3. Kitchen (Infra) ---
	g.Go(func() error {
		kitchens, err := c.hospitalityDomainSvc.
			KitchenService().
			GetKitchensByPropertyID(ctx, propertyID)

		if err != nil && !utils.IsNotFoundError(err) {
			return err
		}

		if len(kitchens) == 0 {
			return nil
		}

		kitchenDTOs := make([]orchestrator_dtos.PropertyInfra_res_dto, 0, len(kitchens))
		for _, k := range kitchens {
			kitchenDTOs = append(kitchenDTOs, orchestrator_dtos.PropertyInfra_res_dto{
				ID:         k.ID,
				Name:       k.Name,
				Status:     k.Status,
				Infra_Type: "KITCHEN",
			})
		}

		mu.Lock()
		infra = append(infra, kitchenDTOs...)
		mu.Unlock()

		return nil
	})

	// Wait for all goroutines
	if err := g.Wait(); err != nil {
		return err
	}

	// Attach only if data exists
	if len(modules) > 0 {
		res.Modules = &modules
	}
	if len(infra) > 0 {
		res.Infra = &infra
	}

	return nil
}

// =---------------------------------------------------------------//

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

		if req.Modules != nil {
			for _, module := range *req.Modules {
				switch module {
				case "HOTEL":
					_, err := s.hospitalityDomainSvc.
						HotelService().
						SetupDefaultHotel(txCtx, orgID, propertyID)
					if err != nil {
						return err
					}

				case "BAR":
					_, err := s.hospitalityDomainSvc.
						BarService().
						SetupDefaultBar(txCtx, orgID, propertyID)
					if err != nil {
						return err
					}

				case "RESTAURANT":
					_, err := s.hospitalityDomainSvc.
						RestaurantService().
						SetupDefaultRestaurant(txCtx, orgID, propertyID)
					if err != nil {
						return err
					}

				default:
					return fmt.Errorf("unsupported module: %s", module)

				}

			}
		}

		if req.Infrastructure != nil {
			for _, infra := range *req.Infrastructure {

				switch infra.Infra_Name {
				case "KITCHEN":
					if err := s.hospitalityDomainSvc.
						KitchenService().
						CreateDefaultKitchen(txCtx, orgID, propertyID, "DEFAULT"); err != nil {
						return err
					}
				default:
					return fmt.Errorf("unsupported module: %s", infra.Infra_Name)
				}
			}
		}

		return nil
	})

}

// =---------------------------------------------------------------//
func (s *orchestratorService) CreateProperty(ctx context.Context, orgID uuid.UUID, req *orchestrator_dtos.CreatePropertyReq_dto) error {
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

	if err := s.orgPropertyService.CreatePropertyDefault(ctx, prInp); err != nil {
		return err
	}

	return nil
}

//=---------------------------------------------------------------//
