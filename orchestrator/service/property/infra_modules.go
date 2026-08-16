package property_orchestrator

import (
	"context"

	orgPropertyResources "sharpnr-engine/core/org_property/mod_resources"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"

	"github.com/google/uuid"
)

/* ---------------------------------- <***> --------------------------------- */

// AddInfraModulesInProperty attaches the requested infrastructure modules to
// the property by handing each domain's slice of the request to that domain's
// orchestrator.
//
// Scope for now: only Hospitality infra is provisioned. General infra and
// Healthcare infra are intentionally ignored even if present in the request
// body, until those domains exist.
func (s *PropertyOrchestrator) AddInfraModulesInProperty(
	ctx context.Context,
	orgID, emplID, propertyID uuid.UUID,
	request orchestrator_dtos.AddInfraModulesRequestDto,
) error {

	if err := s.guard.Authorize(ctx, orgID, emplID, orgPropertyResources.ActionCanAddInfraModules, orgPropertyResources.ResourceTypeProperty, propertyID.String()); err != nil {
		return err
	}

	return s.HospitalityOrchestrator.SetupHospitalityInfraModules(ctx, orgID, propertyID, request.HospitalityInfra)
}

/* ---------------------------------- <***> --------------------------------- */

// ListAllInfraModulesOfProperty assembles the property's infra modules from
// every domain attached to it.
func (s *PropertyOrchestrator) ListAllInfraModulesOfProperty(
	ctx context.Context,
	orgID, emplID, propertyID uuid.UUID,
) (orchestrator_dtos.PropertyInfraModulesResponseDto, error) {

	if err := s.guard.Authorize(ctx, orgID, emplID, orgPropertyResources.ActionCanViewProperty, orgPropertyResources.ResourceTypeProperty, propertyID.String()); err != nil {
		return orchestrator_dtos.PropertyInfraModulesResponseDto{}, err
	}

	hospitalityInfra, err := s.HospitalityOrchestrator.ListHospitalityInfraModules(ctx, propertyID)
	if err != nil {
		return orchestrator_dtos.PropertyInfraModulesResponseDto{}, err
	}

	// General and Healthcare infra are not wired up yet — their slices stay
	// empty (but non-nil, so they serialize as [] rather than null).
	return orchestrator_dtos.PropertyInfraModulesResponseDto{
		GeneralInfraModules:     make([]orchestrator_dtos.GeneralInfraModule, 0),
		HospitalityInfraModules: hospitalityInfra,
		HealthCareInfraModules:  make([]orchestrator_dtos.HealthCareInfraModule, 0),
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */
