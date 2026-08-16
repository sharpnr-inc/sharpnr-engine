// Package orgCostCenterErrs defines the sentinel errors returned by the cost
// center module. Keeping them in one place lets the service layer signal a
// precise failure and the controllers map it to a status code with errors.Is,
// rather than matching on message text.
package orgCostCenterErrs

import "errors"

/* -------------------------------------------------------------------------- */
/*                          Cost Center — Lookup                              */
/* -------------------------------------------------------------------------- */

var (
	NoCostCenterFound = errors.New("no cost centers found for this organization")

	// CostCenterNotFoundErr is a lookup by id that matched nothing (or matched
	// a soft-deleted row).
	CostCenterNotFoundErr = errors.New("cost center not found")

	// InvalidCostCenterErr is the tenant guard: the row exists but belongs to a
	// different organization. Deliberately distinct from NotFound internally,
	// but both should surface to the caller as a 404 so the API never confirms
	// the existence of another org's cost center.
	InvalidCostCenterErr = errors.New("invalid cost center for this organization")
)

/* -------------------------------------------------------------------------- */
/*                        Cost Center — Validation                            */
/* -------------------------------------------------------------------------- */

var (
	// CostCenterCodeAlreadyExistsErr maps the unique violation on
	// (organization_id, code) among live rows.
	CostCenterCodeAlreadyExistsErr = errors.New("a cost center with this code already exists in the organization")

	CostCenterCodeRequiredErr = errors.New("cost center code is required")
	CostCenterNameRequiredErr = errors.New("cost center name is required")

	InvalidCostCenterTypeErr   = errors.New("invalid cost center type")
	InvalidCostCenterStatusErr = errors.New("invalid cost center status")

	// InvalidTimezoneErr guards the IANA name — an unparseable zone would break
	// every attendance cutoff and payroll period computed for the site.
	InvalidTimezoneErr = errors.New("invalid timezone for the cost center")
)

/* -------------------------------------------------------------------------- */
/*                      Cost Center — Legal Entity Link                       */
/* -------------------------------------------------------------------------- */

var (
	// LegalEntityRequiredErr is raised by the operations that cannot proceed
	// without one — invoicing, statutory filing, tax treatment — not by plain
	// creation, where legal_entity_id is allowed to be null.
	LegalEntityRequiredErr = errors.New("cost center must be linked to a legal entity for this operation")

	// InvalidLegalEntityForOrgErr is the cross-tenant guard on the link: the
	// entity exists but is not this organization's.
	InvalidLegalEntityForOrgErr = errors.New("legal entity does not belong to this organization")
)

/* -------------------------------------------------------------------------- */
/*                       Cost Center — Lifecycle Guards                       */
/* -------------------------------------------------------------------------- */

var (
	// CostCenterNotActiveErr blocks writes against an INACTIVE or ARCHIVED cost
	// center — an archived one stays readable so historical reports resolve,
	// but must not take new departments or postings.
	CostCenterNotActiveErr = errors.New("cost center is not active")

	// CostCenterHasDepartmentsErr blocks deletion while departments still hang
	// off it, so the departments are never orphaned.
	CostCenterHasDepartmentsErr = errors.New("cost center still has departments and cannot be deleted")
)

/* -------------------------------------------------------------------------- */
/*                          Cost Center Departments                           */
/* -------------------------------------------------------------------------- */

var (
	NoDepartmentFoundInCostCenter = errors.New("no departments found for this cost center")

	DepartmentNotFoundErr = errors.New("cost center department not found")

	// InvalidDepartmentForCostCenterErr is the scope guard: the department
	// exists but belongs to a different cost center than the one in the request
	// path.
	InvalidDepartmentForCostCenterErr = errors.New("department does not belong to this cost center")

	// DepartmentCodeAlreadyExistsErr maps the unique violation on
	// (org_cost_center_id, code) among live rows.
	DepartmentCodeAlreadyExistsErr = errors.New("a department with this code already exists in this cost center")

	DepartmentCodeRequiredErr = errors.New("department code is required")
	DepartmentNameRequiredErr = errors.New("department name is required")

	// InvalidDepartmentHeadErr covers both an unknown employee and one from
	// another organization.
	InvalidDepartmentHeadErr = errors.New("invalid department head employee for this organization")

	// DepartmentHasEmployeesErr blocks deletion while employees are still
	// assigned, so no assignment is left pointing at a dead department.
	DepartmentHasEmployeesErr = errors.New("department still has assigned employees and cannot be deleted")
)
