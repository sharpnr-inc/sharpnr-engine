package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

/*
TODO :- change the add modules to business-opperations
*/

func registerPropertyRoutes(r *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController) {
	pr := r.Group("/property")
	{
		pr.GET("", c.ListAllOrgProperty)
		pr.POST("/create", c.AddProperty)
		pr.GET("/:propertyID", c.GetPropertyDetails)

		// infra modules
		pr.GET("/:propertyID/infra_modules", c.GetPropertyInfra)
		pr.POST("/:propertyID/infra_modules")

		// business modules
		pr.GET("/:propertyID/business_modules")
		pr.POST("/:propertyID/business_modules")
	}
}
