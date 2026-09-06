package cortexEngineConstants

/* -------------------------------------------------------------------------- */
/*                                 Audit log                                  */
/* -------------------------------------------------------------------------- */

// AuditModule is the emitting engine stamped on every Cortex audit entry, and
// the resource types those entries point at.
//
// There are no Sentinel action ids in this package yet: Cortex has no
// mod_resources registration, so the orchestrator has nothing to authorize
// against (see the note on cortexOrchestratorService). They belong here beside
// these, the day that registration lands.
const (
	AuditModule = "cortex"

	AuditResourceTypeCortexAccount = "cortex_account"

	AuditActionCortexServiceStarted = "cortex.service.started"
	AuditActionCortexServiceResumed = "cortex.service.resumed"
)
