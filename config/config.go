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
	ttl := loadTTL()

	cfg := &Config{
		App:      app,
		Database: loadDatabase(),
		Redis:    loadRedis(),
		JWT:      loadJWT(),
		TTL:      ttl,
		Cookies:  loadCookies(app, ttl),
		SMTP:     loadSMTP(),
		Email:    loadEmail(),
		Cors:     loadCORS(app),
	}

	validate(cfg)
	return cfg
}
