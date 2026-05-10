// types.go

package cookieManager

import "net/http"

// Type represents a logical cookie role.
type Type string

const (
	// Temporary signup/onboarding guard
	Signup Type = "signup"

	// Authenticated session
	Auth Type = "auth"

	// Security
	CSRF Type = "csrf"
)

// Scope controls JS access.
type Scope int

const (
	// Public cookies are readable by JS (e.g. CSRF)
	Public Scope = iota

	// Protected cookies are HttpOnly
	Protected
)

// Definition is the canonical cookie definition.
type Definition struct {
	Name     string
	Scope    Scope
	TTL      int // seconds
	SameSite http.SameSite
	Secure   bool
	Path     string
	Domain   string
	HttpOnly bool
}
