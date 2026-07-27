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

// loadCookies builds the cookie definitions from the lifetime of the token each
// cookie carries, not from a standalone COOKIE_TTL. A cookie is only a transport
// for its token: if its Max-Age outlives the token inside it, the browser keeps
// sending a cookie the server has already stopped accepting, and the user looks
// signed in until the next request 401s. Auth and Signup therefore inherit
// JWT_AUTH_WEB_COOKIE_TTL and JWT_SIGNUP_TTL respectively.
//
// The auth cookie is re-issued on every authenticated request by
// WebAuthMiddleware, so an active session keeps rolling forward regardless.
func loadCookies(app AppConfig, jwt JwtManagerConfigs) map[cookieManager.Type]cookieManager.Definition {
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
		cookieManager.Auth:   newDef(p.authName, p.domain, jwt.WebCookieTokenTTL, true),
		cookieManager.Signup: newDef(p.signupName, p.domain, jwt.SignupTTL, true),
		cookieManager.CSRF:   newDef(p.csrfName, p.domain, p.csrfTTL, false),
	}
}
