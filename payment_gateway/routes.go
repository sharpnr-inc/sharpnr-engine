package paymentgateway

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PaymentGatewaysRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	pay_routes := v1.Group("/payments")

	pay_routes.POST("/create")
}
