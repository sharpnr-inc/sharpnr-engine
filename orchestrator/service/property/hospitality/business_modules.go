package hospitality_orchestrator

import (
	"context"
	"fmt"

	hospitality_utilities "sharpnr-engine/business_ops/hospitality/utilities"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"

	"github.com/google/uuid"
)

// SetupBusinessModules provisions the requested hospitality business modules
// for the property, one default instance per unit of quantity.
//
// No guard call here: the caller
// (property_orchestrator.AddBusinessOperationModulesInProperty) authorizes
// canAddBusinessModules on the property before fanning out. Module types and
// quantities are already validated by the controller's request binding.
func (s *HospitalityOrchestrator) SetupBusinessModules(
	ctx context.Context,
	orgID, propertyID uuid.UUID,
	modules []orchestrator_dtos.HospitalityModuleQuantity,
) error {
	for _, m := range modules {
		for i := 0; i < m.Quantity; i++ {
			if err := s.setupModule(ctx, orgID, propertyID, m.Module); err != nil {
				return fmt.Errorf(
					"setup hospitality module %q (instance %d of %d) for property %s: %w",
					m.Module, i+1, m.Quantity, propertyID, err,
				)
			}
		}
	}

	return nil
}

// setupModule provisions a single default instance of one hospitality module.
func (s *HospitalityOrchestrator) setupModule(
	ctx context.Context,
	orgID, propertyID uuid.UUID,
	module hospitality_utilities.HospitalityModulesType,
) error {
	switch module {
	case hospitality_utilities.Hotel:
		_, err := s.domainSvc.HotelService().SetupDefaultHotel(ctx, orgID, propertyID)
		return err

	case hospitality_utilities.Bar:
		_, err := s.domainSvc.BarService().SetupDefaultBar(ctx, orgID, propertyID)
		return err

	case hospitality_utilities.Restaurant:
		_, err := s.domainSvc.RestaurantService().SetupDefaultRestaurant(ctx, orgID, propertyID)
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

// ListBusinessModules returns the hospitality slice of a property's business
// modules: every bar, hotel and restaurant attached to it.
//
// No guard call here: the caller
// (property_orchestrator.ListAllBusinessOpsModules) authorizes canViewProperty
// before collecting each domain's contribution.
func (s *HospitalityOrchestrator) ListBusinessModules(
	ctx context.Context,
	orgID, propertyID uuid.UUID,
) ([]orchestrator_dtos.HospitalityBusinessModule, error) {

	// Non-nil so the caller's JSON response is [] rather than null.
	modules := make([]orchestrator_dtos.HospitalityBusinessModule, 0)

	bars, err := s.domainSvc.BarService().GetBarDetailsOfProperty(ctx, orgID, propertyID)
	if err != nil {
		return nil, fmt.Errorf("list bars for property %s: %w", propertyID, err)
	}
	for _, b := range bars {
		if b == nil {
			continue
		}
		modules = append(modules, mapBarToBusinessModule(b))
	}

	hotels, err := s.domainSvc.HotelService().GetHotelByPropertyID(ctx, propertyID)
	if err != nil {
		return nil, fmt.Errorf("list hotels for property %s: %w", propertyID, err)
	}
	for _, h := range hotels {
		modules = append(modules, mapHotelToBusinessModule(h))
	}

	restaurants, err := s.domainSvc.RestaurantService().GetAllRestaurantOfProperty(ctx, propertyID)
	if err != nil {
		return nil, fmt.Errorf("list restaurants for property %s: %w", propertyID, err)
	}
	for _, r := range restaurants {
		if r == nil {
			continue
		}
		modules = append(modules, mapRestaurantToBusinessModule(r))
	}

	return modules, nil
}
