package siloEngineConstants

/* -------------------------------------------------------------------------- */
/*                                 Audit log                                  */
/* -------------------------------------------------------------------------- */

// AuditModule is the emitting engine stamped on every Silo audit entry, and the
// resource types those entries point at.
//
// There are no Sentinel action ids in this package yet: Silo has no
// mod_resources registration, so the orchestrator has nothing to authorize
// against (see the note on siloOrchestratorService). They belong here beside
// these, the day that registration lands.
const (
	AuditModule = "silo"

	AuditResourceTypeSiloAccount = "silo_account"

	AuditActionSiloServiceStarted = "silo.service.started"
	AuditActionSiloServiceResumed = "silo.service.resumed"
)
