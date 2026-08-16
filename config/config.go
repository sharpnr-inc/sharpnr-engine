package config

import "sharpnr-engine/pkg/cookieManager"

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Redis    RedisDBConfig
	JWT      JwtManagerConfigs
	TTL      TTLConfig
	Cookies  map[cookieManager.Type]cookieManager.Definition // was: Cookie CookieConfig
	SMTP     SMTPConfig
	Email    EmailConfig
	Cors     CORSConfig
}

func Load() *Config {
	loadEnv()

	app := loadApp()

	// Loaded ahead of the literal because the cookie definitions derive their
	// Max-Age from the JWT TTLs — see loadCookies.
	jwt := loadJWT()

	cfg := &Config{
		App:      app,
		Database: loadDatabase(),
		Redis:    loadRedis(),
		JWT:      jwt,
		TTL:      loadTTL(),
		Cookies:  loadCookies(app, jwt),
		SMTP:     loadSMTP(),
		Email:    loadEmail(),
		Cors:     loadCORS(app),
	}

	validate(cfg)
	return cfg
}
