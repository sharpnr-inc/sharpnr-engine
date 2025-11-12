package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CoreOrgRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	coreOrg := v1.Group("/core/org/")

	coreOrg.GET("get-details")
	coreOrg.POST("update")
	coreOrg.POST("add-bank-accounts")
}
