package config

import mailer "sharpnr-engine/sharp_mailer"

type EmailConfig struct {
	Identities map[mailer.Type]mailer.Identity
}

func loadEmail() EmailConfig {
	return EmailConfig{
		Identities: map[mailer.Type]mailer.Identity{
			mailer.OTP: {
				FromAddress: mustGet("EMAIL_FROM_OTP"),
			},
			mailer.Support: {
				FromAddress: mustGet("EMAIL_FROM_SUPPORT"),
			},
			mailer.Billing: {
				FromAddress: mustGet("EMAIL_FROM_BILLING"),
			},
			mailer.Notification: {
				FromAddress: mustGet("EMAIL_FROM_NOTIFICATION"),
			},
		},
	}
}
