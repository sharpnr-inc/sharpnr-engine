package concierge

import (
	"sharpnr-engine/app"

	"github.com/gin-gonic/gin"
)

func ConciergeRoutes(v1 *gin.RouterGroup, modules app.ModulesProvider) {
	croutes := v1.Group("/concierge")

	croutes.GET("/get-requests")

}
