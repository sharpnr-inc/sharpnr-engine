package cablesEngineConstants

/* -------------------------------------------------------------------------- */
/*                                 Audit log                                  */
/* -------------------------------------------------------------------------- */

// AuditModule is the emitting engine stamped on every cables audit entry, and
// the resource types those entries point at.
//
// There are no Sentinel action ids in this package yet: cables has no
// mod_resources registration, so the orchestrator has nothing to authorize
// against (see the ⚠ note on cablesEngineOrchestrator). They belong here beside
// these, the day that registration lands.
const (
	AuditModule = "cables"

	AuditResourceTypeCablesAccount = "cables_account"
	AuditResourceTypeChannel       = "cables_channel"

	AuditActionCablesServiceStarted = "cables.service.started"
	AuditActionCablesServiceResumed = "cables.service.resumed"
	AuditActionChannelCreated       = "cables.channel.created"
)

/* -------------------------------------------------------------------------- */
/*                                  Channels                                  */
/* -------------------------------------------------------------------------- */

// MaxChannelSlugLength bounds a slug derived from a channel name. It stays
// under the column width so a long name is truncated here, where the caller
// gets a usable slug, rather than by the database.
const MaxChannelSlugLength = 100

// DefaultChannelType and DefaultChannelVisibility are what a create request
// that names neither gets: an ordinary text channel anyone in the cables
// account can find. Private is never the default — a channel nobody can
// discover has to be asked for.
const (
	DefaultChannelType       = ChannelTypeText
	DefaultChannelVisibility = ChannelVisibilityPublic
)
