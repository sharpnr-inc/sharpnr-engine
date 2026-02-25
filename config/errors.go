package config

func validate(cfg *Config) {
	must(cfg.Database.Host, "DB_HOST")
	must(cfg.Database.User, "DB_USER")
	must(cfg.Database.Password, "DB_PASSWORD")
	must(cfg.JWT.Secret, "JWT_SECRET")
}
