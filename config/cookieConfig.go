package config

import (
	"net/http"
	"sharpnr-engine/cookieManager"
	"time"
)

type CookieConfig struct {
	Domain string
	Secure bool
}

func loadCookie(app AppConfig, ttl TTLConfig) CookieConfig {
	if app.Env == "production" {
		return CookieConfig{
			Domain: "",
			Secure: true,
		}
	}

	// development / test
	return CookieConfig{
		Domain: "",
		Secure: false,
	}
}

func CookieRegistry(app AppConfig, ttl TTLConfig) map[cookieManager.Type]cookieManager.Definition {

	// ---------- PRODUCTION ----------
	if app.Env == "production" {
		return map[cookieManager.Type]cookieManager.Definition{
			cookieManager.Auth: {
				Name:     "__Host_Auth",
				Scope:    cookieManager.Protected,
				TTL:      int(ttl.Cookie.Seconds()),
				SameSite: http.SameSiteNoneMode,
				Secure:   true,
				Path:     "/",
				Domain:   "",
				HttpOnly: true,
			},
			cookieManager.Signup: {
				Name:     "__Host_Signup",
				Scope:    cookieManager.Protected,
				TTL:      int(ttl.Cookie.Seconds()),
				SameSite: http.SameSiteNoneMode,
				Secure:   true,
				Path:     "/",
				Domain:   "",
				HttpOnly: true,
			},
			cookieManager.CSRF: {
				Name:     "_CSRF",
				Scope:    cookieManager.Public,
				TTL:      int((30 * time.Minute).Seconds()),
				SameSite: http.SameSiteNoneMode,
				Secure:   true,
				Path:     "/",
				Domain:   ".sharpnr.com",
				HttpOnly: false,
			},
		}
	}

	// ---------- DEVELOPMENT ----------
	return map[cookieManager.Type]cookieManager.Definition{
		cookieManager.Auth: {
			Name:     "_dev_auth",
			Scope:    cookieManager.Protected,
			TTL:      int(ttl.Cookie.Seconds()),
			SameSite: http.SameSiteLaxMode,
			Secure:   false,
			Path:     "/",
		},
		cookieManager.Signup: {
			Name:     "_signup",
			Scope:    cookieManager.Protected,
			TTL:      int(ttl.Cookie.Seconds()),
			SameSite: http.SameSiteLaxMode,
			Secure:   false,
			Path:     "/",
		},
		cookieManager.CSRF: {
			Name:     "_dev_csrf",
			Scope:    cookieManager.Public,
			TTL:      int((300 * time.Minute).Seconds()),
			SameSite: http.SameSiteLaxMode,
			Secure:   false,
			Path:     "/",
		},
	}
}
