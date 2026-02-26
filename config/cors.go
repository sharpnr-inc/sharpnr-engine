package config

type CORSConfig struct {
	WebOrgin []string
}

func loadCORS(app AppConfig) CORSConfig {
	if app.Env == "production" {
		return CORSConfig{
			WebOrgin: []string{
				"https://app.sharpnr.com",
			},
		}
	}

	return CORSConfig{
		WebOrgin: []string{
			"http://localhost:5174",
		},
	}
}
