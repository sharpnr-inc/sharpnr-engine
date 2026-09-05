package sharpCrmEngineConstants

/* -------------------------------------------------------------------------- */
/*                                 Audit log                                  */
/* -------------------------------------------------------------------------- */

// AuditModule is the emitting engine stamped on every SharpCRM audit entry, and
// the resource types those entries point at.
//
// There are no Sentinel action ids in this package yet: SharpCRM has no
// mod_resources registration, so the orchestrator has nothing to authorize
// against (see the note on sharpCrmOrchestratorService). They belong here
// beside these, the day that registration lands.
const (
	AuditModule = "sharpcrm"

	AuditResourceTypeCrmAccount = "sharpcrm_account"

	AuditActionCrmServiceStarted = "sharpcrm.service.started"
	AuditActionCrmServiceResumed = "sharpcrm.service.resumed"
)
