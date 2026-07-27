package config

import "time"

// TTLConfig holds business-level lifetimes that are not owned by another
// config block. Token lifetimes live on JwtManagerConfigs; cookie Max-Age is
// derived from the lifetime of the token each cookie carries (see cookieConfig.go).
//
// Only add a field here when something actually reads it — a TTL that nothing
// consumes is worse than no TTL at all, because it reads like a working knob.
type TTLConfig struct {
	// SignupStaging is how long an organization signup staging record survives
	// in Redis (key: org:signup:staging:{id} plus its email/phone reverse
	// lookups). The staging record carries the signup OTP, and nothing else
	// enforces OTP expiry — VerifySlugCreateDefaultOrg only checks that the key
	// still exists — so this value *is* the OTP validity window.
	SignupStaging time.Duration
}

func loadTTL() TTLConfig {
	return TTLConfig{
		SignupStaging: getEnvAsDuration("SIGNUP_STAGING_TTL", "40m"),
	}
}
