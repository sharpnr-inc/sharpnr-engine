package orchestrator_service

import (
	"context"
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

	// in the request we have the hospitality and healthcare
	// healthcare is on-hold and it will empty or will completey missing

	// s.hospitalityDomainSvc.HotelService().SetupDefaultHotel(ctx, orgID, propertyID)
	// s.hospitalityDomainSvc.BarService().SetupDefaultBar()
	// s.hospitalityDomainSvc.RestaurantService().SetupDefaultRestaurant()

	return nil

}

/* ---------------------------------- <***> --------------------------------- */
