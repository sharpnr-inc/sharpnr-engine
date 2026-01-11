package config

import "time"

type JwtManagerConfigs struct {
	Secret     string
	Issuer     string
	Audience   string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
	SignupTTL  time.Duration
}

func loadJWT() JwtManagerConfigs {
	return JwtManagerConfigs{
		Secret:     mustGet("JWT_SECRET"),
		Issuer:     getEnv("JWT_ISSUER", "sharpnr"),
		Audience:   getEnv("JWT_AUDIENCE", "sharpnr-client"),
		AccessTTL:  getEnvAsDuration("JWT_ACCESS_TTL", "15m"),
		RefreshTTL: getEnvAsDuration("JWT_REFRESH_TTL", "168h"),
		SignupTTL:  getEnvAsDuration("JWT_SIGNUP_TTL", "24h"),
	}
}
