package hrmsEngineConstant

/* -------------------------------------------------------------------------- */
/*                                 Audit log                                  */
/* -------------------------------------------------------------------------- */

// AuditModule is the emitting engine stamped on every HRMS audit entry, and the
// resource types those entries point at.
//
// There are no Sentinel action ids in this package yet: HRMS has no
// mod_resources registration, so the orchestrator has nothing to authorize
// against (see the note on hrmsOrchestratorService). They belong here beside
// these, the day that registration lands.
const (
	AuditModule = "hrms"

	AuditResourceTypeHrmsAccount = "hrms_account"

	AuditActionHrmsServiceStarted = "hrms.service.started"
	AuditActionHrmsServiceResumed = "hrms.service.resumed"
)
