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
}
