package jwt_manager

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

// --------------------------------------------------------------------------
// verifier.go — token verification
//
// All Verify* methods follow the same pattern:
//   1. Parse and cryptographically verify the token via parse()
//   2. Check the Subject claim matches the expected token type
//   3. Return the typed claims struct
//
// Callers distinguish errors with errors.Is():
//   errors.Is(err, ErrExpiredToken)    → re-authenticate
//   errors.Is(err, ErrInvalidToken)    → reject hard (possible tampering)
//   errors.Is(err, ErrInvalidIssuer)   → wrong service issued this token
//   errors.Is(err, ErrInvalidAudience) → token not intended for this service
// --------------------------------------------------------------------------

// parse is the shared verification core used by all Verify* methods.
// It handles:
//   - Cryptographic signature verification (HMAC-SHA256)
//   - Algorithm confusion attack prevention (rejects alg:none and RS*/ES*)
//   - Expiry vs tampering distinction (ErrExpiredToken vs ErrInvalidToken)
//   - Issuer and audience validation
func (m *manager) parse(tokenStr string, claims jwt.Claims) error {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			// Enforce HMAC — reject tokens that claim to use a different
			// algorithm. Without this check an attacker can switch the
			// header to alg:none or an asymmetric algorithm to bypass
			// signature verification entirely.
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrInvalidToken
			}
			return []byte(m.cfg.Secret), nil
		},
	)

	if err != nil {
		// Distinguish expiry from tampering so callers can handle them differently.
		// An expired token was once valid — trigger re-auth.
		// A tampered token was never valid — reject hard.
		if errors.Is(err, jwt.ErrTokenExpired) {
			return ErrExpiredToken
		}
		return ErrInvalidToken
	}

	if !token.Valid {
		return ErrInvalidToken
	}

	// Verify the issuer matches our service.
	// Prevents tokens issued by a foreign service from being accepted.
	issuer, err := claims.GetIssuer()
	if err != nil || issuer != m.cfg.Issuer {
		return ErrInvalidIssuer
	}

	// Verify the audience contains our expected value.
	// Prevents tokens intended for another service from being accepted here.
	audiences, err := claims.GetAudience()
	if err != nil {
		return ErrInvalidAudience
	}
	for _, aud := range audiences {
		if aud == m.cfg.Audience {
			return nil // ✅ all checks passed
		}
	}
	return ErrInvalidAudience
}

// checkSubject guards against cross-type token misuse.
// Example: a signup token should never be accepted by VerifyAccess.
// We issue tokens with a Subject set to their type ("access", "refresh", etc.)
// and check it here before returning the claims to the caller.
func checkSubject(claims jwt.Claims, expected string) error {
	sub, err := claims.GetSubject()
	if err != nil || sub != expected {
		return ErrInvalidToken
	}
	return nil
}

// VerifyAccess verifies a mobile access token.
// Returns ErrExpiredToken if the token has expired so the caller
// can trigger the refresh flow rather than forcing a full re-login.
func (m *manager) VerifyAccess(token string) (*AccessClaims, error) {
	claims := &AccessClaims{}
	if err := m.parse(token, claims); err != nil {
		return nil, err
	}
	if err := checkSubject(claims, "access"); err != nil {
		return nil, err
	}
	return claims, nil
}

// VerifyRefresh verifies a mobile refresh token.
// On success the caller should:
//   1. Look up the SessionID in the sessions table to check it hasn't been revoked
//   2. Issue a new access token
//   3. Optionally rotate the refresh token (recommended)
func (m *manager) VerifyRefresh(token string) (*RefreshClaims, error) {
	claims := &RefreshClaims{}
	if err := m.parse(token, claims); err != nil {
		return nil, err
	}
	if err := checkSubject(claims, "refresh"); err != nil {
		return nil, err
	}
	return claims, nil
}

// VerifySignup verifies an onboarding token.
// Called by signup-gated routes to confirm the user completed OTP verification.
func (m *manager) VerifySignup(token string) (*SignupClaims, error) {
	claims := &SignupClaims{}
	if err := m.parse(token, claims); err != nil {
		return nil, err
	}
	if err := checkSubject(claims, "signup"); err != nil {
		return nil, err
	}
	return claims, nil
}

// VerifyWebCookie verifies the HttpOnly cookie token for the web frontend.
// Called by WebAuthMiddleware on every authenticated web request.
// If this returns ErrExpiredToken, the middleware returns 401 and the
// frontend redirects to login.
func (m *manager) VerifyWebCookie(token string) (*WebCookieTokenClaims, error) {
	claims := &WebCookieTokenClaims{}
	if err := m.parse(token, claims); err != nil {
		return nil, err
	}
	if err := checkSubject(claims, "web_cookie"); err != nil {
		return nil, err
	}
	return claims, nil
}