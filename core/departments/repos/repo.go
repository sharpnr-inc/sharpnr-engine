package orgDepartmentRepo

import (
	"context"
	"errors"
	orgDepartmentsDtos "sharpnr-engine/core/departments/dtos"
	orgDepartmentErrs "sharpnr-engine/core/departments/errs"
	orgDepartmentModels "sharpnr-engine/core/departments/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// conn returns the transaction the caller injected, or the repository's own
// handle when there is none, always bound to the request context so tracing and
// cancellation carry through.
func (o *orgDepartmentRepository) conn(ctx context.Context) *gorm.DB {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = o.db
	}
	return db.WithContext(ctx)
}

/* ---------------------------------- <***> --------------------------------- */

// CreateDepartment implements [OrgDepartmentRepositoryInterface].
func (o *orgDepartmentRepository) CreateDepartment(ctx context.Context, department *orgDepartmentModels.OrgDeparment) error {
	return o.conn(ctx).Create(department).Error
}

/* ---------------------------------- <***> --------------------------------- */

// DeleteDepartment implements [OrgDepartmentRepositoryInterface].
//
// Soft delete: GORM stamps deleted_at, so the row stays readable to anything
// that reports on history while dropping out of every live read and out of the
// partial unique indexes, which frees the name and code for re-use.
func (o *orgDepartmentRepository) DeleteDepartment(ctx context.Context, organizationID, departmentID uuid.UUID) error {
	result := o.conn(ctx).
		Where("organization_id = ? AND id = ?", organizationID, departmentID).
		Delete(&orgDepartmentModels.OrgDeparment{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return orgDepartmentErrs.ErrDepartmentNotFound
	}
	return nil
}

/* ---------------------------------- <***> --------------------------------- */

// GetDepartmentByID implements [OrgDepartmentRepositoryInterface].
//
// A row belonging to another organization is not found here rather than
// returned — the tenant boundary is a predicate on the query.
func (o *orgDepartmentRepository) GetDepartmentByID(ctx context.Context, organizationID, departmentID uuid.UUID) (*orgDepartmentModels.OrgDeparment, error) {
	var department orgDepartmentModels.OrgDeparment

	err := o.conn(ctx).
		Where("organization_id = ? AND id = ?", organizationID, departmentID).
		First(&department).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, orgDepartmentErrs.ErrDepartmentNotFound
		}
		return nil, err
	}

	return &department, nil
}

/* ---------------------------------- <***> --------------------------------- */

// ListDepartments implements [OrgDepartmentRepositoryInterface].
func (o *orgDepartmentRepository) ListDepartments(ctx context.Context, organizationID uuid.UUID) ([]orgDepartmentModels.OrgDeparment, error) {
	var departments []orgDepartmentModels.OrgDeparment

	err := o.conn(ctx).
		Where("organization_id = ?", organizationID).
		Order("name ASC").
		Find(&departments).
		Error

	if err != nil {
		return nil, err
	}

	return departments, nil
}

/* ---------------------------------- <***> --------------------------------- */

// UpdateDepartment implements [OrgDepartmentRepositoryInterface].
//
// Only the fields the caller sent are written: a nil pointer means "leave it
// alone", which is what keeps a partial update from blanking the rest of the
// row. DepartmentHead is the one field that can be cleared — an empty string
// writes NULL, i.e. "nobody runs this department any more".
func (o *orgDepartmentRepository) UpdateDepartment(ctx context.Context, organizationID uuid.UUID, department orgDepartmentsDtos.UpdateDepartmentRequestDto) error {
	updates := map[string]any{}

	if department.Name != nil {
		updates["name"] = *department.Name
	}
	if department.Description != nil {
		updates["description"] = *department.Description
	}
	if department.Code != nil {
		updates["code"] = *department.Code
	}
	if department.DepartmentHead != nil {
		if *department.DepartmentHead == "" {
			updates["head_employee_id"] = nil
		} else {
			headID, err := uuid.Parse(*department.DepartmentHead)
			if err != nil {
				return orgDepartmentErrs.ErrInvalidDepartmentHead
			}
			updates["head_employee_id"] = headID
		}
	}

	// Nothing to write, but the caller still expects the id to have existed.
	if len(updates) == 0 {
		_, err := o.GetDepartmentByID(ctx, organizationID, department.ID)
		return err
	}

	result := o.conn(ctx).
		Model(&orgDepartmentModels.OrgDeparment{}).
		Where("organization_id = ? AND id = ?", organizationID, department.ID).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return orgDepartmentErrs.ErrDepartmentNotFound
	}
	return nil
}

/* ---------------------------------- <***> --------------------------------- */

// DepartmentExists implements [OrgDepartmentRepositoryInterface].
//
// Name is compared case-insensitively, matching the partial unique index on
// lower(name): "Finance" and "finance" are the same department to a human, so
// they are the same department here.
func (o *orgDepartmentRepository) DepartmentExists(ctx context.Context, organizationID uuid.UUID, name, code string, excludeID uuid.UUID) (bool, error) {
	if name == "" && code == "" {
		return false, nil
	}

	query := o.conn(ctx).
		Model(&orgDepartmentModels.OrgDeparment{}).
		Where("organization_id = ?", organizationID)

	switch {
	case name != "" && code != "":
		// Parenthesised in the string rather than relying on the query
		// builder to group it: this has to stay ANDed with the tenant
		// predicate, not OR its way past it.
		query = query.Where("(lower(name) = lower(?) OR code = ?)", name, code)
	case name != "":
		query = query.Where("lower(name) = lower(?)", name)
	default:
		query = query.Where("code = ?", code)
	}

	if excludeID != uuid.Nil {
		query = query.Where("id <> ?", excludeID)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

/* ---------------------------------- <***> --------------------------------- */

// CountEmployeesInDepartment implements [OrgDepartmentRepositoryInterface].
func (o *orgDepartmentRepository) CountEmployeesInDepartment(ctx context.Context, organizationID, departmentID uuid.UUID) (int64, error) {
	var count int64

	err := o.conn(ctx).
		Model(&orgDepartmentModels.OrgDepartmentEmployee{}).
		Where("organization_id = ? AND org_department_id = ?", organizationID, departmentID).
		Count(&count).
		Error

	if err != nil {
		return 0, err
	}

	return count, nil
}
