package orgCostCenterRepo

import (
	"context"
	"errors"
	orgCostCenterErrs "sharpnr-engine/core/cost_center/errs"
	orgCostCenterModels "sharpnr-engine/core/cost_center/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// conn returns the transaction the caller injected, or the repository's own
// handle when there is none, always bound to the request context so tracing and
// cancellation carry through.
func (o *orgCostCenterRepoConstructor) conn(ctx context.Context) *gorm.DB {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = o.db
	}
	return db.WithContext(ctx)
}

/* ---------------------------------- <***> --------------------------------- */

// CreateCostCenter implements [OrgCostCenterRepoInterface].
func (o *orgCostCenterRepoConstructor) CreateCostCenter(ctx context.Context, costCenter orgCostCenterModels.OrgCostCenter) error {
	return o.conn(ctx).Create(&costCenter).Error
}

/* ---------------------------------- <***> --------------------------------- */

// GetAllCostCentersOfOrg implements [OrgCostCenterRepoInterface].
func (o *orgCostCenterRepoConstructor) GetAllCostCentersOfOrg(ctx context.Context, orgID uuid.UUID) ([]orgCostCenterModels.OrgCostCenter, error) {
	var costCenters []orgCostCenterModels.OrgCostCenter

	err := o.conn(ctx).
		Where("organization_id = ?", orgID).
		Order("name ASC").
		Find(&costCenters).
		Error

	if err != nil {
		return nil, err
	}

	return costCenters, nil
}

/* ---------------------------------- <***> --------------------------------- */

// GetCostCenterByID implements [OrgCostCenterRepoInterface].
func (o *orgCostCenterRepoConstructor) GetCostCenterByID(ctx context.Context, costCenterID uuid.UUID) (*orgCostCenterModels.OrgCostCenter, error) {
	var costCenter orgCostCenterModels.OrgCostCenter

	err := o.conn(ctx).
		Where("id = ?", costCenterID).
		First(&costCenter).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, orgCostCenterErrs.CostCenterNotFoundErr
		}
		return nil, err
	}

	return &costCenter, nil
}

/* ---------------------------------- <***> --------------------------------- */

// CostCenterCodeExists implements [OrgCostCenterRepoInterface].
func (o *orgCostCenterRepoConstructor) CostCenterCodeExists(ctx context.Context, orgID uuid.UUID, code string) (bool, error) {
	if code == "" {
		return false, nil
	}

	var count int64

	err := o.conn(ctx).
		Model(&orgCostCenterModels.OrgCostCenter{}).
		Where("organization_id = ? AND code = ?", orgID, code).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

/* -------------------------------------------------------------------------- */
/*                          Cost center departments                           */
/* -------------------------------------------------------------------------- */

// CreateCostCenterDepartment implements [OrgCostCenterRepoInterface].
func (o *orgCostCenterRepoConstructor) CreateCostCenterDepartment(ctx context.Context, department orgCostCenterModels.CostCenterDepartment) error {
	return o.conn(ctx).Create(&department).Error
}

/* ---------------------------------- <***> --------------------------------- */

// DepartmentCodeExists implements [OrgCostCenterRepoInterface].
//
// Scoped to the cost center, not the organization: two sites may each run their
// own "FIN", which is exactly what the partial unique index on
// (org_cost_center_id, code) allows.
func (o *orgCostCenterRepoConstructor) DepartmentCodeExists(ctx context.Context, costCenterID uuid.UUID, code string) (bool, error) {
	if code == "" {
		return false, nil
	}

	var count int64

	err := o.conn(ctx).
		Model(&orgCostCenterModels.CostCenterDepartment{}).
		Where("org_cost_center_id = ? AND code = ?", costCenterID, code).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
