package orgCostCenterConstants

// Audit vocabulary for the cost center module: the emitting engine stamped on
// every entry, the resource those entries point at, and the event names.
const (
	AuditModule = "cost_center"

	AuditResourceType           = "org_cost_center"
	AuditResourceTypeDepartment = "cost_center_department"

	AuditActionCostCenterCreated = "cost_center.created"
	AuditActionDepartmentCreated = "cost_center_department.created"
)

// MaxCostCenterCodeLength caps a derived code at the column width. Cost centers
// and their departments share the same varchar(50).
const MaxCostCenterCodeLength = 50
