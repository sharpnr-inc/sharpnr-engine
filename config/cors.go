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
				"https://marketplace.sharpnr.com",
				"https://folio.sharpnr.com",
				"https://cortex.sharpnr.com",
				"https://ivacore.sharpnr.com",
				"https://sharpcrm.sharpnr.com",
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
			"https://accounting.sharpnr.test",
			"https://orchestrator.sharpnr.test",
			"https://cables.sharpnr.test",
			"https://silo.sharpnr.test",
			"https://folio.sharpnr.test",
			"https://sharpcrm.sharpnr.test",
			"https://cortex.sharpnr.test",
			"https://marketplace.sharpnr.test",
		},
	}
}
