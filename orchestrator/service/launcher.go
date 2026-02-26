package orchestrator_service

import (
	"context"
	property_constants "sharpnr-engine/core/org_property/constants"
	org_property_contracts "sharpnr-engine/core/org_property/contracts"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	"sharpnr-engine/shared/utils"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

func (c *orchestratorService) GetSharpnrLauncherData(
	ctx context.Context,
	orgID uuid.UUID,
	emplID uuid.UUID,
) (orchestrator_dtos.OrgLauncherResponse, error) {

	res := orchestrator_dtos.OrgLauncherResponse{
		Domains: []map[string]orchestrator_dtos.DomainConfig{},
	}

	properties, err := c.orgPropertyService.GetPropertiesOfOrg(ctx, orgID)
	if err != nil {
		return res, err
	}

	if len(properties.Properties) == 0 {
		return res, nil
	}

	// check hospitality
	if hasDomain(
		properties.Properties,
		property_constants.Hospitality,
	) {

		hospitalityConfig, err := c.buildHospitalityLauncher(
			ctx,
			orgID,
			properties.Properties,
		)
		if err != nil {
			return res, err
		}

		// only add if something exists
		if len(hospitalityConfig.Modules) > 0 ||
			len(hospitalityConfig.Infras) > 0 {

			res.Domains = append(res.Domains,
				map[string]orchestrator_dtos.DomainConfig{
					"Hospitality": hospitalityConfig,
				},
			)
		}
	}

	return res, nil
}

func hasDomain(
	properties []org_property_contracts.PropertiesReader,
	domain property_constants.PropertyDomainTypes,
) bool {
	for _, p := range properties {
		if p.Domain == domain {
			return true
		}
	}
	return false
}

func (c *orchestratorService) buildHospitalityLauncher(
	ctx context.Context,
	orgID uuid.UUID,
	properties []org_property_contracts.PropertiesReader,
) (orchestrator_dtos.DomainConfig, error) {

	g, ctx := errgroup.WithContext(ctx)

	moduleSet := sync.Map{}
	infraSet := sync.Map{}

	for _, p := range properties {

		// only hospitality properties
		if p.Domain != property_constants.Hospitality {
			continue
		}

		propertyID := p.ID

		/* ---------- HOTEL ---------- */
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

			if hotel != nil && len(hotel.Hotels) > 0 {
				moduleSet.Store("HOTEL", true)
			}
			return nil
		})

		/* ---------- RESTAURANT ---------- */
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

			if len(restaurant) > 0 {
				moduleSet.Store("RESTAURANT", true)
			}
			return nil
		})

		// g.Go(func() error {
		// 	restaurant, err := c.hospitalityDomainSvc.
		// 		BarService().
		// 		GetAllRestaurantOfProperty(ctx, propertyID)

		// 	if err != nil {
		// 		if utils.IsNotFoundError(err) {
		// 			return nil
		// 		}
		// 		return err
		// 	}

		// 	if len(restaurant) > 0 {
		// 		moduleSet.Store("RESTAURANT", true)
		// 	}
		// 	return nil
		// })

		/* ---------- KITCHEN ---------- */
		g.Go(func() error {
			kitchens, err := c.hospitalityDomainSvc.
				KitchenService().
				GetKitchensByPropertyID(ctx, propertyID)

			if err != nil && !utils.IsNotFoundError(err) {
				return err
			}

			if len(kitchens) > 0 {
				infraSet.Store("KITCHEN", true)
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return orchestrator_dtos.DomainConfig{}, err
	}

	/* ---------- convert sets → slices ---------- */

	var modules []string
	moduleSet.Range(func(k, _ any) bool {
		modules = append(modules, k.(string))
		return true
	})

	var infras []string
	infraSet.Range(func(k, _ any) bool {
		infras = append(infras, k.(string))
		return true
	})

	return orchestrator_dtos.DomainConfig{
		Modules: modules,
		Infras:  infras,
	}, nil
}
