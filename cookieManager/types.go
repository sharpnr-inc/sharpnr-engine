package cookieManager

import "net/http"

type Type string

const (
	Signup Type = "signup"
	Auth   Type = "auth"
	CSRF   Type = "csrf"
)

type Definition struct {
	Name     string
	TTL      int // seconds
	SameSite http.SameSite
	Secure   bool
	Path     string
	Domain   string
	HttpOnly bool
}
