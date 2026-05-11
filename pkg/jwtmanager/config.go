// config.go
package jwt_manager

import "time"

type Config struct {
	Secret            string
	Issuer            string
	Audience          string
	AccessTTL         time.Duration
	RefreshTTL        time.Duration
	WebCookieTokenTTL time.Duration
	SignupTTL         time.Duration
}
