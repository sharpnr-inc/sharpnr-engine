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

	/* ------------------------- Cost center departments ------------------------ */

	// CreateCostCenterDepartment adds one department to a cost center. A cost
	// center has as many as it needs.
	CreateCostCenterDepartment(ctx context.Context, department orgCostCenterModels.CostCenterDepartment) error

	// DepartmentCodeExists reports whether a live department in this cost
	// center already uses the code. Scoped to the cost center, not the org:
	// two sites may each run their own "FIN", which is what the partial unique
	// index on (org_cost_center_id, code) allows.
	DepartmentCodeExists(ctx context.Context, costCenterID uuid.UUID, code string) (bool, error)
}
