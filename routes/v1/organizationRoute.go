// routes/v1/organizationRoute.go
package v1

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/sharpnr-inc/sharpnr-engine/controllers/organisationControllers"
)

func RegisterOrganizationRoutes(rg *gin.RouterGroup) {
	org := rg.Group("/org")
	{
		org.POST("/login", controllers.OrganizationRootLogin)
		org.POST("/register", controllers.RegisterOrganization)

	}
}
