package v1

import (
	"github.com/gin-gonic/gin"
	orgservices "github.com/sharpnr-inc/sharpnr-engine/core/services/orgnisation"
)

// RegisterV1Routes registers all v1 routes.
func RegisterV1Routes(router *gin.Engine, orgRegisterService *orgservices.OrganizationRegistrationService) {
	v1 := router.Group("/v1")

	// Register core routes
	registerCoreRoutes(v1, orgRegisterService)

	// Register module routes
	registerModuleRoutes(v1)
}