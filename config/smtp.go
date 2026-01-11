package config

import "time"

type SMTPConfig struct {
	Host           string
	Port           int
	Username       string
	Password       string
	Encryption     string
	ConnectTimeout time.Duration
	SendTimeout    time.Duration
}

func loadSMTP() SMTPConfig {
	return SMTPConfig{
		Host:           getEnv("SMTP_HOST", "smtp.gmail.com"),
		Port:           getEnvAsInt("SMTP_PORT", 587),
		Username:       mustGet("SMTP_USERNAME"),
		Password:       mustGet("SMTP_PASSWORD"),
		Encryption:     getEnv("SMTP_ENCRYPTION", "tls"),
		ConnectTimeout: getEnvAsDuration("SMTP_CONNECT_TIMEOUT", "10s"),
		SendTimeout:    getEnvAsDuration("SMTP_SEND_TIMEOUT", "10s"),
	}

}
