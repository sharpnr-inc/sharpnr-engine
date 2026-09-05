package folioEngineConstants

/* -------------------------------------------------------------------------- */
/*                                 Audit log                                  */
/* -------------------------------------------------------------------------- */

// AuditModule is the emitting engine stamped on every Folio audit entry, and
// the resource types those entries point at.
//
// There are no Sentinel action ids in this package yet: Folio has no
// mod_resources registration, so the orchestrator has nothing to authorize
// against (see the note on folioOrchestratorService). They belong here beside
// these, the day that registration lands.
const (
	AuditModule = "folio"

	AuditResourceTypeFolioAccount = "folio_account"

	AuditActionFolioServiceStarted = "folio.service.started"
	AuditActionFolioServiceResumed = "folio.service.resumed"
)
