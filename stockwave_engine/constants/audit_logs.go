package stockwaveEngineConstants

/* -------------------------------------------------------------------------- */
/*                                 Audit log                                  */
/* -------------------------------------------------------------------------- */

// AuditModule is the emitting engine stamped on every stockwave audit entry,
// and the resource types those entries point at.
//
// There are no Sentinel action ids in this package yet: stockwave has no
// mod_resources registration, so the orchestrator has nothing to authorize
// against (see the note on stockwaveOrchestrator). They belong here beside
// these, the day that registration lands.
const (
	AuditModule = "stockwave"

	AuditResourceTypeStockwaveAccount = "stockwave_account"

	AuditActionStockwaveServiceStarted = "stockwave.service.started"
	AuditActionStockwaveServiceResumed = "stockwave.service.resumed"
)
