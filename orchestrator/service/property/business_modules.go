package property_orchestrator

import (
	"context"

	orgPropertyResources "sharpnr-engine/core/org_property/mod_resources"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"

	"github.com/google/uuid"
)

/* ---------------------------------- <***> --------------------------------- */

// AddBusinessOperationModulesInProperty attaches the requested business
// operation modules to the property by handing each domain's slice of the
// request to that domain's orchestrator.
func (s *PropertyOrchestrator) AddBusinessOperationModulesInProperty(
	ctx context.Context,
	orgID, emplID, propertyID uuid.UUID,
	req orchestrator_dtos.AddBusinessOpModulesRequestDto,
) error {

	if err := s.guard.Authorize(ctx, orgID, emplID, orgPropertyResources.ActionCanAddBusinessModules, orgPropertyResources.ResourceTypeProperty, propertyID.String()); err != nil {
		return err
	}

	// Request validation is done in the controller, so the module types and
	// quantities here are already known-valid.

	if err := s.HospitalityOrchestrator.SetupBusinessModules(ctx, orgID, propertyID, req.HospitalityModules); err != nil {
		return err
	}

	// Healthcare is on hold for now — req.HealthcareModules is intentionally
	// ignored until that domain is wired up.

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

// ListAllBusinessOpsModules assembles the property's business modules from
// every business domain attached to it.
func (s *PropertyOrchestrator) ListAllBusinessOpsModules(
	ctx context.Context,
	orgID, emplID, propertyID uuid.UUID,
) (orchestrator_dtos.PropertyBusinessOperationModulesResponseDto, error) {

	if err := s.guard.Authorize(ctx, orgID, emplID, orgPropertyResources.ActionCanViewProperty, orgPropertyResources.ResourceTypeProperty, propertyID.String()); err != nil {
		return orchestrator_dtos.PropertyBusinessOperationModulesResponseDto{}, err
	}

	hospitalityModules, err := s.HospitalityOrchestrator.ListBusinessModules(ctx, orgID, propertyID)
	if err != nil {
		return orchestrator_dtos.PropertyBusinessOperationModulesResponseDto{}, err
	}

	// Healthcare domain not ready yet — the slice stays empty (but non-nil, so
	// it serializes as [] rather than null).
	return orchestrator_dtos.PropertyBusinessOperationModulesResponseDto{
		HospitalityModules: hospitalityModules,
		HealthCareModules:  make([]orchestrator_dtos.HealthCareBusinessModule, 0),
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */
