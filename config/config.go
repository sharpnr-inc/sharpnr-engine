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
	Cors     CORSConfig
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
		Cors:     loadCORS(app),
	}

	validate(cfg)
	return cfg
}
