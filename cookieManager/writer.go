package cookieManager

import "github.com/gin-gonic/gin"

func write(ctx *gin.Context, def Definition, value string) {
	ctx.SetSameSite(def.SameSite)
	ctx.SetCookie(def.Name, value, def.TTL, def.Path, def.Domain, def.Secure, def.HttpOnly)
}

func clear(ctx *gin.Context, def Definition) {
	ctx.SetSameSite(def.SameSite)
	ctx.SetCookie(def.Name, "", -1, def.Path, def.Domain, def.Secure, def.HttpOnly)
}
