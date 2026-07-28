package jwt_manager

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// --------------------------------------------------------------------------
// signer.go — token issuance
//
// All Issue* methods follow the same pattern:
//   1. Build the specific claims struct for that token type
//   2. Attach standard registered claims via buildRegistered()
//   3. Sign with HMAC-SHA256 via sign()
// --------------------------------------------------------------------------

// buildRegistered constructs the standard JWT fields shared by all token types.
//
// subject identifies the token type ("access", "refresh", "signup", "web_cookie").
// It is checked during verification to prevent a token issued for one purpose
// from being accepted by a verifier for a different purpose.
func (m *manager) buildRegistered(subject string, ttl time.Duration) jwt.RegisteredClaims {
	now := m.now() // captured once so IssuedAt and ExpiresAt are always consistent
	return jwt.RegisteredClaims{
		Subject:   subject,
		Issuer:    m.cfg.Issuer,
		Audience:  jwt.ClaimStrings{m.cfg.Audience},
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
	}
}

// sign serialises the claims and signs them with HMAC-SHA256.
// All token types use the same signing algorithm and secret.
// To change the algorithm later (e.g. RS256), update only this function.
func (m *manager) sign(claims jwt.Claims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(m.cfg.Secret))
}

// IssueAccess creates a short-lived access token for the mobile API.
// Carries the user's identity, org, and role so downstream services
// don't need a database lookup on every request.
func (m *manager) IssueAccess(userID, orgID uuid.UUID, role string) (string, error) {
	return m.sign(AccessClaims{
		UserID:           userID,
		OrgID:            orgID,
		Role:             role,
		RegisteredClaims: m.buildRegistered("access", m.cfg.AccessTTL),
	})
}

// IssueRefresh creates a long-lived refresh token for the mobile API.
//
// sessionID must match a row in employee.employee_refresh_session — that is
// what allows a single session to be revoked ("log out this device").
//
// jti identifies this specific token within the session. The caller must store
// it as the session's current jti in the same operation that issues the token,
// otherwise the token is born already-superseded and its first use will look
// like replay.
func (m *manager) IssueRefresh(userID, orgID, sessionID, jti uuid.UUID) (string, error) {
	registered := m.buildRegistered("refresh", m.cfg.RefreshTTL)
	registered.ID = jti.String()

	return m.sign(RefreshClaims{
		UserID:           userID,
		OrgID:            orgID,
		SessionID:        sessionID,
		RegisteredClaims: registered,
	})
}

// IssueSignup creates a short-lived onboarding token.
// Issued after OTP is verified. The token gates the "complete your profile"
// step — without it, the signup routes reject the request.
func (m *manager) IssueSignup(userID uuid.UUID) (string, error) {
	return m.sign(SignupClaims{
		UserID:           userID,
		RegisteredClaims: m.buildRegistered("signup", m.cfg.SignupTTL),
	})
}

// IssueWebAuthCookieToken creates the HttpOnly cookie token for the web frontend.
// Unlike mobile, the web does not use access/refresh token pairs — this single
// token covers the full session and is rolled (re-issued with a fresh TTL) by
// WebAuthMiddleware on every authenticated request so active users never get
// logged out mid-session.
func (m *manager) IssueWebAuthCookieToken(userID, orgID uuid.UUID) (string, error) {
	return m.sign(WebCookieTokenClaims{
		UserID:           userID,
		OrgID:            orgID,
		RegisteredClaims: m.buildRegistered("web_cookie", m.cfg.WebCookieTokenTTL),
	})
}
