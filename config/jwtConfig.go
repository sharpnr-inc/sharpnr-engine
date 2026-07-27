package config

import (
	"time"

	jwt_manager "sharpnr-engine/pkg/jwtmanager"
)

// JwtManagerConfigs holds raw JWT settings loaded from environment variables.
// It is stored on Config and converted to jwt_manager.Config via ToManagerConfig()
// before being passed to jwt_manager.New() in main.go.
type JwtManagerConfigs struct {
	Secret            string
	Issuer            string
	Audience          string
	AccessTTL         time.Duration
	RefreshTTL        time.Duration
	SignupTTL         time.Duration
	WebCookieTokenTTL time.Duration
}

// ToManagerConfig converts JwtManagerConfigs into the jwt_manager package's
// own Config type. This keeps the config package as the single place that
// reads from env, while jwt_manager stays independent of the config package.
func (j JwtManagerConfigs) ToManagerConfig() jwt_manager.Config {
	return jwt_manager.Config{
		Secret:            j.Secret,
		Issuer:            j.Issuer,
		Audience:          j.Audience,
		AccessTTL:         j.AccessTTL,
		RefreshTTL:        j.RefreshTTL,
		SignupTTL:         j.SignupTTL,
		WebCookieTokenTTL: j.WebCookieTokenTTL,
	}
}

// Issuer and audience are required rather than defaulted. They are verified on
// every parse, so a deployment that supplies JWT_SECRET but not JWT_ISSUER
// would boot cleanly and then reject every token minted by the previous
// release — a silent, total auth outage with nothing in the logs. Failing at
// startup makes that misconfiguration loud instead.
func loadJWT() JwtManagerConfigs {
	return JwtManagerConfigs{
		Secret:            mustGet("JWT_SECRET"),
		Issuer:            mustGet("JWT_ISSUER"),
		Audience:          mustGet("JWT_AUDIENCE"),
		AccessTTL:         getEnvAsDuration("JWT_ACCESS_TTL", "15m"),
		RefreshTTL:        getEnvAsDuration("JWT_REFRESH_TTL", "24h"),
		SignupTTL:         getEnvAsDuration("JWT_SIGNUP_TTL", "24h"),
		WebCookieTokenTTL: getEnvAsDuration("JWT_AUTH_WEB_COOKIE_TTL", "10h"),
	}
}
