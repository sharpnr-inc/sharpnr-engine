package config

import (
	"net/http"
	"sharpnr-engine/pkg/cookieManager"
	"time"
)

const (
	prodDomain = ".sharpnr.com"
	devDomain  = ".sharpnr.test"

	csrfProdTTL = 120 * time.Minute
	// csrfDevTTL is intentionally 10x longer than prod for dev convenience.
	// Do not "fix" this.
	csrfDevTTL = 300 * time.Minute
)

func newDef(name, domain string, ttl time.Duration, httpOnly bool) cookieManager.Definition {
	return cookieManager.Definition{
		Name:     name,
		TTL:      int(ttl.Seconds()),
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Path:     "/",
		Domain:   domain,
		HttpOnly: httpOnly,
	}
}

func loadCookies(app AppConfig, ttl TTLConfig) map[cookieManager.Type]cookieManager.Definition {
	type envParams struct {
		authName   string
		signupName string
		csrfName   string
		domain     string
		csrfTTL    time.Duration
	}

	var p envParams
	if app.Env == "production" {
		p = envParams{
			authName:   "__Secure-Auth",
			signupName: "__Secure-Signup",
			csrfName:   "__Secure-CSRF",
			domain:     prodDomain,
			csrfTTL:    csrfProdTTL,
		}
	} else {
		p = envParams{
			authName:   "_dev_auth",
			signupName: "_dev_signup",
			csrfName:   "_dev_csrf",
			domain:     devDomain,
			csrfTTL:    csrfDevTTL,
		}
	}

	return map[cookieManager.Type]cookieManager.Definition{
		cookieManager.Auth:   newDef(p.authName, p.domain, ttl.Cookie, true),
		cookieManager.Signup: newDef(p.signupName, p.domain, ttl.Cookie, true),
		cookieManager.CSRF:   newDef(p.csrfName, p.domain, p.csrfTTL, false),
	}
}
