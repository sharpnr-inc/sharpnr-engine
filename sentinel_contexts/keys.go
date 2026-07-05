package sentinelContexts

// Canonical Cedar attribute keys. These are the ONLY names policies should be
// authored against; the engine stamps every one of them on every evaluation
// (zero values included), so a condition can never silently no-op because a
// key is missing — in Cedar, reading a missing attribute errors and the whole
// policy is skipped. They are embedded in stored policy text, so they must
// stay stable once released.

// Request-environment keys, read in policy text as `context.<key>`.
const (
	ContextKeyChannel     = "channel"     // "web" | "mobile" (empty when unauthenticated paths evaluate)
	ContextKeyIP          = "ip"          // client IP as a string
	ContextKeyRequestedAt = "requestedAt" // evaluation time (datetime)
	ContextKeyMFAVerified = "mfaVerified" // whether the session completed MFA
)

// Principal (user) keys, read in policy text as `principal.<key>`. The engine
// also stamps `principal.orgId` itself — that one is engine-owned tenancy, not
// settable through this package.
const (
	PrincipalKeyDepartment  = "department"
	PrincipalKeyDesignation = "designation"
)

// Channel values for RequestContext.Channel / `context.channel`.
const (
	ChannelWeb    = "web"
	ChannelMobile = "mobile"
)
