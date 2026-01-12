package config

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Redis    RedisDBConfig
	JWT      JwtManagerConfigs
	TTL      TTLConfig
	Cookie   CookieConfig
	SMTP     SMTPConfig
	Email    EmailConfig
}

func Load() *Config {
	loadEnv()

	ttl := loadTTL()
	app := loadApp()

	cfg := &Config{
		App:      loadApp(),
		Database: loadDatabase(),
		Redis:    loadRedis(),
		JWT:      loadJWT(),
		TTL:      loadTTL(),
		Cookie:   loadCookie(app, ttl),
		SMTP:     loadSMTP(),
		Email:    loadEmail(),
	}

	validate(cfg)
	return cfg
}
