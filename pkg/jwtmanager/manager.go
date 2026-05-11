package jwt_manager

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// --------------------------------------------------------------------------
// Manager is the public API for this package.
//
// Web flow (cookie-based):
//   IssueWebAuthCookieToken → stored in HttpOnly cookie by cookieManager
//   VerifyWebCookie         → called by WebAuthMiddleware on every request
//
// Mobile flow (Bearer token):
//   IssueAccess   → short-lived (15 min), sent in Authorization header
//   IssueRefresh  → long-lived (24 h), stored securely on device
//   VerifyAccess  → called by mobile auth middleware
//   VerifyRefresh → called by the token-refresh endpoint
//
// Onboarding flow:
//   IssueSignup  → issued after OTP verification
//   VerifySignup → called by signup-protected routes
// --------------------------------------------------------------------------

type Manager interface {
	// --- Issuers ---

	// IssueAccess creates a short-lived access token for the mobile API.
	IssueAccess(userID, orgID uuid.UUID, role string) (string, error)

	// IssueRefresh creates a long-lived refresh token for the mobile API.
	// sessionID ties the token to a revocable server-side session.
	IssueRefresh(userID, orgID uuid.UUID) (string, error)

	// IssueSignup creates an onboarding token valid for the signup TTL.
	// Issued once OTP is verified; consumed when the account is created.
	IssueSignup(signupSlug uuid.UUID) (string, error)

	// IssueWebAuthCookieToken creates the HttpOnly cookie token for the web frontend.
	// This is the only token type the web uses — there is no separate access/refresh
	// split for web; the cookie is rolled on every request by WebAuthMiddleware.
	IssueWebAuthCookieToken(userID, orgID uuid.UUID) (string, error)

	// --- Verifiers ---

	VerifyAccess(token string) (*AccessClaims, error)
	VerifyRefresh(token string) (*RefreshClaims, error)
	VerifySignup(token string) (*SignupClaims, error)
	VerifyWebCookie(token string) (*WebCookieTokenClaims, error)
}

// manager is the private implementation of Manager.
type manager struct {
	cfg Config

	// now is a function that returns the current time.
	// In production this is always time.Now.
	// In tests it can be replaced with a fixed time to simulate
	// expired tokens without sleeping:
	//
	//   mgr := &manager{cfg: cfg, now: func() time.Time {
	//       return time.Now().Add(-24 * time.Hour) // pretend it's yesterday
	//   }}
	now func() time.Time
}

// New constructs a Manager and validates the config eagerly.
// Returns an error at startup rather than failing silently at runtime.
func New(cfg Config) (Manager, error) {
	if cfg.Secret == "" {
		return nil, errors.New("jwt: Secret must not be empty")
	}
	if cfg.Issuer == "" {
		return nil, errors.New("jwt: Issuer must not be empty")
	}
	if cfg.Audience == "" {
		return nil, errors.New("jwt: Audience must not be empty")
	}

	return &manager{
		cfg: cfg,
		now: time.Now,
	}, nil
}
