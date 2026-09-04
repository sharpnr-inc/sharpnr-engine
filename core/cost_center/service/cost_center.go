package orgCostCenterService

import (
	"context"
	"strings"

	orgCostCenterConstants "sharpnr-engine/core/cost_center/constants"
	orgCostCenterDtos "sharpnr-engine/core/cost_center/dtos"
	orgCostCenterErrs "sharpnr-engine/core/cost_center/errs"
	orgCostCenterModels "sharpnr-engine/core/cost_center/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateNewCostCenterInOrg implements [OrgCostCenterServiceInterface].
//
// A new site starts ACTIVE and unlinked from any legal entity: the entity is
// what decides tax treatment and is often registered after the office is stood
// up, which is why the column is nullable. Linking it is a later operation.
func (c *orgCostCenterServiceConstructor) CreateNewCostCenterInOrg(
	ctx context.Context,
	orgID, emplID uuid.UUID,
	request orgCostCenterDtos.CreateCostCenterRequestDto,
) error {

	name := strings.TrimSpace(request.Name)
	if name == "" {
		return orgCostCenterErrs.CostCenterNameRequiredErr
	}

	code := strings.ToUpper(strings.TrimSpace(request.Code))
	if code == "" {
		code = deriveCostCenterCode(name)
	}
	if code == "" {
		// A name with nothing alphanumeric in it derives to nothing; the
		// column is NOT NULL, so say what is missing rather than let the
		// insert fail.
		return orgCostCenterErrs.CostCenterCodeRequiredErr
	}

	exists, err := c.repo.CostCenterCodeExists(ctx, orgID, code)
	if err != nil {
		return err
	}
	if exists {
		return orgCostCenterErrs.CostCenterCodeAlreadyExistsErr
	}

	costCenter := orgCostCenterModels.OrgCostCenter{
		ID:             uuid.New(),
		OrganizationID: orgID,
		Name:           name,
		Description:    strings.TrimSpace(request.Description),
		Code:           code,
		Status:         orgCostCenterConstants.CostCenterActive,
		AddressLine:    strings.TrimSpace(request.Address),
		City:           strings.TrimSpace(request.City),
		State:          strings.TrimSpace(request.State),
		Country:        strings.TrimSpace(request.Contry),
	}

	return c.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := c.repo.CreateCostCenter(txCtx, costCenter); err != nil {
			return err
		}

		return c.recordAudit(txCtx, orgID, emplID, costCenter.ID,
			orgCostCenterConstants.AuditActionCostCenterCreated,
			describeCostCenter("created", costCenter),
			map[string]any{
				"name":    costCenter.Name,
				"code":    costCenter.Code,
				"city":    costCenter.City,
				"country": costCenter.Country,
			},
		)
	})
}

/* ---------------------------------- <***> --------------------------------- */

// ListAllCostCentersOfOrg implements [OrgCostCenterServiceInterface].
//
// An organization with no cost centers is an empty list, not an error: an org
// on the department-only track never creates one, and that is a normal state
// rather than a failure. NoCostCenterFound stays for the flows that genuinely
// cannot proceed without a site.
func (c *orgCostCenterServiceConstructor) ListAllCostCentersOfOrg(
	ctx context.Context,
	orgID, emplID uuid.UUID,
) ([]orgCostCenterDtos.CostCenterResponseDto, error) {

	costCenters, err := c.repo.GetAllCostCentersOfOrg(ctx, orgID)
	if err != nil {
		return nil, err
	}

	views := make([]orgCostCenterDtos.CostCenterResponseDto, 0, len(costCenters))
	for _, costCenter := range costCenters {
		views = append(views, toView(costCenter))
	}

	return views, nil
}

/* ---------------------------------- <***> --------------------------------- */

// GetCostCenterByID implements [OrgCostCenterServiceInterface].
//
// The repo looks the row up by id alone, so the tenant check is here: a cost
// center belonging to another organization is InvalidCostCenterErr, which the
// controller answers as a 404 like a genuine miss — the API must never confirm
// that another org's id exists.
func (c *orgCostCenterServiceConstructor) GetCostCenterByID(
	ctx context.Context,
	orgID, emplID, costCenterID uuid.UUID,
) (*orgCostCenterDtos.CostCenterResponseDto, error) {

	costCenter, err := c.repo.GetCostCenterByID(ctx, costCenterID)
	if err != nil {
		return nil, err
	}
	if costCenter.OrganizationID != orgID {
		return nil, orgCostCenterErrs.InvalidCostCenterErr
	}

	view := toView(*costCenter)

	return &view, nil
}
