package orgDepartmentService

import (
	"context"
	"strings"

	orgDepartmentsConstants "sharpnr-engine/core/departments/constants"
	orgDepartmentsDtos "sharpnr-engine/core/departments/dtos"
	orgDepartmentErrs "sharpnr-engine/core/departments/errs"
	orgDepartmentModels "sharpnr-engine/core/departments/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateDepartment implements [OrgDepartmentServiceInterface].
//
// Name and code are unique per organization among live departments; the check
// here gives the caller a precise error instead of the partial unique index's
// constraint violation, which the index still enforces underneath as the race
// backstop.
func (c *coreOrgDepartmentService) CreateDepartment(ctx context.Context, orgID, emplID uuid.UUID, request orgDepartmentsDtos.CreateDepartmentRequestDto) error {
	name := strings.TrimSpace(request.Name)
	if name == "" {
		return orgDepartmentErrs.ErrDepartmentNameRequired
	}

	code := strings.ToUpper(strings.TrimSpace(request.Code))
	if code == "" {
		code = deriveDepartmentCode(name)
	}

	var headEmployeeID *uuid.UUID
	if head := strings.TrimSpace(request.DepartmentHead); head != "" {
		parsed, err := uuid.Parse(head)
		if err != nil {
			return orgDepartmentErrs.ErrInvalidDepartmentHead
		}
		headEmployeeID = &parsed
	}

	exists, err := c.orgDepartmentRepository.DepartmentExists(ctx, orgID, name, code, uuid.Nil)
	if err != nil {
		return err
	}
	if exists {
		return orgDepartmentErrs.ErrDepartmentAlreadyExists
	}

	department := &orgDepartmentModels.OrgDeparment{
		ID:             uuid.New(),
		OrganizationID: orgID,
		Name:           name,
		Description:    strings.TrimSpace(request.Description),
		Code:           code,
		HeadEmployeeID: headEmployeeID,
	}

	return c.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := c.orgDepartmentRepository.CreateDepartment(txCtx, department); err != nil {
			return err
		}

		return c.recordAudit(txCtx, orgID, emplID, department.ID,
			orgDepartmentsConstants.AuditActionDepartmentCreated,
			describeDepartment("created", *department),
			map[string]any{
				"name": department.Name,
				"code": department.Code,
			},
		)
	})
}

/* ---------------------------------- <***> --------------------------------- */

// DeleteDepartment implements [OrgDepartmentServiceInterface].
//
// A department people are still assigned to is not deleted: move them out
// first, so nobody silently loses their department.
func (c *coreOrgDepartmentService) DeleteDepartment(ctx context.Context, orgID, emplID uuid.UUID, DepartmentID uuid.UUID) error {
	department, err := c.orgDepartmentRepository.GetDepartmentByID(ctx, orgID, DepartmentID)
	if err != nil {
		return err
	}

	employeeCount, err := c.orgDepartmentRepository.CountEmployeesInDepartment(ctx, orgID, DepartmentID)
	if err != nil {
		return err
	}
	if employeeCount > 0 {
		return orgDepartmentErrs.ErrDepartmentHasEmployees
	}

	return c.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := c.orgDepartmentRepository.DeleteDepartment(txCtx, orgID, DepartmentID); err != nil {
			return err
		}

		return c.recordAudit(txCtx, orgID, emplID, DepartmentID,
			orgDepartmentsConstants.AuditActionDepartmentDeleted,
			describeDepartment("deleted", *department),
			map[string]any{
				"name": department.Name,
				"code": department.Code,
			},
		)
	})
}

/* ---------------------------------- <***> --------------------------------- */

// GetDepartmentByID implements [OrgDepartmentServiceInterface].
func (c *coreOrgDepartmentService) GetDepartmentByID(ctx context.Context, orgID, emplID uuid.UUID, departmentID uuid.UUID) (orgDepartmentsDtos.DepartmentViewResponseDto, error) {
	department, err := c.orgDepartmentRepository.GetDepartmentByID(ctx, orgID, departmentID)
	if err != nil {
		return orgDepartmentsDtos.DepartmentViewResponseDto{}, err
	}

	return toView(*department), nil
}

/* ---------------------------------- <***> --------------------------------- */

// ListDepartments implements [OrgDepartmentServiceInterface].
func (c *coreOrgDepartmentService) ListDepartments(ctx context.Context, orgID, emplID uuid.UUID) ([]orgDepartmentsDtos.DepartmentViewResponseDto, error) {
	departments, err := c.orgDepartmentRepository.ListDepartments(ctx, orgID)
	if err != nil {
		return nil, err
	}

	views := make([]orgDepartmentsDtos.DepartmentViewResponseDto, 0, len(departments))
	for _, department := range departments {
		views = append(views, toView(department))
	}

	return views, nil
}

/* ---------------------------------- <***> --------------------------------- */

// UpdateDepartment implements [OrgDepartmentServiceInterface].
//
// Partial: only the fields present in the request are written. A rename or a
// new code is checked against the org's other live departments, excluding this
// one so re-sending the current values is not a conflict with itself.
func (c *coreOrgDepartmentService) UpdateDepartment(ctx context.Context, orgID, emplID uuid.UUID, updateReq orgDepartmentsDtos.UpdateDepartmentRequestDto) error {
	department, err := c.orgDepartmentRepository.GetDepartmentByID(ctx, orgID, updateReq.ID)
	if err != nil {
		return err
	}

	if updateReq.Name != nil {
		name := strings.TrimSpace(*updateReq.Name)
		if name == "" {
			return orgDepartmentErrs.ErrDepartmentNameRequired
		}
		updateReq.Name = &name
	}
	if updateReq.Code != nil {
		code := strings.ToUpper(strings.TrimSpace(*updateReq.Code))
		if code == "" {
			code = deriveDepartmentCode(department.Name)
			if updateReq.Name != nil {
				code = deriveDepartmentCode(*updateReq.Name)
			}
		}
		updateReq.Code = &code
	}
	if updateReq.Description != nil {
		description := strings.TrimSpace(*updateReq.Description)
		updateReq.Description = &description
	}
	if updateReq.DepartmentHead != nil {
		head := strings.TrimSpace(*updateReq.DepartmentHead)
		if head != "" {
			if _, err := uuid.Parse(head); err != nil {
				return orgDepartmentErrs.ErrInvalidDepartmentHead
			}
		}
		updateReq.DepartmentHead = &head
	}

	var newName, newCode string
	if updateReq.Name != nil {
		newName = *updateReq.Name
	}
	if updateReq.Code != nil {
		newCode = *updateReq.Code
	}

	if newName != "" || newCode != "" {
		exists, err := c.orgDepartmentRepository.DepartmentExists(ctx, orgID, newName, newCode, department.ID)
		if err != nil {
			return err
		}
		if exists {
			return orgDepartmentErrs.ErrDepartmentAlreadyExists
		}
	}

	return c.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := c.orgDepartmentRepository.UpdateDepartment(txCtx, orgID, updateReq); err != nil {
			return err
		}

		return c.recordAudit(txCtx, orgID, emplID, department.ID,
			orgDepartmentsConstants.AuditActionDepartmentUpdated,
			describeDepartment("updated", *department),
			changedFields(updateReq),
		)
	})
}
