package orgCostCenterRepo

import (
	"context"
	orgCostCenterModels "sharpnr-engine/core/cost_center/models"

	"github.com/google/uuid"
)

type OrgCostCenterRepoInterface interface {
	CreateCostCenter(ctx context.Context, costCenter orgCostCenterModels.OrgCostCenter) error
	GetAllCostCentersOfOrg(ctx context.Context, orgID uuid.UUID) ([]orgCostCenterModels.OrgCostCenter, error)

	// GetCostCenterByID looks the row up by id alone. The tenant check belongs
	// to the service, which compares the row's organization against the
	// caller's and answers InvalidCostCenterErr — see the errs package.
	GetCostCenterByID(ctx context.Context, costCenterID uuid.UUID) (*orgCostCenterModels.OrgCostCenter, error)

	// CostCenterCodeExists reports whether a live cost center in the
	// organization already uses this code. The partial unique index on
	// (organization_id, code) is still the race backstop; this is what turns
	// the common case into a precise error instead of a constraint violation.
	CostCenterCodeExists(ctx context.Context, orgID uuid.UUID, code string) (bool, error)
}
