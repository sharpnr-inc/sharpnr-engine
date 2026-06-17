package orchestrator_service

import (
	"context"
	org_property_contracts "sharpnr-engine/core/org_property/contracts"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/shared/utils"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
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

	err = c.enrichHospitalityData(ctx, orgID, propertyID, response)
	if err != nil {
		return nil, err
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
