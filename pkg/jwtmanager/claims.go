package jwt_manager

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// --------------------------------------------------------------------------
// Claims are the payloads embedded inside each JWT.
// Every token type has its own claims struct so the verifier knows exactly
// what fields to expect — and can reject a token issued for a different
// purpose (e.g. a signup token used as an auth token).
//
// Shared conventions across all claims:
//   - UserID  → "uid"    (who the token belongs to)
//   - OrgID   → "org_id" (which organisation)
//   - Subject → set automatically by buildRegistered() to the token type
//               e.g. "access", "refresh", "signup", "web_cookie"
//
// uuid.UUID marshals to a standard UUID string in JSON automatically,
// so the wire format is identical to using plain strings.
// --------------------------------------------------------------------------

// AccessClaims is embedded in short-lived access tokens (default 15 min).
// Used by the mobile API — sent as a Bearer token in the Authorization header.
type AccessClaims struct {
	UserID uuid.UUID `json:"uid"`
	OrgID  uuid.UUID `json:"org_id"`
	Role   string    `json:"role"`
	jwt.RegisteredClaims
}

// RefreshClaims is embedded in long-lived refresh tokens (default 24 h).
// Used by the mobile API to obtain a new access token without re-login.
// SessionID ties the token to a specific server-side session record so it
// can be revoked individually (e.g. "log out this device").
type RefreshClaims struct {
	UserID    uuid.UUID `json:"uid"`
	OrgID     uuid.UUID `json:"org_id"`
	SessionID string    `json:"sid"` // string — not necessarily a UUID
	jwt.RegisteredClaims
}

// WebCookieTokenClaims is embedded in the HttpOnly cookie token (default 10 h).
// Used by the web frontend — the browser sends it automatically via cookie,
// the server reads and verifies it in WebAuthMiddleware.
// Intentionally minimal: no role, no session ID — the server looks those up
// from the database when needed.
type WebCookieTokenClaims struct {
	UserID uuid.UUID `json:"uid"`
	OrgID  uuid.UUID `json:"org_id"`
	jwt.RegisteredClaims
}

// SignupClaims is embedded in short-lived onboarding tokens (default 24 h).
// Issued after OTP verification — proves the user completed the first step
// of signup but has not yet created a full account.
// No OrgID — the org does not exist yet at this point in the flow.
type SignupClaims struct {
	UserID uuid.UUID `json:"uid"`
	jwt.RegisteredClaims
}
