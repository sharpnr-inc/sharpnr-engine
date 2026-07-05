package config

import mailerEngineTypes "sharpnr-engine/sharp_mailer_engine/types"

type EmailConfig struct {
	Identities map[mailerEngineTypes.Type]mailerEngineTypes.Identity
}

func loadEmail() EmailConfig {
	return EmailConfig{
		Identities: map[mailerEngineTypes.Type]mailerEngineTypes.Identity{
			mailerEngineTypes.OTP: {
				FromAddress: mustGet("EMAIL_FROM_OTP"),
			},
			mailerEngineTypes.Support: {
				FromAddress: mustGet("EMAIL_FROM_SUPPORT"),
			},
			mailerEngineTypes.Billing: {
				FromAddress: mustGet("EMAIL_FROM_BILLING"),
			},
			mailerEngineTypes.Notification: {
				FromAddress: mustGet("EMAIL_FROM_NOTIFICATION"),
			},
		},
	}
}
