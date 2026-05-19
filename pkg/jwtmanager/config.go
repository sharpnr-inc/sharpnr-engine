package jwt_manager

import (
	"time"
)

// Config holds everything the manager needs to issue and verify tokens.
// It is populated from environment variables via config.loadJWT() and
// injected into New() at startup — never read from env inside this package.
type Config struct {
	// Secret is the HMAC-SHA256 signing key shared by all token types.
	// Must be set — New() will return an error if it is empty.
	// Minimum recommended length: 32 bytes.
	Secret string

	// Issuer is embedded in every token ("iss" claim).
	// Verified on every parse — tokens from a different issuer are rejected.
	// Example: "sharpnr"
	Issuer string

	// Audience is embedded in every token ("aud" claim).
	// Verified on every parse — tokens intended for a different audience are rejected.
	// Example: "sharpnr-client"
	Audience string

	// TTLs control how long each token type is valid.
	// These are loaded from env vars with sensible defaults (see config/jwt.go).
	AccessTTL         time.Duration // JWT_ACCESS_TTL         (default 15m)  — mobile access token
	RefreshTTL        time.Duration // JWT_REFRESH_TTL        (default 24h)  — mobile refresh token
	SignupTTL         time.Duration // JWT_SIGNUP_TTL         (default 24h)  — onboarding gate token
	WebCookieTokenTTL time.Duration // JWT_AUTH_WEB_COOKIE_TTL (default 10h) — web HttpOnly cookie token
}
