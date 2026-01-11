package config

import "time"

type TTLConfig struct {
	OTP     time.Duration
	Session time.Duration
	Cookie  time.Duration
}

func loadTTL() TTLConfig {
	return TTLConfig{
		OTP:     getEnvAsDuration("OTP_TTL", "5m"),
		Session: getEnvAsDuration("SESSION_TTL", "24h"),
		Cookie:  getEnvAsDuration("COOKIE_TTL", "168h"),
	}
}
