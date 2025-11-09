package concierge

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConciergeRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	croutes := v1.Group("/concierge")

	croutes.GET("/get-requests")

}
