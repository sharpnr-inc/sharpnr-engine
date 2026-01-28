package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/core/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerPropertyRoutes(r *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController) {

	pr := r.Group("/property")
	{
		pr.GET("", c.ListAllOrgProperty)
		pr.GET("/:propertyID", c.GetPropertyDetails)
		pr.POST("/:propertyID/add-domain", c.AddDomainInProperty)
		pr.POST("/:propertyID/add-modules", c.AddModulesInProperty)
	}

}
