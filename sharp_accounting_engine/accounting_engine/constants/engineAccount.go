package accountingconstants

/* -------------------------------------------------------------------------- */
/*                                 Audit log                                  */
/* -------------------------------------------------------------------------- */

// AuditModule is the emitting engine stamped on every accounting audit entry,
// and the resource types those entries point at.
//
// There are no Sentinel action ids here yet: the accounting engine has no
// mod_resources registration, so the orchestrator has nothing to authorize
// against (see the note on accountingOrchestratorService). They belong here
// beside these, the day that registration lands.
const (
	AuditModule = "accounting"

	AuditResourceTypeAccountingAccount = "accounting_account"

	AuditActionAccountingServiceStarted = "accounting.service.started"
	AuditActionAccountingServiceResumed = "accounting.service.resumed"
)
