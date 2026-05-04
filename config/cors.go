package config

type CORSConfig struct {
	WebOrgin []string
}

func loadCORS(app AppConfig) CORSConfig {
	if app.Env == "production" {
		return CORSConfig{
			WebOrgin: []string{
				"https://sharpnr.com",
				"https://www.sharpnr.com",
				"https://auth.sharpnr.com",
				"https://iam.sharpnr.com",
				"https://hr.sharpnr.com",
				"https://accounting.sharpnr.com",
				"https://orchestrator.sharpnr.com",
				"https://hospitality.sharpnr.com",
				"https://silo.sharpnr.com",
				"https://cables.sharpnr.com",
				"https://marketing.sharpnr.com",
				"https://marketplace.sharpnr.com",
				"https://folio.sharpnr.com",
				"https://cortex.sharpnr.com",
			},
		}
	}

	return CORSConfig{
		WebOrgin: []string{
			"https://hospitality.sharpnr.test",
			"https://app.sharpnr.test",
			"https://auth.sharpnr.test",
			"https://iam.sharpnr.test",
			"https://hr.sharpnr.test",
			"https://marketing.sharpnr.test",
			"https://accounting.sharpnr.test",
			"https://orchestrator.sharpnr.test",
		},
	}
}
