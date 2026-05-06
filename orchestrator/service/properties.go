package orchestrator_service

import (
	"context"
	"errors"
	"fmt"
	dbtx "sharpnr-engine/core/dbTx"
	org_property_contracts "sharpnr-engine/core/org_property/contracts"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	"sharpnr-engine/shared/utils"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

/* ---------------------------------- <***> --------------------------------- */

// GetAllPropertiesOfOrg retrieves all properties associated with a given organization ID. It returns a structured response containing the list of properties or an error if the retrieval fails.
// ? Point to improve, :--> Need to add the emplID, sharpnr_sentinel can be check is employee is authorised to see this or not, currently any employee of the org can see all the properties, in future we can have
// ? some restriction based on role or permission

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
) (*orchestrator_dtos.ModuleOperationResult, error) {

	result := &orchestrator_dtos.ModuleOperationResult{
		Created: []string{},
		Skipped: []string{},
	}

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// ----------------------------------------------------
		// 1. Property Validation
		// ----------------------------------------------------
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

		// ----------------------------------------------------
		// 2. Identification Phase
		// ----------------------------------------------------
		var modulesToCreate []string

		if req.Modules != nil {
			for _, module := range *req.Modules {

				exists := false

				switch module {

				case "HOTEL":
					h, _ := s.hospitalityDomainSvc.
						HotelService().
						GetHotelByPropertyID(txCtx, propertyID)

					if h != nil && len(h.Hotels) > 0 {
						exists = true
					}

				case "BAR":
					b, _ := s.hospitalityDomainSvc.
						BarService().
						GetBarDetailsOfProperty(txCtx, orgID, propertyID)

					if len(b) > 0 {
						exists = true
					}

				case "RESTAURANT":
					r, _ := s.hospitalityDomainSvc.
						RestaurantService().
						GetAllRestaurantOfProperty(txCtx, propertyID)

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

		// ----------------------------------------------------
		// 3. Execution Phase (MULTI MODULE SUPPORT)
		// ----------------------------------------------------
		for _, target := range modulesToCreate {

			var setupErr error

			switch target {

			case "HOTEL":
				_, setupErr = s.hospitalityDomainSvc.
					HotelService().
					SetupDefaultHotel(txCtx, orgID, propertyID)

			case "BAR":
				_, setupErr = s.hospitalityDomainSvc.
					BarService().
					SetupDefaultBar(txCtx, orgID, propertyID)

			case "RESTAURANT":
				_, setupErr = s.hospitalityDomainSvc.
					RestaurantService().
					SetupDefaultRestaurant(txCtx, orgID, propertyID)
			}

			if setupErr != nil {
				return setupErr
			}

			result.Created = append(result.Created, target)
		}

		// ----------------------------------------------------
		// 4. Infrastructure (Kitchen qty support)
		// ----------------------------------------------------
		if req.Infrastructure != nil {

			for _, infra := range *req.Infrastructure {

				if infra.Infra_Name == "Kitchen/Service Point" {

					qty := infra.InfraQty
					if qty <= 0 {
						qty = 1
					}

					for i := 0; i < qty; i++ {

						name := fmt.Sprintf("DEFAULT-%d", i+1)

						if err := s.hospitalityDomainSvc.
							KitchenService().
							CreateDefaultKitchen(
								txCtx,
								orgID,
								propertyID,
								name,
							); err != nil {
							return err
						}

						result.Created = append(
							result.Created,
							"Kitchen/Service Point",
						)
					}
				}
			}
		}

		// ----------------------------------------------------
		// 5. Update Property Configuration Status
		// ----------------------------------------------------
		if err := s.orgPropertyService.
			UpdatePropertyConfigurationStatus(
				txCtx,
				propertyID,
				true,
			); err != nil {
			return err
		}

		// ----------------------------------------------------
		// 6. Final Message
		// ----------------------------------------------------
		result.Message = fmt.Sprintf(
			"%v created. %v already existed and were skipped.",
			result.Created,
			result.Skipped,
		)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *orchestratorService) AddInfraInProperty(
	ctx context.Context,
	orgID uuid.UUID,
	propertyID uuid.UUID,
	req *orchestrator_dtos.PropertyInfraConfiguration,
) error {

	// Guard clause
	if req == nil {
		return fmt.Errorf("infra configuration is nil")
	}

	// Handle Kitchen / Service Point infra
	if req.Infra_Name != "Kitchen/Service Point" {
		return nil
	}

	// Default quantity fallback
	qty := req.InfraQty
	if qty <= 0 {
		qty = 1
	}

	kitchenSvc := s.hospitalityDomainSvc.KitchenService()

	for i := 0; i < qty; i++ {
		name := fmt.Sprintf("DEFAULT-%d", i+1)

		if err := kitchenSvc.CreateDefaultKitchen(
			ctx,
			orgID,
			propertyID,
			name,
		); err != nil {
			return fmt.Errorf("failed creating kitchen %s: %w", name, err)
		}
	}

	return nil
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

func (s *orchestratorService) GetPropertyInfra(
	ctx context.Context,
	propertyID uuid.UUID,
	orgID uuid.UUID, infraType *string) (*orchestrator_dtos.ListAllInfraOfProperty_Res_dto, error) {

	_, err := s.orgPropertyService.GetPropertyDetailsByID(ctx, orgID, propertyID)
	if err != nil {
		return nil, err
	}

	// For Now we have only Kitchen as Infra, so directly calling that. In future, we can use infraType to decide which infra details to fetch
	kitchens, err := s.hospitalityDomainSvc.KitchenService().GetKitchensByPropertyID(ctx, propertyID)
	if err != nil {
		return nil, err
	}

	results := make([]*orchestrator_dtos.PropertyInfra_res_dto, 0, len(kitchens))

	for i, inf := range kitchens {
		results[i] = &orchestrator_dtos.PropertyInfra_res_dto{
			ID:         inf.ID,
			Name:       inf.Name,
			Status:     inf.Status,
			Infra_Type: inf.Name,
		}
	}

	return &orchestrator_dtos.ListAllInfraOfProperty_Res_dto{
		PropertyID:      propertyID,
		Infrastructures: results,
	}, nil

}

/* ---------------------------------- <***> --------------------------------- */
