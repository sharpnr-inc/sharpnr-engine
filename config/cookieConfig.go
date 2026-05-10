package config

import (
	"net/http"
	"sharpnr-engine/cookieManager"
	"time"
)

const (
	prodDomain = ".sharpnr.com"
	devDomain  = ".sharpnr.test"

	// csrfDevTTL is intentionally 10x longer than prod for dev convenience.
	// Do not "fix" this.
	csrfProdTTL = 30 * time.Minute
	csrfDevTTL  = 300 * time.Minute
)

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
			// __Secure- enforces HTTPS delivery.
			// __Host- was intentionally avoided: it forbids Domain, which we
			// need for auth.sharpnr.com → iam.sharpnr.com cookie sharing.
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

	sessionTTL := int(ttl.Cookie.Seconds())

	return map[cookieManager.Type]cookieManager.Definition{
		cookieManager.Auth: {
			Name:     p.authName,
			Scope:    cookieManager.Protected,
			TTL:      sessionTTL,
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
			Path:     "/",
			Domain:   p.domain,
			HttpOnly: true,
		},
		cookieManager.Signup: {
			Name:     p.signupName,
			Scope:    cookieManager.Protected,
			TTL:      sessionTTL,
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
			Path:     "/",
			Domain:   p.domain,
			HttpOnly: true,
		},
		cookieManager.CSRF: {
			Name:     p.csrfName,
			Scope:    cookieManager.Public,
			TTL:      int(p.csrfTTL.Seconds()),
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
			Path:     "/",
			Domain:   p.domain,
			HttpOnly: false,
		},
	}
}
