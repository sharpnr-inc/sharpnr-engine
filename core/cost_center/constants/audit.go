package orgCostCenterConstants

// Audit vocabulary for the cost center module: the emitting engine stamped on
// every entry, the resource those entries point at, and the event names.
const (
	AuditModule       = "cost_center"
	AuditResourceType = "org_cost_center"

	AuditActionCostCenterCreated = "cost_center.created"
)

// MaxCostCenterCodeLength caps a derived code at the column width.
const MaxCostCenterCodeLength = 50
