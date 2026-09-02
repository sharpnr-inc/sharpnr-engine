package orchestrator_service

import (
	"context"

	orgCostCenterDtos "sharpnr-engine/core/cost_center/dtos"
	orgCostCenterService "sharpnr-engine/core/cost_center/service"

	"github.com/google/uuid"
)

// costCenterOrchestrator implements OrgCostCenterOrchestratorServiceInterface.
//
// Every method is a single delegated call, so there is no transaction here
// (rule 2 in doc.go) — the cost center service owns its own consistency, and it
// already commits each write together with the audit entry describing it.
//
// ⚠ No guard.Authorize calls yet (doc.go rule 1): cost centers have no Sentinel
// resource registered, so there is no action id to check against. Reading and
// writing them is currently open to any authenticated employee of the org —
// wire the guard here the day the resource lands.
type costCenterOrchestrator struct {
	coreCostCenterService orgCostCenterService.OrgCostCenterServiceInterface
}

/* ---------------------------------- <***> --------------------------------- */

func (s *costCenterOrchestrator) CreateNewCostCenter(ctx context.Context, orgID, emplID uuid.UUID, request orgCostCenterDtos.CreateCostCenterRequestDto) error {
	return s.coreCostCenterService.CreateNewCostCenterInOrg(ctx, orgID, emplID, request)
}

/* ---------------------------------- <***> --------------------------------- */

func (s *costCenterOrchestrator) ListAllOrgCostCenters(ctx context.Context, orgID, emplID uuid.UUID) ([]orgCostCenterDtos.CostCenterResponseDto, error) {
	return s.coreCostCenterService.ListAllCostCentersOfOrg(ctx, orgID, emplID)
}

/* ---------------------------------- <***> --------------------------------- */

func (s *costCenterOrchestrator) GetOrgCostCenterDetails(ctx context.Context, orgID, emplID, costCenterID uuid.UUID) (*orgCostCenterDtos.CostCenterResponseDto, error) {
	return s.coreCostCenterService.GetCostCenterByID(ctx, orgID, emplID, costCenterID)
}

/* ---------------------------------- <***> --------------------------------- */
