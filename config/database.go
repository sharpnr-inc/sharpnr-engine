package config

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	SSLMode  string
}

func loadDatabase() DatabaseConfig {
	return DatabaseConfig{
		Driver:   getEnv("DB_DRIVER", "postgres"),
		Host:     getEnv("DB_HOST", ""),
		Port:     getEnv("DB_PORT", "5432"),
		Name:     getEnv("DB_NAME", ""),
		User:     getEnv("DB_USER", ""),
		Password: getEnv("DB_PASSWORD", ""),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}
