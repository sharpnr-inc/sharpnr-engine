// Package orgDepartmentResources is the departments module's mod_resources: it
// describes the Department resource to the Sentinel aggregator and lists an
// org's live departments as policy targets.
package orgDepartmentResources

import orgDepartmentsConstants "sharpnr-engine/core/departments/constants"

// Cedar ids, re-exported from the leaf constants package.
//
// They live there rather than here so the orchestrator can use them in its
// guard calls without importing this package — which imports the service, and
// would cycle. Both sides therefore read the same constant, so enforcement and
// the policy-builder palette cannot drift apart.
const (
	ResourceTypeDepartment = orgDepartmentsConstants.ResourceTypeDepartment

	ActionCanViewDepartment   = orgDepartmentsConstants.ActionCanViewDepartment
	ActionCanCreateDepartment = orgDepartmentsConstants.ActionCanCreateDepartment
	ActionCanEditDepartment   = orgDepartmentsConstants.ActionCanEditDepartment
	ActionCanDeleteDepartment = orgDepartmentsConstants.ActionCanDeleteDepartment
)
