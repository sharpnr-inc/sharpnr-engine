// writer.go
package cookieManager

import "github.com/gin-gonic/gin"

type writer struct {
	secure bool
	domain string
}

func newWriter(secure bool, domain string) *writer {
	return &writer{
		secure: secure,
		domain: domain,
	}
}

func (w *writer) set(ctx *gin.Context, def Definition, value string) {
	ctx.SetSameSite(def.SameSite)

	domain := w.domain
	if def.Domain != "" {
		domain = def.Domain // ⭐ override
	}

	ctx.SetCookie(
		def.Name,
		value,
		def.TTL,
		def.Path,
		domain,
		def.Secure,
		def.Scope == Protected,
	)
}

func (w *writer) clear(ctx *gin.Context, def Definition) {
	domain := w.domain
	if def.Domain != "" {
		domain = def.Domain
	}
	ctx.SetCookie(
		def.Name,
		"",
		-1,
		"/",
		domain,
		def.Secure, // 🔹 FIX: Use def.Secure, not w.secure
		def.Scope == Protected,
	)
}
