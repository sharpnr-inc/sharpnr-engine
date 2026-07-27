package config

// Environment names. Anything not listed here is treated as a non-production
// environment: cookie names/domains, CORS origins and gin's mode all key off
// Env == EnvProduction, so an unrecognised value must never fall through to
// production settings by accident.
const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
	EnvDocker      = "docker"
)

type AppConfig struct {
	Env     string
	Port    string
	GinMode string
}

func loadApp() AppConfig {
	// Defaulted rather than read raw: an empty APP_ENV used to produce
	// Env == "", which is not production and so happened to pick the dev
	// branch everywhere — correct by accident. Name the default instead.
	env := getEnv("APP_ENV", EnvDevelopment)

	ginMode := "debug"
	if env == EnvProduction {
		ginMode = "release"
	}

	return AppConfig{
		Env:     env,
		Port:    getEnv("PORT", "8080"),
		GinMode: ginMode,
	}
}
