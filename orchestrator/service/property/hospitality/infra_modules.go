package hospitality_orchestrator

import (
	"context"
	"fmt"

	hospitality_utilities "sharpnr-engine/business_ops/hospitality/utilities"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"

	"github.com/google/uuid"
)

// SetupInfraModules provisions the requested hospitality infra for the property.
//
// Scope for now: only Kitchen is wired up. Any other hospitality infra type is
// skipped rather than failing the whole request, until its setup service exists.
//
// No guard call here: the caller
// (property_orchestrator.AddInfraModulesInProperty) authorizes
// canAddInfraModules on the property before fanning out.
func (s *HospitalityOrchestrator) SetupInfraModules(
	ctx context.Context,
	orgID, propertyID uuid.UUID,
	modules []orchestrator_dtos.HospitalityInfraModuleQuantity,
) error {
	for _, m := range modules {
		if m.Module != hospitality_utilities.Kitchen {
			// Not yet supported — skip rather than fail the whole request.
			continue
		}

		for i := 0; i < m.Quantity; i++ {
			err := s.domainSvc.KitchenService().
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

// ListInfraModules returns the hospitality slice of a property's infra modules.
//
// Only Kitchen is wired up today. NOTE: kitchen lives directly on the
// hospitality domain for now; later it folds into domainSvc.InfraModules().Kitchen().
//
// No guard call here: the caller
// (property_orchestrator.ListAllInfraModulesOfProperty) authorizes
// canViewProperty before collecting each domain's contribution.
func (s *HospitalityOrchestrator) ListInfraModules(
	ctx context.Context,
	propertyID uuid.UUID,
) ([]orchestrator_dtos.HospitalityInfraModule, error) {

	// Non-nil so the caller's JSON response is [] rather than null.
	modules := make([]orchestrator_dtos.HospitalityInfraModule, 0)

	kitchens, err := s.domainSvc.KitchenService().GetKitchensByPropertyID(ctx, propertyID)
	if err != nil {
		return nil, fmt.Errorf("list kitchens for property %s: %w", propertyID, err)
	}
	for _, k := range kitchens {
		if k == nil {
			continue
		}
		modules = append(modules, mapKitchenToInfraModule(k))
	}

	return modules, nil
}
