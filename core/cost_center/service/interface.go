package orgCostCenterService

import (
	"context"
	orgCostCenterDtos "sharpnr-engine/core/cost_center/dtos"

	"github.com/google/uuid"
)

type OrgCostCenterServiceInterface interface {
	CreateNewCostCenterInOrg(ctx context.Context, orgID, emplID uuid.UUID, request orgCostCenterDtos.CreateCostCenterRequestDto) error
	ListAllCostCentersOfOrg(ctx context.Context, orgID, emplID uuid.UUID) ([]orgCostCenterDtos.CostCenterResponseDto, error)
	GetCostCenterByID(ctx context.Context, orgID, emplID, costCenterID uuid.UUID) (*orgCostCenterDtos.CostCenterResponseDto, error)

	/* ------------------------- Cost center departments ------------------------ */

	// CreateNewDepartmentInCostCenter adds a department to one of the
	// organization's cost centers. A cost center has as many as it needs.
	CreateNewDepartmentInCostCenter(ctx context.Context, orgID, emplID, costCenterID uuid.UUID, request orgCostCenterDtos.CreateCostCenterDepartmentRequestDto) error
}
