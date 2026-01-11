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

	ctx.SetCookie(
		def.Name,
		value,
		def.TTL,
		"/",
		w.domain,
		w.secure,
		def.Scope == Protected,
	)
}

func (w *writer) clear(ctx *gin.Context, def Definition) {
	ctx.SetCookie(
		def.Name,
		"",
		-1,
		"/",
		w.domain,
		w.secure,
		def.Scope == Protected,
	)
}
