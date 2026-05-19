package jwt_manager

import "errors"

// Sentinel errors returned by all Verify* methods.
//
// Callers should use errors.Is() to distinguish them:
//
//	claims, err := jwtMgr.VerifyAccess(token)
//	if errors.Is(err, jwt_manager.ErrExpiredToken) {
//	    // ask the user to re-authenticate
//	}
//	if errors.Is(err, jwt_manager.ErrInvalidToken) {
//	    // token was tampered with — reject hard
//	}
var (
	// ErrInvalidToken covers tampered signatures, wrong format,
	// wrong token type (e.g. signup token used as access token),
	// and algorithm confusion attacks.
	ErrInvalidToken = errors.New("invalid token")

	// ErrExpiredToken means the token was valid but the "exp" claim
	// has passed. Kept separate from ErrInvalidToken so callers can
	// trigger a re-auth flow rather than a hard rejection.
	ErrExpiredToken = errors.New("token expired")

	// ErrInvalidIssuer means the "iss" claim does not match Config.Issuer.
	ErrInvalidIssuer = errors.New("invalid issuer")

	// ErrInvalidAudience means the "aud" claim does not contain Config.Audience.
	ErrInvalidAudience = errors.New("invalid audience")
)
