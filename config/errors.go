package config

import "log"

// minJWTSecretBytes is the smallest HMAC-SHA256 key we accept. HS256 keys
// shorter than the 256-bit digest weaken the signature, and a short secret is
// the kind of thing that gets pasted in "just for now" and then deployed.
const minJWTSecretBytes = 32

func validate(cfg *Config) {
	must(cfg.Database.Host, "DB_HOST")
	must(cfg.Database.Name, "DB_NAME")
	must(cfg.Database.User, "DB_USER")
	must(cfg.Database.Password, "DB_PASSWORD")

	must(cfg.Redis.Host, "REDIS_HOST")
	must(cfg.Redis.Port, "REDIS_PORT")

	must(cfg.JWT.Secret, "JWT_SECRET")
	if len(cfg.JWT.Secret) < minJWTSecretBytes {
		log.Fatalf(
			"❌ JWT_SECRET is too short: %d bytes, need at least %d",
			len(cfg.JWT.Secret), minJWTSecretBytes,
		)
	}

	// An access token that outlives its refresh token makes the refresh flow
	// pointless — the short-lived credential becomes the long-lived one.
	if cfg.JWT.AccessTTL >= cfg.JWT.RefreshTTL {
		log.Fatalf(
			"❌ JWT_ACCESS_TTL (%s) must be shorter than JWT_REFRESH_TTL (%s)",
			cfg.JWT.AccessTTL, cfg.JWT.RefreshTTL,
		)
	}

	mustPositiveDuration(cfg.JWT.AccessTTL, "JWT_ACCESS_TTL")
	mustPositiveDuration(cfg.JWT.RefreshTTL, "JWT_REFRESH_TTL")
	mustPositiveDuration(cfg.JWT.SignupTTL, "JWT_SIGNUP_TTL")
	mustPositiveDuration(cfg.JWT.WebCookieTokenTTL, "JWT_AUTH_WEB_COOKIE_TTL")

	// A non-positive TTL on a Redis SET means "no expiry" — the signup staging
	// record, OTP and hashed password included, would live forever.
	mustPositiveDuration(cfg.TTL.SignupStaging, "SIGNUP_STAGING_TTL")

	// Warn rather than fail: this is legitimate when Postgres is on the same
	// host (as DB_HOST=127.0.0.1 currently is), and refusing to boot over it
	// would take production down on the next release. Loud enough to act on.
	if cfg.App.Env == EnvProduction && cfg.Database.SSLMode == "disable" {
		log.Printf("⚠️  DB_SSLMODE=disable in production — database traffic is unencrypted")
	}
}
