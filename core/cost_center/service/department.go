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

// CreateNewDepartmentInCostCenter implements [OrgCostCenterServiceInterface].
//
// The cost center is loaded before anything is written, for three reasons at
// once: it is the tenant guard (a site belonging to another organization must
// read as not-found), it is the lifecycle guard (an INACTIVE or ARCHIVED site
// takes no new departments), and it is where organization_id comes from — the
// column is denormalised from the cost center, so copying it from the caller
// instead would let the two disagree.
func (c *orgCostCenterServiceConstructor) CreateNewDepartmentInCostCenter(
	ctx context.Context,
	orgID, emplID, costCenterID uuid.UUID,
	request orgCostCenterDtos.CreateCostCenterDepartmentRequestDto,
) error {

	name := strings.TrimSpace(request.Name)
	if name == "" {
		return orgCostCenterErrs.DepartmentNameRequiredErr
	}

	costCenter, err := c.repo.GetCostCenterByID(ctx, costCenterID)
	if err != nil {
		return err
	}
	if costCenter.OrganizationID != orgID {
		return orgCostCenterErrs.InvalidCostCenterErr
	}
	if costCenter.Status != orgCostCenterConstants.CostCenterActive {
		return orgCostCenterErrs.CostCenterNotActiveErr
	}

	code := strings.ToUpper(strings.TrimSpace(request.Code))
	if code == "" {
		code = deriveCostCenterCode(name)
	}
	if code == "" {
		// A name with nothing alphanumeric in it derives to nothing; the
		// column is NOT NULL, so say what is missing rather than let the
		// insert fail.
		return orgCostCenterErrs.DepartmentCodeRequiredErr
	}

	exists, err := c.repo.DepartmentCodeExists(ctx, costCenterID, code)
	if err != nil {
		return err
	}
	if exists {
		return orgCostCenterErrs.DepartmentCodeAlreadyExistsErr
	}

	// ⚠ The head is only checked for shape here. InvalidDepartmentHeadErr is
	// also meant to cover an unknown employee and one from another
	// organization, which needs the employee service this module does not
	// hold — wire that check in when it is injected.
	var headEmployeeID *uuid.UUID
	if head := strings.TrimSpace(request.DepartmentHead); head != "" {
		parsed, err := uuid.Parse(head)
		if err != nil {
			return orgCostCenterErrs.InvalidDepartmentHeadErr
		}
		headEmployeeID = &parsed
	}

	department := orgCostCenterModels.CostCenterDepartment{
		ID:              uuid.New(),
		OrganizationID:  costCenter.OrganizationID,
		OrgCostCenterID: costCenter.ID,
		Code:            code,
		Name:            name,
		Description:     strings.TrimSpace(request.Description),
		HeadEmployeeID:  headEmployeeID,
		Status:          orgCostCenterConstants.CostCenterActive,
	}

	return c.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := c.repo.CreateCostCenterDepartment(txCtx, department); err != nil {
			return err
		}

		return c.recordDepartmentAudit(txCtx, orgID, emplID, department.ID,
			orgCostCenterConstants.AuditActionDepartmentCreated,
			describeCostCenterDepartment("created", department, *costCenter),
			map[string]any{
				"name":            department.Name,
				"code":            department.Code,
				"cost_center_id":  costCenter.ID.String(),
				"cost_center":     costCenter.Name,
				"department_head": request.DepartmentHead,
			},
		)
	})
}
