package orchestrator_service

import (
	"context"
	"errors"
	"fmt"
	dbtx "sharpnr-engine/core/dbTx"
	org_property_contracts "sharpnr-engine/core/org_property/contracts"
	hotel_contracts "sharpnr-engine/modules/hospitality/hotel/contracts"
	restaurant_contracts "sharpnr-engine/modules/hospitality/restaurant/contracts"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
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
) (*orchestrator_dtos.ModuleOperationResult, error) { // Changed signature

	result := &orchestrator_dtos.ModuleOperationResult{
		Created: []string{},
		Skipped: []string{},
	}

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// 1. Property Validation
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

		// 2. Identification Phase
		var modulesToCreate []string
		if req.Modules != nil {
			for _, module := range *req.Modules {
				exists := false
				switch module {
				case "HOTEL":
					h, _ := s.hospitalityDomainSvc.HotelService().GetHotelByPropertyID(ctx, propertyID)
					if h != nil && len(h.Hotels) > 0 {
						exists = true
					}
				case "BAR":
					b, _ := s.hospitalityDomainSvc.BarService().GetBarDetailsOfProperty(ctx, orgID, propertyID)
					if len(b) > 0 {
						exists = true
					}
				case "RESTAURANT":
					r, _ := s.hospitalityDomainSvc.RestaurantService().GetAllRestaurantOfProperty(ctx, propertyID)
					if len(r) > 0 {
						exists = true
					}
				}

				if exists {
					result.Skipped = append(result.Skipped, module)
				} else {
					modulesToCreate = append(modulesToCreate, module)
				}
			}
		}

		// 3. Constraint Check
		if len(modulesToCreate) > 1 {
			return fmt.Errorf("only one new module type is allowed at a time; requested: %v", modulesToCreate)
		}

		// 4. Execution Phase
		if len(modulesToCreate) == 1 {
			target := modulesToCreate[0]
			var setupErr error
			switch target {
			case "HOTEL":
				_, setupErr = s.hospitalityDomainSvc.HotelService().SetupDefaultHotel(txCtx, orgID, propertyID)
			case "BAR":
				_, setupErr = s.hospitalityDomainSvc.BarService().SetupDefaultBar(txCtx, orgID, propertyID)
			case "RESTAURANT":
				_, setupErr = s.hospitalityDomainSvc.RestaurantService().SetupDefaultRestaurant(txCtx, orgID, propertyID)
			}

			if setupErr != nil {
				return setupErr
			}
			result.Created = append(result.Created, target)
		}

		// 5. Infrastructure (Kitchen)
		if req.Infrastructure != nil {
			for _, infra := range *req.Infrastructure {
				if infra.Infra_Name == "KITCHEN" {
					if err := s.hospitalityDomainSvc.KitchenService().CreateDefaultKitchen(txCtx, orgID, propertyID, "DEFAULT"); err != nil {
						return err
					}
					result.Created = append(result.Created, "KITCHEN")
				}
			}
		}

		// 6. Build final message
		result.Message = fmt.Sprintf("%v created. %v already existed and were skipped.", result.Created, result.Skipped)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
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

func (s *orchestratorService) ConfigureRestaurant(
	ctx context.Context,
	orgID uuid.UUID,
	propertyID uuid.UUID,
	req *orchestrator_dtos.ConfigurePropertysRestaurant_Req_dto,
) error {
	// 1. Initialize the slice for the domain service contract
	// We use the length of the incoming layout to allocate memory efficiently
	generalDetails := &restaurant_contracts.UpdateGeneralDetailsInput{
		Name:         req.GeneralDetails.Name,
		IsConfigured: true,
		IsPureVeg:    req.GeneralDetails.IsPureVeg,
		OpeningTime:  req.GeneralDetails.OpeningTime,
		ClosingTime:  req.GeneralDetails.ClosingTime,
	}

	err := s.hospitalityDomainSvc.
		RestaurantService().
		UpdateRestaurantGeneralDetails(ctx, req.RestaurantID, generalDetails)
	if err != nil {
		return err
	}
	layoutInputs := make([]restaurant_contracts.SetupRestaurantLayoutInput, 0, len(req.RestaurantLayout))

	// 2. Map the DTO data to the Service Contract
	for _, f := range req.RestaurantLayout {
		layoutInputs = append(layoutInputs, restaurant_contracts.SetupRestaurantLayoutInput{
			Floor_Name:       f.FloorName, // Map from DTO field to Contract field
			Floor_Number:     f.Floor_Number,
			Number_of_tables: f.NumberOfTables,
		})
	}

	// 3. Pass the mapped slice to your domain service
	// Ensure you pass layoutInputs as the final argument
	if err := s.hospitalityDomainSvc.RestaurantService().SetupRestaurantLayout(
		ctx,
		propertyID,
		req.RestaurantID,
		layoutInputs,
	); err != nil {
		return err
	}

	// Future: Add logic here to configure Bar and Kitchen modules

	return nil
}

func (c *orchestratorService) ConfigureHotel(
	ctx context.Context,
	orgID uuid.UUID,
	propertyID uuid.UUID,
	req *orchestrator_dtos.ConfigureHotel_Request_dto,
) error {

	hotelLayouts := make([]hotel_contracts.FloorInput, 0, len(req.HotelLayout))

	for _, floorReq := range req.HotelLayout {
		hotelLayouts = append(hotelLayouts, hotel_contracts.FloorInput{
			HotelID:       req.HotelID,
			FloorName:     floorReq.FloorName,
			FloorNumber:   floorReq.FloorNumber,
			NumberOfRooms: floorReq.NumberOfRooms,
		})
	}

	if err := c.hospitalityDomainSvc.HotelService().SetUpLayout(ctx, hotelLayouts); err != nil {
		return err
	}

	genDetails := hotel_contracts.AddHotelGeneralInfoInput{
		HotelID:          req.HotelID,
		HotelName:        req.GenDetails.HotelDisplayName,
		HotelOpeningTime: req.GenDetails.OpeningTime,
		HotelClosingTime: req.GenDetails.ClosingTime,
		Is24X7:           req.GenDetails.Is24X7,
	}

	if err := c.hospitalityDomainSvc.HotelService().AddHotelGenDetails(ctx, &genDetails); err != nil {
		return err
	}

	return nil

}

func (s *orchestratorService) GetPropertyInfra(ctx context.Context, proertyID uuid.UUID, orgID uuid.UUID) (string, error){
	return "", nil
}
