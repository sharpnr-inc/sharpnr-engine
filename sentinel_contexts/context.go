// Package sentinelContexts is the rigid context vocabulary for Sentinel
// authorization. It defines the two typed fact bags an evaluation can carry —
// RequestContext (the environment of this one request → Cedar `context.*`) and
// UserContext (who the principal is → Cedar `principal.*`) — plus the
// context.Context carrier the auth middlewares use to hand them to the engine.
//
// The flow mirrors the dbtx pattern: middleware builds AuthzContext once per
// request and Injects it; every service that calls engine.Authorize(ctx, ...)
// gets it for free, with no signature changes; the engine Extracts and stamps
// the canonical keys itself. Leaf package — it imports only the standard
// library, so anything may depend on it.
package sentinelContexts

import (
	"context"
	"time"
)

// RequestContext is the environment of one request. The engine exposes it to
// policies as `context.<key>` (see keys.go). Zero values are meaningful and
// always stamped: an un-injected evaluation sees channel "" and mfaVerified
// false — conditions evaluate to false rather than erroring the policy.
type RequestContext struct {
	Channel     string    // ChannelWeb | ChannelMobile
	IPAddress   string    // client IP
	RequestedAt time.Time // when the request happened; zero → engine uses evaluation time
	MFAVerified bool      // whether the session completed MFA
}

// CedarKeys returns the canonical `context.*` map with every key present.
func (rc RequestContext) CedarKeys() map[string]any {
	return map[string]any{
		ContextKeyChannel:     rc.Channel,
		ContextKeyIP:          rc.IPAddress,
		ContextKeyRequestedAt: rc.RequestedAt,
		ContextKeyMFAVerified: rc.MFAVerified,
	}
}

// UserContext is who the principal is, beyond their id: HR-ish facts policies
// may condition on. The engine exposes it as `principal.<key>` attributes,
// next to the engine-owned `principal.orgId`. Populated from what the session
// already carries — do not put DB lookups behind it on the authorize hot path.
type UserContext struct {
	Department  string
	Designation string
}

// CedarKeys returns the canonical `principal.*` map with every key present.
func (uc UserContext) CedarKeys() map[string]any {
	return map[string]any{
		PrincipalKeyDepartment:  uc.Department,
		PrincipalKeyDesignation: uc.Designation,
	}
}

// AuthzContext bundles both fact bags for one request.
type AuthzContext struct {
	Request RequestContext
	User    UserContext
}

/* ---------------------------------- <***> --------------------------------- */

type ctxKeyType struct{}

var ctxKey = ctxKeyType{}

// Inject returns a context carrying the AuthzContext. Called once per request
// by the auth middlewares.
func Inject(ctx context.Context, authz AuthzContext) context.Context {
	return context.WithValue(ctx, ctxKey, authz)
}

// Extract returns the AuthzContext carried by ctx, if any. The engine treats a
// missing one as all-zero facts (keys still stamped).
func Extract(ctx context.Context) (AuthzContext, bool) {
	authz, ok := ctx.Value(ctxKey).(AuthzContext)
	return authz, ok
}
