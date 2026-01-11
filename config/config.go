package config

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Redis    RedisDBConfig
	JWT      JwtManagerConfigs
	TTL      TTLConfig
	Cookie   CookieConfig
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
	}

	validate(cfg)
	return cfg
}
