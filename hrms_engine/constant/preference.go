package hrmsEngineConstant

// HRModel is how the organization runs its HR function.
//
// It is a statement of INTENT, not an access-control switch. Authorization is
// always decided by Sentinel policy: a central HR carries a type-level grant
// (every cost center), a local HR carries an instance-level grant (one cost
// center). HR tables are cost-center-scoped in both models, so switching
// between them is a change of policy attachments and never a data migration.
//
// What this drives instead: what the HR app renders (a cost-center picker or
// not), what policy scope is attached by default when an HR role is created,
// and the warning surfaced when a configuration contradicts the declared model.
type HRModel string

const (
	// HRModelCentralized — one HR function over the whole organization.
	// New HR roles default to a type-level grant covering every cost center.
	HRModelCentralized HRModel = "CENTRALIZED"

	// HRModelDistributed — HR staff sit at individual cost centers, typically
	// reporting to a central HR through the ordinary EmployeeReporting chain.
	// New HR roles default to an instance-level grant and the UI asks which
	// cost centers they cover.
	HRModelDistributed HRModel = "DISTRIBUTED"
)

func (m HRModel) Valid() bool {
	switch m {
	case HRModelCentralized, HRModelDistributed:
		return true
	default:
		return false
	}
}
