package orchestrator_service

import (
	"context"
	"fmt"

	bar_contracts "sharpnr-engine/business_ops/hospitality/bar/contracts"
	hotel_contracts "sharpnr-engine/business_ops/hospitality/hotel/contracts"
	kitchen_contracts "sharpnr-engine/business_ops/hospitality/kitchen/contracts"
	restaurant_contracts "sharpnr-engine/business_ops/hospitality/restaurant/contracts"
	hospitality_utilities "sharpnr-engine/business_ops/hospitality/utilities"
	org_property_contracts "sharpnr-engine/core/org_property/contracts"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// resourceTypeProperty is the Cedar entity type for a property resource.
const resourceTypeProperty = "Property"

/* ---------------------------------- <***> --------------------------------- */

func (s *orchestratorService) GetAllPropertiesOfOrg(ctx context.Context, orgID, emplID uuid.UUID) (*orchestrator_dtos.PropertiesOfOrganization_ResDto, error) {

	// Authorization: can this employee view properties in this org? This is an
	// org-wide listing (no single property instance), so the org id doubles as
	// the resource handle; the engine stamps orgId on both principal and
	// resource, so the tenant guard in the policy holds.
	// decision, err := s.sentinelEngine.IsAuthorized(ctx, sentinelengine.AuthorizeInput{
	// 	OrgID:  orgID,
	// 	UserID: emplID,
	// 	Action: string(sentinel_actions.ViewProperty),
	// 	Resource: sentinelengine.Resource{
	// 		Type: resourceTypeProperty,
	// 		ID:   orgID.String(),
	// 	},
	// })
	// if err != nil {
	// 	return nil, fmt.Errorf("authorize view properties: %w", err)
	// }
	// if !decision.Allowed {
	// 	return nil, fmt.Errorf("forbidden: not allowed to view properties")
	// }

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

func (s *orchestratorService) CreateProperty(ctx context.Context, orgID, emplID uuid.UUID, req *orchestrator_dtos.CreatePropertyReq_dto) error {
	// Authorize BEFORE opening a transaction — there's no point starting a write
	// the caller may not be allowed to make. The property doesn't exist yet, so
	// the org id is the resource handle; the engine stamps orgId on both the
	// principal and the resource, satisfying the tenant guard in the owner policy
	// (`resource.orgId == principal.orgId`).
	// decision, err := s.sentinelEngine.IsAuthorized(ctx, sentinelengine.AuthorizeInput{
	// 	OrgID:  uuid.New(),
	// 	UserID: emplID,
	// 	Action: string(sentinel_actions.CreateProperty),
	// 	Resource: sentinelengine.Resource{
	// 		Type: resourceTypeProperty,
	// 		ID:   orgID.String(),
	// 	},
	// })
	// if err != nil {
	// 	// Infrastructure error (couldn't load policies/membership). Fail closed.
	// 	slog.Error("sentinel: authorization check errored",
	// 		"action", sentinel_actions.CreateProperty,
	// 		"org_id", orgID,
	// 		"user_id", emplID,
	// 		"error", err.Error(),
	// 	)
	// 	return fmt.Errorf("authorize create property: %w", err)
	// }

	// // Log the full decision so the engine can be seen working end to end:
	// // allowed/denied, which policy ids drove it (Reasons), and any per-policy
	// // evaluation errors (Errors) that did not change the authoritative outcome.
	// slog.Info("sentinel: authorization decision",
	// 	"action", sentinel_actions.CreateProperty,
	// 	"org_id", uuid.New(),
	// 	"user_id", emplID,
	// 	"resource_type", resourceTypeProperty,
	// 	"resource_id", orgID.String(),
	// 	"allowed", decision.Allowed,
	// 	"reasons", decision.Reasons,
	// 	"eval_errors", decision.Errors,
	// )

	// if !decision.Allowed {
	// 	slog.Warn("sentinel: create property denied",
	// 		"action", sentinel_actions.CreateProperty,
	// 		"org_id", orgID,
	// 		"user_id", emplID,
	// 	)
	// 	return fmt.Errorf("forbidden: not allowed to create a property in this organization")
	// }

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

func (s *orchestratorService) ListAllBusinessOpsModules(
	ctx context.Context,
	orgID, emplID, propertyID uuid.UUID,
) (orchestrator_dtos.PropertyBusinessOperationModulesResponseDto, error) {

	// Initialised to empty (non-nil) slices so the JSON response is [] rather
	// than null when a domain has no modules.
	resp := orchestrator_dtos.PropertyBusinessOperationModulesResponseDto{
		HospitalityModules: make([]orchestrator_dtos.HospitalityBusinessModule, 0),
		HealthCareModules:  make([]orchestrator_dtos.HealthCareBusinessModule, 0),
	}

	// ---------------------------- Hospitality ----------------------------

	bars, err := s.hospitalityDomainSvc.BarService().
		GetBarDetailsOfProperty(ctx, orgID, propertyID)
	if err != nil {
		return orchestrator_dtos.PropertyBusinessOperationModulesResponseDto{},
			fmt.Errorf("list bars for property %s: %w", propertyID, err)
	}
	for _, b := range bars {
		if b == nil {
			continue
		}
		resp.HospitalityModules = append(resp.HospitalityModules, mapBarToBusinessModule(b))
	}

	// Single hotel for now; comment in source says this becomes a slice later.
	hotels, err := s.hospitalityDomainSvc.HotelService().
		GetHotelByPropertyID(ctx, propertyID)
	if err != nil {
		return orchestrator_dtos.PropertyBusinessOperationModulesResponseDto{},
			fmt.Errorf("list hotels for property %s: %w", propertyID, err)
	}
	for _, h := range hotels {
		resp.HospitalityModules = append(resp.HospitalityModules, mapHotelToBusinessModule(h))
	}

	restaurants, err := s.hospitalityDomainSvc.RestaurantService().
		GetAllRestaurantOfProperty(ctx, propertyID)
	if err != nil {
		return orchestrator_dtos.PropertyBusinessOperationModulesResponseDto{},
			fmt.Errorf("list restaurants for property %s: %w", propertyID, err)
	}
	for _, r := range restaurants {
		if r == nil {
			continue
		}
		resp.HospitalityModules = append(resp.HospitalityModules, mapRestaurantToBusinessModule(r))
	}

	// ----------------------------- Healthcare -----------------------------
	// Domain not ready yet — HealthCareModules intentionally left empty.

	return resp, nil
}

// --- mapping helpers ---
// ModuleType is set from the originating service (known & correct).
// Field accessors below are ASSUMPTIONS — adjust once the contract structs are known.

func mapBarToBusinessModule(b *bar_contracts.BarDetailsReader) orchestrator_dtos.HospitalityBusinessModule {
	return orchestrator_dtos.HospitalityBusinessModule{
		ID:         b.ID,
		Name:       b.Name,
		ModuleType: hospitality_utilities.Bar,
		// Status:       b.Status,
		// IsConfigured: b.IsConfigured,
	}
}

func mapHotelToBusinessModule(h hotel_contracts.HotelReader) orchestrator_dtos.HospitalityBusinessModule {
	return orchestrator_dtos.HospitalityBusinessModule{
		ID:           h.ID,
		Name:         h.Name,
		ModuleType:   hospitality_utilities.Hotel,
		Status:       h.Status,
		IsConfigured: h.IsConfigured,
	}
}

func mapRestaurantToBusinessModule(r *restaurant_contracts.RestaurantDetails) orchestrator_dtos.HospitalityBusinessModule {
	return orchestrator_dtos.HospitalityBusinessModule{
		ID:           r.ID,
		Name:         r.Name,
		ModuleType:   hospitality_utilities.Restaurant,
		Status:       r.Status,
		IsConfigured: r.IsConfigured,
	}
}

/* ---------------------------------- <***> --------------------------------- */

func (s *orchestratorService) ListAllInfraModulesOfProperty(
	ctx context.Context,
	orgID, emplID, propertyID uuid.UUID,
) (orchestrator_dtos.PropertyInfraModulesResponseDto, error) {

	// Empty (non-nil) slices so absent domains serialize as [] rather than null.
	resp := orchestrator_dtos.PropertyInfraModulesResponseDto{
		GeneralInfraModules:     make([]orchestrator_dtos.GeneralInfraModule, 0),
		HospitalityInfraModules: make([]orchestrator_dtos.HospitalityInfraModule, 0),
		HealthCareInfraModules:  make([]orchestrator_dtos.HealthCareInfraModule, 0),
	}

	// -------------------------- Hospitality infra --------------------------
	// Only Kitchen is wired up today. NOTE: kitchen lives directly on the
	// hospitality domain for now; later it folds into
	// hospitalityDomainSvc.InfraModules().Kitchen().
	kitchens, err := s.hospitalityDomainSvc.KitchenService().
		GetKitchensByPropertyID(ctx, propertyID)
	if err != nil {
		return orchestrator_dtos.PropertyInfraModulesResponseDto{},
			fmt.Errorf("list kitchens for property %s: %w", propertyID, err)
	}
	for _, k := range kitchens {
		if k == nil {
			continue
		}
		resp.HospitalityInfraModules = append(resp.HospitalityInfraModules, mapKitchenToInfraModule(k))
	}

	// General infra and Healthcare infra are intentionally skipped for now and
	// returned as empty slices.

	return resp, nil
}

// mapKitchenToInfraModule maps a kitchen read-model to the infra DTO.
func mapKitchenToInfraModule(k *kitchen_contracts.Kitchen_Reader) orchestrator_dtos.HospitalityInfraModule {
	return orchestrator_dtos.HospitalityInfraModule{
		ID:        k.ID,
		Name:      k.Name,
		Status:    k.Status,
		InfraType: hospitality_utilities.Kitchen,
		Meta:      nil, // shape TBD
	}
}

/* ---------------------------------- <***> --------------------------------- */
