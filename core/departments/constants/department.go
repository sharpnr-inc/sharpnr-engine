package orgDepartmentsConstants

/* -------------------------------------------------------------------------- */
/*                            Sentinel — authorization                        */
/* -------------------------------------------------------------------------- */

// Cedar ids for the org-level department.
//
// They live in this leaf constants package (rather than in mod_resources) so
// the orchestrator's guard calls can use them without importing mod_resources,
// which imports the service — that would cycle. Both sides therefore read the
// same constant, so enforcement and the policy-builder palette cannot drift.
//
// The action ids are exactly what contracts.Standard generates for the type,
// i.e. can<Verb><ResourceType>; keep them in sync with the registration in
// core/departments/mod_resources.
const (
	// ResourceTypeDepartment is the Cedar entity type of an org department.
	ResourceTypeDepartment = "Department"

	ActionCanViewDepartment   = "canViewDepartment"
	ActionCanCreateDepartment = "canCreateDepartment"
	ActionCanEditDepartment   = "canEditDepartment"
	ActionCanDeleteDepartment = "canDeleteDepartment"
)

/* -------------------------------------------------------------------------- */
/*                              Audit log                                     */
/* -------------------------------------------------------------------------- */

// AuditModule is the emitting engine stamped on every department audit entry,
// and AuditResourceType the resource those entries point at.
const (
	AuditModule       = "departments"
	AuditResourceType = "org_department"

	AuditActionDepartmentCreated = "department.created"
	AuditActionDepartmentUpdated = "department.updated"
	AuditActionDepartmentDeleted = "department.deleted"
)

/* -------------------------------------------------------------------------- */
/*                              Department code                               */
/* -------------------------------------------------------------------------- */

// MaxDepartmentCodeLength caps the generated code at the column width.
const MaxDepartmentCodeLength = 50
