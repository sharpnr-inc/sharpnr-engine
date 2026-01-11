package config

import "os"

type AppConfig struct {
	Env     string
	Port    string
	GinMode string
}

func loadApp() AppConfig {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	ginMode := "debug"
	if env == "production" {
		ginMode = "release"
	}

	return AppConfig{
		Env:     env,
		Port:    getEnv("PORT", "8080"),
		GinMode: ginMode,
	}
}
