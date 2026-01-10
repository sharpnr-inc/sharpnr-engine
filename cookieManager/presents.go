package cookieManager

import "net/http"

// registry is the single source of truth for all cookies.
var registry = map[Type]Definition{

	// 🟡 Temporary cookie after OTP verification
	// Allows access to /onboarding
	Signup: {
		Name:     "_signup",
		Scope:    Protected,
		TTL:      60 * 60, // 10 minutes
		SameSite: http.SameSiteLaxMode,
	},

	// 🟢 Authenticated session cookie (refresh token)
	Auth: {
		Name:     "__Host-auth",
		Scope:    Protected,
		TTL:      60 * 60 * 24, // 24 hours (example)
		SameSite: http.SameSiteLaxMode,
	},

	// 🔐 CSRF double-submit cookie
	CSRF: {
		Name:     "XSRF-TOKEN",
		Scope:    Public,
		TTL:      60 * 60 * 24,
		SameSite: http.SameSiteLaxMode,
	},
}
