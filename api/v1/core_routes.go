package v1

import (
	"github.com/gin-gonic/gin"
	orgservices "github.com/sharpnr-inc/sharpnr-engine/core/services/orgnisation"
	"github.com/sharpnr-inc/sharpnr-engine/controllers/organisationControllers"
)

func registerCoreRoutes(rg *gin.RouterGroup, orgRegisterService *orgservices.OrganizationRegistrationService) {
	// Organization Registration Route
	rg.POST("/organizations/register", func(c *gin.Context) {
		organisationControllers.RegisterOrganization(c, orgRegisterService)
	})
}